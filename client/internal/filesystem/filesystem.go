package filesystem

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"

	pb "github.com/poy/loggrebutterfly/api/v1"
	"github.com/poy/petasos/reader"
	"github.com/poy/petasos/router"
)

type RouteCache interface {
	List() []string
	FetchRoute(name string) (client pb.DataNodeClient, addr string)
	Reset()
}

type FileSystem struct {
	cache RouteCache
}

func New(cache RouteCache) *FileSystem {
	return &FileSystem{
		cache: cache,
	}
}

func (f *FileSystem) List() (files []string, err error) {
	list := f.cache.List()
	if len(list) == 0 {
		return nil, fmt.Errorf("unable to fetch route list")
	}
	return list, nil
}

func (f *FileSystem) Writer(name string) (writer router.Writer, err error) {
	client, addr := f.cache.FetchRoute(name)
	if client == nil {
		return nil, fmt.Errorf("unknown file: %s", name)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	sender, err := client.Write(ctx)
	if err != nil {
		f.cache.Reset()
		return nil, err
	}

	wrapper := &senderWrapper{
		sender: sender,
		addr:   addr,
		reset:  f.cache.Reset,
	}

	return wrapper, nil
}

func (f *FileSystem) Reader(name string, startingIndex uint64) (reader reader.Reader, err error) {
	client, addr := f.cache.FetchRoute(name)
	if client == nil {
		return nil, fmt.Errorf("unknown file: %s", name)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	rx, err := client.Read(ctx, &pb.ReadInfo{Name: name, Index: startingIndex})
	if err != nil {
		return nil, err
	}

	return &receiverWrapper{rx: rx, cancel: cancel, addr: addr, resetCache: f.cache.Reset}, nil
}

type senderWrapper struct {
	addr   string
	sender pb.DataNode_WriteClient
	err    error
	reset  func()
}

func (w *senderWrapper) Write(data []byte) error {
	if w.err != nil {
		return w.err
	}

	err := w.sender.Send(&pb.WriteInfo{Payload: data})
	if err != nil {
		w.err = err
		w.reset()

		return fmt.Errorf("[WRITE TO %s]: %s", w.addr, w.err)
	}

	return nil
}

func (w *senderWrapper) Close() {
	w.sender.CloseAndRecv()
}

type receiverWrapper struct {
	addr       string
	rx         pb.DataNode_ReadClient
	cancel     func()
	resetCache func()
}

func (w *receiverWrapper) Read() (reader.DataPacket, error) {
	data, err := w.rx.Recv()
	if err != nil && grpc.ErrorDesc(err) == "EOF" {
		w.resetCache()
		return reader.DataPacket{}, io.EOF
	}

	if err != nil {
		w.resetCache()
		return reader.DataPacket{}, fmt.Errorf("[READ FROM %s]: %s", w.addr, err)
	}

	return reader.DataPacket{
		Payload:  data.Payload,
		Filename: data.File,
		Index:    data.Index,
	}, nil
}

func (w *receiverWrapper) Close() {
	w.cancel()
}
