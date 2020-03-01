package route_guide

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type GRPCServer struct {
	logger                       log.Logger
	GetFeatureGRPCHandler        kitGRPC.Handler
	HalfDuplexGRPCHandler        kitGRPC.Handler // TODO : half duplex
	ReverseHalfDuplexGRPCHandler kitGRPC.Handler // TODO : reverse half duplex
	FullDuplexGRPCHandler        kitGRPC.Handler // TODO : full duplex
}

func NewGRPCServer(endpoints Endpoints, logger log.Logger, options ...kitGRPC.ServerOption) (*GRPCServer, error) {
	return &GRPCServer{
		logger: logger,

		GetFeatureGRPCHandler: kitGRPC.NewServer(
			endpoints.GetFeatureEndpoint,
			decodeGetFeatureGRPCRequest(logger),
			encodeGetFeatureGRPCResponse(logger),
			options...,
		),
		HalfDuplexGRPCHandler: kitGRPC.NewServer(
			endpoints.HalfDuplexEndpoint,
			decodeHalfDuplexGRPCStreamRequest(logger),
			encodeHalfDuplexGRPCStreamResponse(logger),
			options..., // TODO : maybe no options for streaming?
		),
		ReverseHalfDuplexGRPCHandler: kitGRPC.NewServer(
			endpoints.ReverseHalfDuplexEndpoint,
			decodeReverseHalfDuplexGRPCStreamRequest(logger),
			encodeReverseHalfDuplexGRPCStreamResponse(logger),
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

func decodeGetFeatureGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.Point)
		if !ok {
			err := fmt.Errorf("expecting *pb.Point received %T", req)
			level.Error(logger).Log("decodeGetFeatureGRPCRequest", err)
			return nil, err
		}
		// TODO : validate protobuf
		return NewPointFromPB(pbReq), nil
	}
}

func encodeGetFeatureGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*Feature)
		if !ok {
			err := fmt.Errorf("expecting *Feature received %T", resp)
			level.Error(logger).Log("encodeGetFeatureGRPCRequest", err)
			return nil, err
		}
		return NewPBFromFeature(domResp), nil
	}
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
func decodeReverseHalfDuplexGRPCStreamRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return req, nil
	}
}

// streaming encoder : nothing to do, just pass it over
// it will be service responsibility to encode
func encodeReverseHalfDuplexGRPCStreamResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
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
	Request *pb.Rectangle
	Stream  pb.RouteGuide_HalfDuplexServer
}

// no streaming for GetFeature
func (s *GRPCServer) GetFeature(ctx context.Context, req *pb.Point) (*pb.Feature, error) {
	_, resp, err := s.GetFeatureGRPCHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	pbResp, ok := resp.(*pb.Feature)
	if !ok {
		err := fmt.Errorf("expecting *pb.Feature received %T", resp)
		return nil, err
	}
	return pbResp, nil
}

// half duplex for HalfDuplex
func (s *GRPCServer) HalfDuplex(req *pb.Rectangle, stream pb.RouteGuide_HalfDuplexServer) error {
	reqNStream := &RequestAndStreamHalfDuplex{Request: req, Stream: stream}
	_, _, err := s.HalfDuplexGRPCHandler.ServeGRPC(stream.Context(), reqNStream)
	return err
}

// reverse half duplex for ReverseHalfDuplex
func (s *GRPCServer) ReverseHalfDuplex(stream pb.RouteGuide_ReverseHalfDuplexServer) error {
	_, _, err := s.ReverseHalfDuplexGRPCHandler.ServeGRPC(stream.Context(), stream)
	return err
}

// full duplex for FullDuplex
func (s *GRPCServer) FullDuplex(stream pb.RouteGuide_FullDuplexServer) error {
	_, _, err := s.FullDuplexGRPCHandler.ServeGRPC(stream.Context(), stream)
	return err
}
