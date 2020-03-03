package route_guide

import (
	"context"
	"fmt"
	"io"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

type ClientOption func(*grpcClient)
type ClientRequestFunc func(context.Context, *metadata.MD) context.Context
type ClientResponseFunc func(context.Context, metadata.MD, metadata.MD) context.Context

func ClientBefore(before ...ClientRequestFunc) ClientOption {
	return func(c *grpcClient) { c.before = append(c.before, before...) }
}

func ClientAfter(after ...ClientResponseFunc) ClientOption {
	return func(c *grpcClient) { c.after = append(c.after, after...) }
}

type grpcClient struct {
	log                        log.Logger
	GetFeatureEndpoint         endpoint.Endpoint
	receiveHalfDuplex          chan *Feature      // TODO : collect payloads from this channel
	broadcastReverseHalfDuplex chan *pb.Point     // TODO : send payloads to this channel
	receiveFullDuplex          chan *RouteNote    // TODO : collect payloads from this channel
	broadcastFullDuplex        chan *pb.RouteNote // TODO : send payloads to this channel
	directClient               pb.RouteGuideClient
	before                     []ClientRequestFunc
	after                      []ClientResponseFunc
}

func (c *grpcClient) Log() log.Logger {
	return c.log
}

func NewClient(conn *grpc.ClientConn, logger log.Logger, options ...kitGRPC.ClientOption) Service {
	result := &grpcClient{
		log:          logger,
		directClient: pb.NewRouteGuideClient(conn),
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
	return result
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

// usable : implementation of direct client for HalfDuplex - half duplex
func (c *grpcClient) CallHalfDuplex(extCtx context.Context, lo *Point, hi *Point) (string, *Point, error) {
	var err error
	ctx, cancel := context.WithCancel(extCtx)

	ctx = context.WithValue(ctx, kitGRPC.ContextKeyRequestMethod, "CallHalfDuplex")

	req := NewPBFromRectangle(&Rectangle{Lo: lo, Hi: hi})

	md := &metadata.MD{}
	for _, f := range c.before {
		ctx = f(ctx, md)
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	var header, trailer metadata.MD

	stream, err := c.directClient.HalfDuplex(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		c.log.Log("client_error", err)
		return "", nil, err
	}

	var closing = func() {
		cancel()
		stream.CloseSend()
		for _, f := range c.after {
			ctx = f(ctx, header, trailer)
		}
	}
	// receiving from server loop
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			closing()
			// read done.
			return "", nil, nil
		}
		if err != nil {
			// TODO : if server needs to close, with an error, which is a known error (no-error)
			//if err == CloseCommunication {
			//    closing()
			//    return "",nil, nil
			//}
			c.log.Log("client_error", fmt.Sprintf("server return error : %v\n", err))
			closing()
			return "", nil, err
		}

		domResp := NewFeatureFromPB(message)
		c.receiveHalfDuplex <- domResp // writing payloads to this channel, so dev can collect them
	}

	return "", nil, nil
}

// usable : implementation of direct client for ReverseHalfDuplex - reverse half duplex
func (c *grpcClient) CallReverseHalfDuplex(extCtx context.Context) error {
	var err error
	ctx := context.WithValue(extCtx, kitGRPC.ContextKeyRequestMethod, "CallReverseHalfDuplex")

	md := &metadata.MD{}
	for _, f := range c.before {
		ctx = f(ctx, md)
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	var header, trailer metadata.MD

	stream, err := c.directClient.ReverseHalfDuplex(ctx, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		c.log.Log("client_error", err)
		return err
	}

	var sendErr error
	// sending messages loop
	waitc := make(chan struct{})
	go func(grCtx context.Context) {
		for {
			select {
			case req := <-c.broadcastReverseHalfDuplex:
				// TODO : building a request as following :
				// req := NewPBFromPoint(&Point{ Latitude:latitude,Longitude:longitude, })
				// then write it to the channel:
				// c.broadcastReverseHalfDuplex <- req
				if sendErr = stream.Send(req); sendErr != nil {
					c.log.Log("client_error", fmt.Sprintf("Failed to send a request: %v\n", sendErr))
					close(waitc)
					return
				}

			case <-grCtx.Done():
				c.log.Log("client", "Context done.")
				close(waitc)
				return
			}
		}
	}(ctx)
	<-waitc // block here until context done or send error

	if sendErr == nil { // was context Done
		stream.CloseSend()
		for _, f := range c.after {
			ctx = f(ctx, header, trailer)
		}
	}
	return sendErr
}

// usable : implementation of direct client for FullDuplex - full duplex
func (c *grpcClient) CallFullDuplex(extCtx context.Context) error {
	var err error
	ctx, cancel := context.WithCancel(extCtx)

	ctx = context.WithValue(ctx, kitGRPC.ContextKeyRequestMethod, "CallFullDuplex")

	md := &metadata.MD{}
	for _, f := range c.before {
		ctx = f(ctx, md)
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	var header, trailer metadata.MD

	stream, err := c.directClient.FullDuplex(ctx, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		c.log.Log("client_error", err)
		return err
	}

	var closing = func() {
		cancel()
		stream.CloseSend()
		for _, f := range c.after {
			ctx = f(ctx, header, trailer)
		}
	}
	var sendErr error
	// sending messages loop
	go func(grCtx context.Context) {
		for {
			select {
			case req := <-c.broadcastFullDuplex:
				// TODO : building a request as following :
				// req := NewPBFromRouteNote(&RouteNote{ Location:location,Message:message, })
				// then write it to the channel:
				// c.broadcastFullDuplex <- req
				if sendErr = stream.Send(req); sendErr != nil {
					c.log.Log("client_error", fmt.Sprintf("Failed to send a request: %v\n", sendErr))
					return
				}

			case <-grCtx.Done():
				c.log.Log("client", "Context done.")
				return
			}
		}
	}(ctx)
	// receiving from server loop
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			closing()
			// read done.
			return nil
		}
		if err != nil {
			// TODO : if server needs to close, with an error, which is a known error (no-error)
			//if err == CloseCommunication {
			//    closing()
			//    return nil,"", nil
			//}
			c.log.Log("client_error", fmt.Sprintf("server return error : %v\n", err))
			closing()
			return err
		}
		if sendErr != nil {
			return sendErr
		}
		domResp := NewRouteNoteFromPB(message)
		c.receiveFullDuplex <- domResp // writing payloads to this channel, so dev can collect them
	}

	return nil
}
