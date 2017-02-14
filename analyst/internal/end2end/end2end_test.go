package end2end_test

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"sort"
	"testing"
	"time"

	"github.com/apoydence/eachers/testhelpers"
	v2 "github.com/apoydence/loggrebutterfly/api/loggregator/v2"
	loggrebutterfly "github.com/apoydence/loggrebutterfly/api/v1"
	"github.com/apoydence/loggrebutterfly/internal/end2end"
	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
	talaria "github.com/apoydence/talaria/api/v1"
	"github.com/golang/protobuf/proto"
	"github.com/onsi/gomega/gexec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	analystPort      int
	analystIntraPort int

	nodeAddr, schedAddr string
	mockNode            *mockNodeServer
	mockSched           *mockSchedulerServer
)

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Verbose() {
		log.SetOutput(ioutil.Discard)
		grpclog.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags))
	}

	ps := setup()
	go serveUpData()
	go serveScheduler()

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

type TA struct {
	*testing.T
	client loggrebutterfly.AnalystClient
}

func TestAnalyst(t *testing.T) {
	t.Parallel()
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) TA {
		client := startClient(fmt.Sprintf("localhost:%d", analystPort))

		return TA{
			T:      t,
			client: client,
		}
	})

	o.Spec("it returns only the envelopes from the right time range", func(t TA) {
		var results *loggrebutterfly.QueryResponse
		f := func() bool {
			var err error
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			results, err = t.client.Query(ctx, &loggrebutterfly.QueryInfo{
				SourceUuid: "some-id",
				TimeRange: &loggrebutterfly.TimeRange{
					Start: 10,
					End:   20,
				},
			})
			return err == nil
		}

		Expect(t, f).To(ViaPollingMatcher{
			Duration: 5 * time.Second,
			Matcher:  BeTrue(),
		})

		sort.Sort(envelopes(results.Envelopes))

		Expect(t, results.Envelopes).To(HaveLen(10))
		for i, e := range results.Envelopes {
			Expect(t, e.SourceUuid).To(Equal("some-id"))
			Expect(t, e.Timestamp).To(Equal(int64(i + 10)))
		}
	})
}

func serveUpData() {
	for {
		rx := <-mockNode.ReadInput.Arg1
		go func(rx talaria.Node_ReadServer) {
			for i := 0; i < 100; i++ {
				e := &v2.Envelope{
					SourceUuid: "some-id",
					Timestamp:  int64(i),
				}
				data, err := proto.Marshal(e)
				if err != nil {
					panic(err)
				}
				err = rx.Send(&talaria.ReadDataPacket{Message: data})
				if err != nil {
					panic(err)
				}
			}

			mockNode.ReadOutput.Ret0 <- io.EOF
		}(rx)
	}
}

func serveScheduler() {
	close(mockSched.ListClusterInfoOutput.Ret1)
	testhelpers.AlwaysReturn(mockSched.ListClusterInfoOutput.Ret0, &talaria.ListResponse{
		Info: []*talaria.ClusterInfo{
			{
				Name: `{"Low":0,"High":18446744073709551615}`,
				Nodes: []*talaria.NodeInfo{
					{
						URI: nodeAddr,
					},
				},
			},
		},
	})
}

func startClient(addr string) loggrebutterfly.AnalystClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return loggrebutterfly.NewAnalystClient(conn)
}

func setup() []*os.Process {
	nodeAddr, mockNode = startMockNode()
	schedAddr, mockSched = startMockSched()

	analystPort = end2end.AvailablePort()
	analystIntraPort = end2end.AvailablePort()
	analystPs := startAnalyst(analystPort, analystIntraPort, nodeAddr, schedAddr)

	return []*os.Process{
		analystPs,
	}
}

func startAnalyst(port, intraPort int, dataNodeAddr, schedAddr string) *os.Process {
	log.Printf("Starting analyst on %d...", port)
	defer log.Printf("Done starting analyst on %d.", port)

	path, err := gexec.Build("github.com/apoydence/loggrebutterfly/analyst")
	if err != nil {
		panic(err)
	}

	command := exec.Command(path)
	command.Env = []string{
		fmt.Sprintf("ADDR=127.0.0.1:%d", port),
		fmt.Sprintf("INTRA_ADDR=127.0.0.1:%d", intraPort),
		fmt.Sprintf("TALARIA_NODE_ADDR=%s", dataNodeAddr),
		fmt.Sprintf("TALARIA_SCHEDULER_ADDR=%s", schedAddr),
		fmt.Sprintf("TALARIA_NODE_LIST=%s", dataNodeAddr),
		fmt.Sprintf("INTRA_ANALYST_LIST=127.0.0.1:%d", intraPort),
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
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	mockNodeServer := newMockNodeServer()
	s := grpc.NewServer()
	talaria.RegisterNodeServer(s, mockNodeServer)

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	log.Printf("Starting talaria node on %s", lis.Addr().String())
	return lis.Addr().String(), mockNodeServer
}

func startMockSched() (string, *mockSchedulerServer) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	mockSchedServer := newMockSchedulerServer()
	s := grpc.NewServer()
	talaria.RegisterSchedulerServer(s, mockSchedServer)

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	log.Printf("Starting talaria sched on %s", lis.Addr().String())
	return lis.Addr().String(), mockSchedServer
}

type envelopes []*v2.Envelope

func (e envelopes) Len() int {
	return len(e)
}

func (e envelopes) Less(i, j int) bool {
	return e[i].Timestamp < e[j].Timestamp
}

func (e envelopes) Swap(i, j int) {
	tmp := e[i]
	e[i] = e[j]
	e[j] = tmp
}