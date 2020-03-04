package factory_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	"google.golang.org/grpc"

	. "github.com/badu/gokit-gen/example/http/gen"
	pb "github.com/badu/gokit-gen/example/http/pb"
	kitJWT "github.com/go-kit/kit/auth/jwt"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type GRPCMockedRepository struct{}

func (r *GRPCMockedRepository) MakeBox(ctx context.Context, name string, height int32, width int32, depth int32) (int32, int32, error) {
	// TODO : implement me
	return 0, 0, errors.New("repository not implemented")
}

func (r *GRPCMockedRepository) GetBoxes(ctx context.Context, page int32, perPage int32) ([]BoxSpecification, error) {
	// TODO : implement me
	return nil, errors.New("repository not implemented")
}

func (r *GRPCMockedRepository) Status(ctx context.Context, page int32, perPage int32) (string, bool, error) {
	// TODO : implement me
	return "", false, errors.New("repository not implemented")
}

const (
	grpcHostAndPort = "localhost:8082"
)

func generateJWTMeta() string {
	// TODO : customize
	return ""
}

// classic grpc call for MakeBox
func TestGRPCMakeBox(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	grpcService := NewService(logger, &GRPCMockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(grpcService, []endpoint.Middleware{}) //authMiddleware
	grpcServer := grpc.NewServer()

	serverConn, err := net.Listen("tcp", grpcHostAndPort)
	if err != nil {
		panic(fmt.Sprintf("unable to listen: %+v", err))
	}
	defer grpcServer.GracefulStop()

	options := []kitGRPC.ServerOption{
		kitGRPC.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kitGRPC.ServerBefore(kitJWT.GRPCToContext()),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterBoxFactoryServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.Background()
	// TODO : load the payloads
	page, perPage, err := client.MakeBox(ctx, "", 0, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("page : %v", page)
	t.Logf("perPage : %v", perPage)
}

// classic grpc call for GetBoxes
func TestGRPCGetBoxes(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	grpcService := NewService(logger, &GRPCMockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(grpcService, []endpoint.Middleware{}) //authMiddleware
	grpcServer := grpc.NewServer()

	serverConn, err := net.Listen("tcp", grpcHostAndPort)
	if err != nil {
		panic(fmt.Sprintf("unable to listen: %+v", err))
	}
	defer grpcServer.GracefulStop()

	options := []kitGRPC.ServerOption{
		kitGRPC.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kitGRPC.ServerBefore(kitJWT.GRPCToContext()),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterBoxFactoryServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.Background()
	// TODO : load the payloads
	boxes, err := client.GetBoxes(ctx, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("boxes : %v", boxes)
}

// classic grpc call for Status
func TestGRPCStatus(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.Caller(4))
	logger = level.NewFilter(logger, level.AllowDebug())

	grpcService := NewService(logger, &GRPCMockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(logger)
	endpoints := MakeEndpoints(grpcService, []endpoint.Middleware{}) //authMiddleware
	grpcServer := grpc.NewServer()

	serverConn, err := net.Listen("tcp", grpcHostAndPort)
	if err != nil {
		panic(fmt.Sprintf("unable to listen: %+v", err))
	}
	defer grpcServer.GracefulStop()

	options := []kitGRPC.ServerOption{
		kitGRPC.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kitGRPC.ServerBefore(kitJWT.GRPCToContext()),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterBoxFactoryServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.Background()
	// TODO : load the payloads
	serviceName, ok, err := client.Status(ctx, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("serviceName : %v", serviceName)
	t.Logf("ok : %v", ok)
}
