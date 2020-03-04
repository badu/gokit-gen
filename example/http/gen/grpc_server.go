package factory

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	pb "github.com/badu/gokit-gen/example/http/pb"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type GRPCServer struct {
	logger              log.Logger
	MakeBoxGRPCHandler  kitGRPC.Handler
	GetBoxesGRPCHandler kitGRPC.Handler
	StatusGRPCHandler   kitGRPC.Handler
}

func NewGRPCServer(endpoints Endpoints, logger log.Logger, options ...kitGRPC.ServerOption) (*GRPCServer, error) {
	return &GRPCServer{
		logger: logger,

		MakeBoxGRPCHandler: kitGRPC.NewServer(
			endpoints.MakeBoxEndpoint,
			decodeMakeBoxGRPCRequest(logger),
			encodeMakeBoxGRPCResponse(logger),
			options...,
		),
		GetBoxesGRPCHandler: kitGRPC.NewServer(
			endpoints.GetBoxesEndpoint,
			decodeGetBoxesGRPCRequest(logger),
			encodeGetBoxesGRPCResponse(logger),
			options...,
		),
		StatusGRPCHandler: kitGRPC.NewServer(
			endpoints.StatusEndpoint,
			decodeStatusGRPCRequest(logger),
			encodeStatusGRPCResponse(logger),
			options...,
		),
	}, nil
}

func decodeMakeBoxGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.BoxSpecification)
		if !ok {
			err := fmt.Errorf("expecting *pb.BoxSpecification received %T", req)
			level.Error(logger).Log("decodeMakeBoxGRPCRequest", err)
			return nil, err
		}
		// TODO : validate protobuf
		return NewBoxSpecificationFromPB(pbReq), nil
	}
}

func encodeMakeBoxGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*Empty)
		if !ok {
			err := fmt.Errorf("expecting *Empty received %T", resp)
			level.Error(logger).Log("encodeMakeBoxGRPCRequest", err)
			return nil, err
		}
		return NewPBFromEmpty(domResp), nil
	}
}

func decodeGetBoxesGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.Empty)
		if !ok {
			err := fmt.Errorf("expecting *pb.Empty received %T", req)
			level.Error(logger).Log("decodeGetBoxesGRPCRequest", err)
			return nil, err
		}
		// TODO : validate protobuf
		return NewEmptyFromPB(pbReq), nil
	}
}

func encodeGetBoxesGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*Boxes)
		if !ok {
			err := fmt.Errorf("expecting *Boxes received %T", resp)
			level.Error(logger).Log("encodeGetBoxesGRPCRequest", err)
			return nil, err
		}
		return NewPBFromBoxes(domResp), nil
	}
}

func decodeStatusGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*pb.Empty)
		if !ok {
			err := fmt.Errorf("expecting *pb.Empty received %T", req)
			level.Error(logger).Log("decodeStatusGRPCRequest", err)
			return nil, err
		}
		// TODO : validate protobuf
		return NewEmptyFromPB(pbReq), nil
	}
}

func encodeStatusGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*StatusMessage)
		if !ok {
			err := fmt.Errorf("expecting *StatusMessage received %T", resp)
			level.Error(logger).Log("encodeStatusGRPCRequest", err)
			return nil, err
		}
		return NewPBFromStatusMessage(domResp), nil
	}
}

// protobuf implementation : no streaming for MakeBox
func (s *GRPCServer) MakeBox(ctx context.Context, req *pb.BoxSpecification) (*pb.Empty, error) {
	_, resp, err := s.MakeBoxGRPCHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	pbResp, ok := resp.(*pb.Empty)
	if !ok {
		err := fmt.Errorf("expecting *pb.Empty received %T", resp)
		return nil, err
	}
	return pbResp, nil
}

// protobuf implementation : no streaming for GetBoxes
func (s *GRPCServer) GetBoxes(ctx context.Context, req *pb.Empty) (*pb.Boxes, error) {
	_, resp, err := s.GetBoxesGRPCHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	pbResp, ok := resp.(*pb.Boxes)
	if !ok {
		err := fmt.Errorf("expecting *pb.Boxes received %T", resp)
		return nil, err
	}
	return pbResp, nil
}

// protobuf implementation : no streaming for Status
func (s *GRPCServer) Status(ctx context.Context, req *pb.Empty) (*pb.StatusMessage, error) {
	_, resp, err := s.StatusGRPCHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	pbResp, ok := resp.(*pb.StatusMessage)
	if !ok {
		err := fmt.Errorf("expecting *pb.StatusMessage received %T", resp)
		return nil, err
	}
	return pbResp, nil
}
