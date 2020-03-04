package example

import (
	"context"

	"github.com/go-kit/kit/log"

	pb "github.com/badu/gokit-gen/example/poc/pb"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type GRPCServer struct {
	logger                log.Logger
	HalfDuplexGRPCHandler kitGRPC.Handler // TODO : half duplex
	FullDuplexGRPCHandler kitGRPC.Handler // TODO : full duplex
}

func NewGRPCServer(endpoints Endpoints, logger log.Logger, options ...kitGRPC.ServerOption) (*GRPCServer, error) {
	return &GRPCServer{
		logger: logger,

		HalfDuplexGRPCHandler: kitGRPC.NewServer(
			endpoints.HalfDuplexEndpoint,
			decodeHalfDuplexGRPCStreamRequest(logger),
			encodeHalfDuplexGRPCStreamResponse(logger),
			options..., // TODO : maybe no options for streaming?
		),
		FullDuplexGRPCHandler: kitGRPC.NewServer(
			endpoints.FullDuplexEndpoint,
			decodeFullDuplexGRPCStreamRequest(logger),
			encodeFullDuplexGRPCStreamResponse(logger),
			options..., // TODO : maybe no options for streaming?
		),
	}, nil
}

// streaming decoder : nothing to do, just pass it over
// it will be service responsibility to decode
func decodeHalfDuplexGRPCStreamRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return req, nil
	}
}

// streaming encoder : nothing to do, just pass it over
// it will be service responsibility to encode
func encodeHalfDuplexGRPCStreamResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		return resp, nil
	}
}

// streaming decoder : nothing to do, just pass it over
// it will be service responsibility to decode
func decodeFullDuplexGRPCStreamRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return req, nil
	}
}

// streaming encoder : nothing to do, just pass it over
// it will be service responsibility to encode
func encodeFullDuplexGRPCStreamResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		return resp, nil
	}
}

// request and stream for half duplex
type RequestAndStreamHalfDuplex struct {
	Request *pb.Request
	Stream  pb.Streams_HalfDuplexServer
}

// protobuf implementation : half duplex for HalfDuplex
func (s *GRPCServer) HalfDuplex(req *pb.Request, stream pb.Streams_HalfDuplexServer) error {
	reqNStream := &RequestAndStreamHalfDuplex{Request: req, Stream: stream}
	_, _, err := s.HalfDuplexGRPCHandler.ServeGRPC(stream.Context(), reqNStream)
	return err
}

// protobuf implementation : full duplex for FullDuplex
func (s *GRPCServer) FullDuplex(stream pb.Streams_FullDuplexServer) error {
	_, _, err := s.FullDuplexGRPCHandler.ServeGRPC(stream.Context(), stream)
	return err
}
