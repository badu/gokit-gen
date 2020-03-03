package route_guide_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

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
	return "test", &Point{
		Latitude:  10,
		Longitude: 10,
	}, nil
}

func (r *MockedRepository) HalfDuplex(ctx context.Context, lo *Point, hi *Point) (string, *Point, error) {
	return "test", &Point{
		Latitude:  10,
		Longitude: 10,
	}, nil
}

func (r *MockedRepository) ReverseHalfDuplex(ctx context.Context, latitude int32, longitude int32) (int32, int32, int32, int32, error) {
	return 1, 2, 3, 4, nil
}

func (r *MockedRepository) FullDuplex(ctx context.Context, location *Point, message string) (*Point, string, error) {
	return &Point{
		Latitude:  10,
		Longitude: 10,
	}, "test", nil
}

const (
	grpcHostAndPort = "localhost:8082"
)

func generateJWTMeta() string {
	// TODO : customize
	return ""
}

// classic grpc call for GetFeature
func TestGRPCGetFeature(t *testing.T) {
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
		//kitGRPC.ServerBefore(MakeGRPCDecodeJWTToClaims(jwtSecret, jwt.SigningMethodHS256, logger)),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(logger)),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterRouteGuideServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

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
		//kitGRPC.ServerBefore(MakeGRPCDecodeJWTToClaims(jwtSecret, jwt.SigningMethodHS256, logger)),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(logger)),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterRouteGuideServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.WithValue(context.Background(), kitJWT.JWTTokenContextKey, generateJWTMeta())
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
				grpcService.BroadcastHalfDuplex() <- Feature{}
				logger.Log("broadcasting", "HalfDuplex <- Feature{}")
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

				t.Logf("name : %v", message.Name)
				t.Logf("location : %v", message.Location)
			}
		}
	}()
	logger.Log("client", "placing call to HalfDuplex")
	// TODO : load the payloads
	err = client.CallHalfDuplex(ctx, &Point{Latitude: 0, Longitude: 0}, &Point{Latitude: 0, Longitude: 0})
	if err != nil && err.Error() != "rpc error: code = Canceled desc = context canceled" {
		t.Fatalf("unable to test: %+v", err)
	}
	logger.Log("client", "end of test")
}

// reverse half duplex for ReverseHalfDuplex
func TestGRPCReverseHalfDuplex(t *testing.T) {
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
		//kitGRPC.ServerBefore(MakeGRPCDecodeJWTToClaims(jwtSecret, jwt.SigningMethodHS256, logger)),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(logger)),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterRouteGuideServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.WithValue(context.Background(), kitJWT.JWTTokenContextKey, generateJWTMeta())
	// simulate broadcasting messages
	go func() {
		cliTicker := time.NewTicker(3 * time.Second)
		stopTimer := time.After(12 * time.Second)
		for {
			select {
			case <-ctx.Done():
				logger.Log("server", "stop")
				return
			case <-cliTicker.C:
				client.BroadcastReverseHalfDuplex() <- &pb.Point{}
				logger.Log("from_client", "ReverseHalfDuplex <- &pb.Point{}")
			case <-stopTimer:
				logger.Log("server", "end of life")
				client.CloseReverseHalfDuplex() <- struct{}{}
				cliTicker.Stop()
				return
			}
		}
	}()
	// TODO : load the payloads
	reply, err := client.CallReverseHalfDuplex(ctx)
	if err != nil && err.Error() != "rpc error: code = Canceled desc = context canceled" {
		t.Fatalf("unable to test: %+v", err)
	}
	// TODO : check response (write the actual test)

	t.Logf("pointCount : %v", reply.PointCount)
	t.Logf("featureCount : %v", reply.FeatureCount)
	t.Logf("distance : %v", reply.Distance)
	t.Logf("elapsedTime : %v", reply.ElapsedTime)
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
		//kitGRPC.ServerBefore(MakeGRPCDecodeJWTToClaims(jwtSecret, jwt.SigningMethodHS256, logger)),
		//kitGRPC.ServerBefore(MakeAddStreamUUID(logger)),
	}

	service, err := NewGRPCServer(endpoints, logger, options...)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	go func() {
		pb.RegisterRouteGuideServer(grpcServer, service)
		_ = grpcServer.Serve(serverConn)
	}()

	conn, err := grpc.Dial(grpcHostAndPort, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to Dial: %+v", err)
	}

	clientOptions := []kitGRPC.ClientOption{kitGRPC.ClientBefore(kitJWT.ContextToGRPC())}
	client := NewClient(conn, logger, clientOptions...)

	ctx := context.WithValue(context.Background(), kitJWT.JWTTokenContextKey, generateJWTMeta())
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
				grpcService.BroadcastFullDuplex() <- RouteNote{}
				logger.Log("from_server", "FullDuplex <- RouteNote{}")
			case <-cliTicker.C:
				client.BroadcastFullDuplex() <- &pb.RouteNote{}
				logger.Log("from_client", "FullDuplex <- &pb.RouteNote{}")
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

				t.Logf("location : %v", message.Location)
				t.Logf("message : %v", message.Message)
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
