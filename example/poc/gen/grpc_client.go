package example

import (
	"context"
	"fmt"
	"io"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/badu/gokit-gen/example/poc/pb"
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
	log                 log.Logger
	receiveHalfDuplex   chan *Reply      // TODO : collect payloads from this channel
	receiveFullDuplex   chan *Reply      // TODO : collect payloads from this channel
	broadcastFullDuplex chan *pb.Request // TODO : send payloads to this channel
	directClient        pb.StreamsClient
	before              []ClientRequestFunc
	after               []ClientResponseFunc
}

func (c *grpcClient) Log() log.Logger {
	return c.log
}

func (c *grpcClient) ApplyExtraOptions(options ...ClientOption) {
	for _, option := range options {
		option(c)
	}
}

// getter for receiveHalfDuplex chan Reply
func (c *grpcClient) ReceiveHalfDuplex() chan *Reply {
	return c.receiveHalfDuplex
}

// getter for receiveFullDuplex chan Reply
func (c *grpcClient) ReceiveFullDuplex() chan *Reply {
	return c.receiveFullDuplex
}

// getter for broadcastFullDuplex chan *pb.Request
func (c *grpcClient) BroadcastFullDuplex() chan *pb.Request {
	return c.broadcastFullDuplex
}

func NewClient(conn *grpc.ClientConn, logger log.Logger, options ...kitGRPC.ClientOption) *grpcClient {
	result := &grpcClient{
		log:                 logger,
		directClient:        pb.NewStreamsClient(conn),
		receiveHalfDuplex:   make(chan *Reply),
		receiveFullDuplex:   make(chan *Reply),
		broadcastFullDuplex: make(chan *pb.Request),
	}
	return result
}

// unusable : client has to implement Service interface for half duplex HalfDuplex
func (c *grpcClient) HalfDuplex(req *pb.Request, stream pb.Streams_HalfDuplexServer) error {
	return nil
}

// unusable : client has to implement Service interface for full duplex FullDuplex
func (c *grpcClient) FullDuplex(stream pb.Streams_FullDuplexServer) error {
	return nil
}

// usable : implementation of direct client for HalfDuplex - half duplex
func (c *grpcClient) CallHalfDuplex(extCtx context.Context, msg string) error {
	var err error
	ctx, cancel := context.WithCancel(extCtx)

	ctx = context.WithValue(ctx, kitGRPC.ContextKeyRequestMethod, "CallHalfDuplex")

	req := NewPBFromRequest(&Request{Msg: msg})

	md := &metadata.MD{}
	for _, f := range c.before {
		ctx = f(ctx, md)
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	var header, trailer metadata.MD

	stream, err := c.directClient.HalfDuplex(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
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
			//    return "", nil
			//}
			c.log.Log("client_error", fmt.Sprintf("server return error : %v\n", err))
			closing()
			return err
		}

		domResp := NewReplyFromPB(message)
		c.receiveHalfDuplex <- domResp // writing payloads to this channel, so dev can collect them
	}

	return nil
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
			case <-grCtx.Done():
				c.log.Log("client", "Context done.")
				return
			case req := <-c.broadcastFullDuplex:
				// TODO : building a request as following :
				// req := NewPBFromRequest(&Request{ Msg:msg, })
				// then write it to the channel:
				// c.broadcastFullDuplex <- req
				if sendErr = stream.Send(req); sendErr != nil {
					c.log.Log("client_error", fmt.Sprintf("Failed to send a request: %v\n", sendErr))
					return
				}
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
			//    return "", nil
			//}
			c.log.Log("client_error", fmt.Sprintf("server return error : %v\n", err))
			closing()
			return err
		}
		if sendErr != nil {
			return sendErr
		}
		domResp := NewReplyFromPB(message)
		c.receiveFullDuplex <- domResp // writing payloads to this channel, so dev can collect them
	}

	return nil
}
