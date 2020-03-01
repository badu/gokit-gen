package route_guide

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type grpcClient struct {
	log                log.Logger
	GetFeatureEndpoint endpoint.Endpoint
	//HalfDuplexEndpoint endpoint.Endpoint // half duplex
	//ReverseHalfDuplexEndpoint endpoint.Endpoint // reverse half duplex
	//FullDuplexEndpoint endpoint.Endpoint  // full duplex
}

func (c *grpcClient) Log() log.Logger {
	return c.log
}

func NewClient(conn *grpc.ClientConn, logger log.Logger, options ...kitGRPC.ClientOption) Service {
	return &grpcClient{
		log: logger,

		GetFeatureEndpoint: kitGRPC.NewClient(
			conn,
			"routeguide.RouteGuideService",
			"GetFeature",
			encodeGetFeatureGRPCRequest(logger),
			decodeGetFeatureGRPCResponse(logger),
			&pb.Feature{},
			options...,
		).Endpoint(),
	}
}

func encodeGetFeatureGRPCRequest(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		pbReq, ok := req.(*Point)
		if !ok {
			err := fmt.Errorf("expecting *PointRequest received %T", req)
			level.Error(logger).Log("decodeGetFeatureGRPCRequest", err)
			return nil, err
		}
		return NewPBFromPoint(pbReq), nil
	}
}

func decodeGetFeatureGRPCResponse(logger log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		domResp, ok := resp.(*pb.Feature)
		if !ok {
			err := fmt.Errorf("expecting *pb.Feature received %T", resp)
			level.Error(logger).Log("encodeGetFeatureGRPCRequest", err)
			return nil, err
		}
		return NewFeatureFromPB(domResp), nil
	}
}

func (c *grpcClient) GetFeature(ctx context.Context, latitude int32, longitude int32) (string, *Point, error) {
	resp, err := c.GetFeatureEndpoint(ctx, &Point{Latitude: latitude, Longitude: longitude})
	if err != nil {
		return "", nil, err
	}
	domResp, ok := resp.(*Feature)
	if !ok {
		err := fmt.Errorf("expecting *Feature received %T", resp)
		return "", nil, err
	}
	return domResp.Name, domResp.Location, nil
}

// unusable : client has to implement Service interface for half duplex HalfDuplex
func (c *grpcClient) HalfDuplex(req *pb.Rectangle, stream pb.RouteGuide_HalfDuplexServer) error {
	return nil
}

// unusable : client has to implement Service interface for reverse half duplex ReverseHalfDuplex
func (c *grpcClient) ReverseHalfDuplex(stream pb.RouteGuide_ReverseHalfDuplexServer) error {
	return nil
}

// unusable : client has to implement Service interface for full duplex FullDuplex
func (c *grpcClient) FullDuplex(stream pb.RouteGuide_FullDuplexServer) error {
	return nil
}
