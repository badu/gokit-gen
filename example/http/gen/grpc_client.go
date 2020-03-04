package factory

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	pb "github.com/badu/gokit-gen/example/http/pb"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type grpcClient struct {
	log              log.Logger
	MakeBoxEndpoint  endpoint.Endpoint
	GetBoxesEndpoint endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
}

func (c *grpcClient) Log() log.Logger {
	return c.log
}

func NewClient(conn *grpc.ClientConn, logger log.Logger, options ...kitGRPC.ClientOption) *grpcClient {
	result := &grpcClient{
		log: logger,
		MakeBoxEndpoint: kitGRPC.NewClient(
			conn,
			"pb.BoxFactory",
			"MakeBox",
			encodeMakeBoxGRPCRequest(logger),
			decodeMakeBoxGRPCResponse(logger),
			&pb.Empty{},
			options...,
		).Endpoint(),
		GetBoxesEndpoint: kitGRPC.NewClient(
			conn,
			"pb.BoxFactory",
			"GetBoxes",
			encodeGetBoxesGRPCRequest(logger),
			decodeGetBoxesGRPCResponse(logger),
			&pb.Boxes{},
			options...,
		).Endpoint(),
		StatusEndpoint: kitGRPC.NewClient(
			conn,
			"pb.BoxFactory",
			"Status",
			encodeStatusGRPCRequest(logger),
			decodeStatusGRPCResponse(logger),
			&pb.StatusMessage{},
			options...,
		).Endpoint(),
	}
	return result
}

func encodeMakeBoxGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*BoxSpecification)
		if !ok {
			err := fmt.Errorf("expecting *BoxSpecificationRequest received %T", req)
			level.Error(logger).Log("decodeMakeBoxGRPCRequest", err)
			return nil, err
		}
		return NewPBFromBoxSpecification(pbReq), nil
	}
}

func decodeMakeBoxGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*pb.Empty)
		if !ok {
			err := fmt.Errorf("expecting *pb.Empty received %T", resp)
			level.Error(logger).Log("encodeMakeBoxGRPCRequest", err)
			return nil, err
		}
		return NewEmptyFromPB(domResp), nil
	}
}

func (c *grpcClient) MakeBox(ctx context.Context, name string, height int32, width int32, depth int32) (int32, int32, error) {
	resp, err := c.MakeBoxEndpoint(ctx, &BoxSpecification{Name: name, Height: height, Width: width, Depth: depth})
	if err != nil {
		return 0, 0, err
	}
	domResp, ok := resp.(*Empty)
	if !ok {
		err := fmt.Errorf("expecting *Empty received %T", resp)
		return 0, 0, err
	}
	return domResp.Page, domResp.PerPage, nil
}

func encodeGetBoxesGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*Empty)
		if !ok {
			err := fmt.Errorf("expecting *EmptyRequest received %T", req)
			level.Error(logger).Log("decodeGetBoxesGRPCRequest", err)
			return nil, err
		}
		return NewPBFromEmpty(pbReq), nil
	}
}

func decodeGetBoxesGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*pb.Boxes)
		if !ok {
			err := fmt.Errorf("expecting *pb.Boxes received %T", resp)
			level.Error(logger).Log("encodeGetBoxesGRPCRequest", err)
			return nil, err
		}
		return NewBoxesFromPB(domResp), nil
	}
}

func (c *grpcClient) GetBoxes(ctx context.Context, page int32, perPage int32) ([]BoxSpecification, error) {
	resp, err := c.GetBoxesEndpoint(ctx, &Empty{Page: page, PerPage: perPage})
	if err != nil {
		return nil, err
	}
	domResp, ok := resp.(*Boxes)
	if !ok {
		err := fmt.Errorf("expecting *Boxes received %T", resp)
		return nil, err
	}
	return domResp.Boxes, nil
}

func encodeStatusGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*Empty)
		if !ok {
			err := fmt.Errorf("expecting *EmptyRequest received %T", req)
			level.Error(logger).Log("decodeStatusGRPCRequest", err)
			return nil, err
		}
		return NewPBFromEmpty(pbReq), nil
	}
}

func decodeStatusGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*pb.StatusMessage)
		if !ok {
			err := fmt.Errorf("expecting *pb.StatusMessage received %T", resp)
			level.Error(logger).Log("encodeStatusGRPCRequest", err)
			return nil, err
		}
		return NewStatusMessageFromPB(domResp), nil
	}
}

func (c *grpcClient) Status(ctx context.Context, page int32, perPage int32) (string, bool, error) {
	resp, err := c.StatusEndpoint(ctx, &Empty{Page: page, PerPage: perPage})
	if err != nil {
		return "", false, err
	}
	domResp, ok := resp.(*StatusMessage)
	if !ok {
		err := fmt.Errorf("expecting *StatusMessage received %T", resp)
		return "", false, err
	}
	return domResp.ServiceName, domResp.Ok, nil
}
