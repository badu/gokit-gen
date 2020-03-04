package example_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	. "github.com/badu/gokit-gen/example/poc/gen"
	pb "github.com/badu/gokit-gen/example/poc/pb"
	kitJWT "github.com/go-kit/kit/auth/jwt"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type MockedRepository struct{}

func (r *MockedRepository) HalfDuplex(ctx context.Context, msg string) (string, error) {
	// TODO : implement me
	return "", errors.New("repository not implemented")
}

func (r *MockedRepository) FullDuplex(ctx context.Context, msg string) (string, error) {
	// TODO : implement me
	return "", errors.New("repository not implemented")
}

const (
	grpcHostAndPort = "localhost:8082"
)

func generateJWTMeta() string {
	// TODO : customize
	return ""
}

// half duplex for HalfDuplex
func TestGRPCHalfDuplex(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	grpcService := NewService(logger, &MockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(grpcService, []endpoint.Middleware{}) //authMiddleware
	grpcServer := grpc.NewServer()

	serverConn, err := net.Listen("tcp", grpcHostAndPort)
	if err != nil {
		panic(fmt.Sprintf("unable to listen: %+v", err))
	}
	defer grpcServer.GracefulStop()

	options := []kitGRPC.ServerOption{
		kitGRPC.ServerBefore(kitJWT.GRPCToContext()),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(logger)),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterStreamsServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	// simulate broadcasting messages
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		stopTimer := time.After(12 * time.Second)
		for {
			select {
			case <-ctx.Done():
				logger.Log("server", "stop")
				return
			case <-ticker.C:
				grpcService.BroadcastHalfDuplex() <- Reply{}
				logger.Log("broadcasting", "HalfDuplex <- Reply{}")
			case <-stopTimer:
				logger.Log("server", "end of life")
				ticker.Stop()
				cancel()
			}
		}
	}()
	// client receive loop
	go func() {
		logger.Log("client", "waiting for messages")
		for {
			select {
			case <-ctx.Done():
				t.Log("Context done")
				return
			case message := <-client.ReceiveHalfDuplex():
				// TODO : check response (write the actual test)

				t.Logf("msg : %v", message.Msg)
			}
		}
	}()
	logger.Log("client", "placing call to HalfDuplex")
	// TODO : load the payloads
	err = client.CallHalfDuplex(ctx, "some")
	if err != nil && err.Error() != "rpc error: code = Canceled desc = context canceled" {
		t.Fatalf("unable to test: %+v", err)
	}
	logger.Log("client", "end of test")
}

// full duplex for FullDuplex
func TestGRPCFullDuplex(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	grpcService := NewService(logger, &MockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(grpcService, []endpoint.Middleware{}) //authMiddleware
	grpcServer := grpc.NewServer()

	serverConn, err := net.Listen("tcp", grpcHostAndPort)
	if err != nil {
		panic(fmt.Sprintf("unable to listen: %+v", err))
	}
	defer grpcServer.GracefulStop()

	options := []kitGRPC.ServerOption{
		kitGRPC.ServerBefore(kitJWT.GRPCToContext()),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(logger)),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterStreamsServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	// simulate broadcasting messages
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		cliTicker := time.NewTicker(3 * time.Second)
		stopTimer := time.After(12 * time.Second)
		for {
			select {
			case <-ctx.Done():
				logger.Log("server", "stop")
				return
			case <-ticker.C:
				grpcService.BroadcastFullDuplex() <- Reply{}
				logger.Log("from_server", "FullDuplex <- Reply{}")
			case <-cliTicker.C:
				client.BroadcastFullDuplex() <- &pb.Request{}
				logger.Log("from_client", "FullDuplex <- &pb.Request{}")
			case <-stopTimer:
				logger.Log("server", "end of life")
				ticker.Stop()
				cliTicker.Stop()
				cancel()
			}
		}
	}()
	// client receive loop
	go func() {
		logger.Log("client", "waiting for messages")
		for {
			select {
			case <-ctx.Done():
				t.Log("Context done")
				return
			case message := <-client.ReceiveFullDuplex():
				// TODO : check response (write the actual test)

				t.Logf("msg : %v", message.Msg)
			}
		}
	}()
	logger.Log("client", "placing call to FullDuplex")
	// TODO : load the payloads
	err = client.CallFullDuplex(ctx)
	if err != nil && err.Error() != "rpc error: code = Canceled desc = context canceled" {
		t.Fatalf("unable to test: %+v", err)
	}
	logger.Log("client", "end of test")
}
