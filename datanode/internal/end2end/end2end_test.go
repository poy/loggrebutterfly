package end2end_test

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/poy/eachers/testhelpers"
	"github.com/poy/loggrebutterfly/api/intra"
	v2 "github.com/poy/loggrebutterfly/api/loggregator/v2"
	pb "github.com/poy/loggrebutterfly/api/v1"
	"github.com/poy/loggrebutterfly/internal/end2end"
	"github.com/poy/onpar"
	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
	"github.com/poy/petasos/router"
	talariapb "github.com/poy/talaria/api/v1"
	"github.com/golang/protobuf/proto"
	"github.com/onsi/gomega/gexec"
)

var (
	dataNodePort      int
	dataNodeIntraPort int
	nodeAddr          string
	mockNode          *mockNodeServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
		grpclog.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags))
	}

	ps := setup()

	var status int
	func() {
		defer func() {
			for _, p := range ps {
				p.Kill()
			}
		}()
		status = m.Run()
	}()

	os.Exit(status)
}

type TDN struct {
	*testing.T
	client      pb.DataNodeClient
	intraClient intra.DataNodeClient
	fileName    string
}

func TestDataNode(t *testing.T) {
	t.Parallel()
	o := onpar.New()
	defer o.Run(t)

	fileName := buildRangeName(0, 18446744073709551615, 0)
	testhelpers.AlwaysReturn(mockNode.ListClustersOutput.Ret0, &talariapb.ListClustersResponse{
		Names: []string{fileName},
	})
	close(mockNode.ListClustersOutput.Ret1)
	close(mockNode.WriteOutput.Ret0)

	o.BeforeEach(func(t *testing.T) TDN {
		return TDN{
			T:           t,
			client:      fetchClient(fmt.Sprintf("127.0.0.1:%d", dataNodePort)),
			intraClient: fetchIntraClient(fmt.Sprintf("127.0.0.1:%d", dataNodeIntraPort)),
			fileName:    fileName,
		}
	})

	o.Spec("it writes to the talaria node", func(t TDN) {
		e := &v2.Envelope{
			SourceId: "some-id",
		}

		data, err := proto.Marshal(e)
		Expect(t, err == nil).To(BeTrue())

		f := func() bool {
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			sender, err := t.client.Write(ctx)
			if err != nil {
				return false
			}

			return sender.Send(&pb.WriteInfo{Payload: data}) == nil
		}
		Expect(t, f).To(ViaPollingMatcher{
			Duration: 5 * time.Second,
			Matcher:  BeTrue(),
		})

		var resp *intra.ReadMetricsResponse
		f = func() bool {
			var err error
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			resp, err = t.intraClient.ReadMetrics(ctx, &intra.ReadMetricsInfo{t.fileName})
			if err != nil {
				return false
			}

			return resp.WriteCount != 0
		}

		Expect(t, f).To(ViaPolling(BeTrue()))
		Expect(t, resp.WriteCount).To(Equal(uint64(1)))
	})

	o.Spec("it reads from the talaria node", func(t TDN) {
		defer close(mockNode.ReadOutput.Ret0)

		go func() {
			var server talariapb.Node_ReadServer
			Expect(t, mockNode.ReadInput.Arg1).To(ViaPollingMatcher{
				Duration: 5 * time.Second,
				Matcher:  Chain(Receive(), Fetch(&server)),
			})

			server.Send(&talariapb.ReadDataPacket{
				Message: []byte("some-data"),
				Index:   99,
			})
			server.Send(&talariapb.ReadDataPacket{
				Message: []byte("some-other-data"),
				Index:   101,
			})
		}()

		var rx pb.DataNode_ReadClient
		f := func() bool {
			var err error
			rx, err = t.client.Read(context.Background(), &pb.ReadInfo{Name: t.fileName})
			return err == nil
		}

		Expect(t, f).To(ViaPollingMatcher{
			Duration: 5 * time.Second,
			Matcher:  BeTrue(),
		})

		rxData, err := rx.Recv()
		Expect(t, err == nil).To(BeTrue())
		Expect(t, rxData.Payload).To(Equal([]byte("some-data")))

		var rn router.RangeName
		err = json.Unmarshal([]byte(rxData.File), &rn)
		Expect(t, err == nil).To(BeTrue())
		Expect(t, rn.Low).To(Equal(uint64(0)))
		Expect(t, rn.High).To(Equal(uint64(18446744073709551615)))
		Expect(t, rxData.Index).To(Equal(uint64(99)))

		rxData, err = rx.Recv()
		Expect(t, err == nil).To(BeTrue())
		Expect(t, rxData.Payload).To(Equal([]byte("some-other-data")))

		err = json.Unmarshal([]byte(rxData.File), &rn)
		Expect(t, err == nil).To(BeTrue())
		Expect(t, rn.Low).To(Equal(uint64(0)))
		Expect(t, rn.High).To(Equal(uint64(18446744073709551615)))
		Expect(t, rxData.Index).To(Equal(uint64(101)))

		file := buildRangeName(0, 18446744073709551615, 0)
		f = func() bool {
			var err error
			rx, err = t.client.Read(context.Background(), &pb.ReadInfo{
				Name:  t.fileName,
				Index: 99,
			})
			return err == nil
		}
		Expect(t, f).To(ViaPolling(BeTrue()))

		Expect(t, mockNode.ReadInput.Arg0).To(ViaPolling(
			Chain(Receive(), Equal(&talariapb.BufferInfo{
				Name:       file,
				StartIndex: 99,
			})),
		))
	})
}

func setup() []*os.Process {
	nodeAddr, mockNode = startMockNode()

	dataNodePort = end2end.AvailablePort()
	dataNodeIntraPort = end2end.AvailablePort()
	dataNodePs := startDataNode(dataNodePort, dataNodeIntraPort, nodeAddr)

	return []*os.Process{
		dataNodePs,
	}
}

func startDataNode(port, intraPort int, nodeAddr string) *os.Process {
	log.Printf("Starting data node on %d...", port)
	defer log.Printf("Done starting data node on %d.", port)

	path, err := gexec.Build("github.com/poy/loggrebutterfly/datanode")
	if err != nil {
		panic(err)
	}

	command := exec.Command(path)
	command.Env = []string{
		fmt.Sprintf("ADDR=127.0.0.1:%d", port),
		fmt.Sprintf("INTRA_ADDR=127.0.0.1:%d", intraPort),
		fmt.Sprintf("NODE_ADDR=%s", nodeAddr),
	}

	if testing.Verbose() {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}

	if err = command.Start(); err != nil {
		panic(err)
	}

	return command.Process
}

func startMockNode() (string, *mockNodeServer) {
	mockNodeServer := newMockNodeServer()
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	talariapb.RegisterNodeServer(s, mockNodeServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return lis.Addr().String(), mockNodeServer
}

func fetchClient(addr string) pb.DataNodeClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewDataNodeClient(conn)
}

func fetchIntraClient(addr string) intra.DataNodeClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return intra.NewDataNodeClient(conn)
}

func buildRangeName(low, high, term uint64) string {
	rn := router.RangeName{
		Low:  low,
		High: high,
		Term: term,
	}

	j, _ := json.Marshal(rn)
	return string(j)
}
