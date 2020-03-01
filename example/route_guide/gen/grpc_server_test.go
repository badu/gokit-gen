package route_guide_test

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
	"google.golang.org/grpc"

	. "github.com/badu/gokit-gen/example/route_guide/gen"
	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
	kitJWT "github.com/go-kit/kit/auth/jwt"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type MockedRepository struct{}

func (r *MockedRepository) GetFeature(ctx context.Context, latitude int32, longitude int32) (string, *Point, error) {
	// TODO : implement me
	return "", nil, errors.New("not implemented")
}

func (r *MockedRepository) HalfDuplex(ctx context.Context, lo *Point, hi *Point) (string, *Point, error) {
	// TODO : implement me
	return "", nil, errors.New("not implemented")
}

func (r *MockedRepository) ReverseHalfDuplex(ctx context.Context, latitude int32, longitude int32) (int32, int32, int32, int32, error) {
	// TODO : implement me
	return 0, 0, 0, 0, errors.New("not implemented")
}

func (r *MockedRepository) FullDuplex(ctx context.Context, location *Point) (*Point, error) {
	// TODO : implement me
	return nil, errors.New("not implemented")
}

const (
	grpcHostAndPort = "localhost:8082"
)

func generateJWTMeta() string {
	// TODO : customize
	return ""
}

func TestMain(m *testing.M) {
	serverLogger := log.NewLogfmtLogger(os.Stderr)
	serverLogger = log.With(serverLogger, "caller", log.Caller(7))
	serverLogger = level.NewFilter(serverLogger, level.AllowDebug())

	grpcService := NewService(serverLogger, &MockedRepository{})
	//authMiddleware := MakeEnsureValidJWTMiddleware(serverLogger)
	endpoints := MakeEndpoints(grpcService, []endpoint.Middleware{}) //authMiddleware
	grpcServer := grpc.NewServer()

	serverConn, err := net.Listen("tcp", grpcHostAndPort)
	if err != nil {
		panic(fmt.Sprintf("unable to listen: %+v", err))
	}
	defer grpcServer.GracefulStop()

	options := []kitGRPC.ServerOption{
		kitGRPC.ServerBefore(kitJWT.GRPCToContext()),
		//kitGRPC.ServerBefore(MakeGRPCDecodeJWTToClaims(jwtSecret, jwt.SigningMethodHS256, serverLogger)),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(serverLogger)),
	}

	service, err := NewGRPCServer(endpoints, serverLogger, options...)
	if err != nil {
		level.Error(serverLogger).Log("error", err)
	}

	go func() {
		pb.RegisterRouteGuideServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGRPCGetFeature(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.Caller(7))
	logger = level.NewFilter(logger, level.AllowDebug())

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.WithValue(context.Background(), kitJWT.JWTTokenContextKey, generateJWTMeta())
	// TODO : load the payloads
	name, location, err := client.GetFeature(ctx, 0, 0)
	if err != nil {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("name : %v", name)
	t.Logf("location : %v", location)
}

// TODO : half duplex for HalfDuplex

// TODO : reverse half duplex for ReverseHalfDuplex

// TODO : full duplex for FullDuplex
