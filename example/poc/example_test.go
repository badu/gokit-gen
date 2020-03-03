package poc_test

import (
	"context"
	"github.com/badu/gokit-gen/example/poc"
	"github.com/badu/gokit-gen/example/poc/pb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"net"
	"os"
	"testing"
)

const (
	hostPort string = "localhost:8002"
)

var logger log.Logger

func createLogger() log.Logger {
	// Create a single logger, which we'll use and give to other components.

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.DefaultCaller) // add caller

	logger = level.NewFilter(logger, level.AllowDebug())
	return logger
}

func TestMain(m *testing.M) {
	logger = createLogger()
	var (
		server  = grpc.NewServer()
		service = poc.NewService(logger)
	)

	srvConn, err := net.Listen("tcp", hostPort)
	if err != nil {
		logger.Log("error", "unable to listen: "+err.Error())
	}
	defer server.GracefulStop()

	go func() {
		pb.RegisterStreamsServer(server, poc.NewBinding(service))
		_ = server.Serve(srvConn)
	}()
	exitVal := m.Run()
	os.Exit(exitVal)

}
func TestHalfDuplexDirect(t *testing.T) {
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	client := poc.NewDirectClient(
		logger,
		conn,
		poc.EncodeRequestToServer(logger),
		poc.DecodeResponseFromServer(logger),
	)

	client.ApplyExtraOptions(
		poc.ClientBefore(
			poc.InjectCorrelationID(logger),
			poc.DisplayClientRequestHeaders(logger),
		),
		poc.ClientAfter(
			poc.DisplayClientResponseHeaders(logger),
			poc.DisplayClientResponseTrailers(logger),
			poc.ExtractConsumedCorrelationID(logger),
		),
	)
	var (
		cID = "request-1"
		ctx = poc.SetCorrelationID(context.Background(), cID)
	)

	err = client.CallHalfDuplex(ctx, &poc.InternalRequest{Msg: "Badabam"})
	if err != nil && err.Error() != "rpc error: code = Unknown desc = communication end" {
		t.Fatalf("unable to Test: %+v\n", err)
	}
}

func TestFullDuplexDirect(t *testing.T) {
	conn, err := grpc.Dial(hostPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	client := poc.NewDirectClient(
		logger,
		conn,
		poc.EncodeRequestToServer(logger),
		poc.DecodeResponseFromServer(logger),
	)

	client.ApplyExtraOptions(
		poc.ClientBefore(
			poc.InjectCorrelationID(logger),
			poc.DisplayClientRequestHeaders(logger),
		),
		poc.ClientAfter(
			poc.DisplayClientResponseHeaders(logger),
			poc.DisplayClientResponseTrailers(logger),
			poc.ExtractConsumedCorrelationID(logger),
		),
	)
	var (
		cID = "request-1"
		ctx = poc.SetCorrelationID(context.Background(), cID)
	)

	err = client.CallFullDuplex(ctx)
	if err != nil && err.Error() != "rpc error: code = Unknown desc = communication end" {
		t.Fatalf("unable to Test: %+v\n", err)
	}
}
