package poc

import (
	"context"
	"fmt"
	"github.com/badu/gokit-gen/example/poc/pb"
	"github.com/go-kit/kit/log"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

type Client struct {
	kitGRPC.Client
	log    log.Logger
	client pb.StreamsClient
	enc    kitGRPC.EncodeRequestFunc
	dec    kitGRPC.DecodeResponseFunc
	before []kitGRPC.ClientRequestFunc
	after  []kitGRPC.ClientResponseFunc
}

func (c Client) Logger() log.Logger {
	return c.log
}

func (c Client) CallHalfDuplex(extCtx context.Context, request *InternalRequest) error {
	var err error
	ctx, cancel := context.WithCancel(extCtx)

	ctx = context.WithValue(ctx, kitGRPC.ContextKeyRequestMethod, "HalfDuplex")

	req, err := c.enc(ctx, request)
	if err != nil {
		return err
	}

	md := &metadata.MD{}
	for _, f := range c.before {
		ctx = f(ctx, md)
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	var header, trailer metadata.MD

	stream, err := c.client.HalfDuplex(ctx, req.(*pb.Request), grpc.Header(&header), grpc.Trailer(&trailer))
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

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			closing()
			// read done.
			return nil
		}
		if err != nil {
			if err == CloseCommunication {
				closing()
				return nil
			}
			c.log.Log("client_error", err)
			closing()
			return err
		}

		response, err := c.dec(ctx, message)
		if err != nil {
			c.log.Log("client_error", "Failed to decode : %v\n", err)
			closing()
			return err
		}
		domResp, ok := response.(*InternalReply)
		if ok {
			c.log.Log("client", "---------------------------")
			c.log.Log("HALFDUPLEX", domResp.Msg)
			c.log.Log("client", "---------------------------")
		} else {
			c.log.Log("client_error", "expecting *InternalReply")
		}
	}

	return nil
}

func (c Client) CallFullDuplex(cctx context.Context) error {
	var err error
	ctx, cancel := context.WithCancel(cctx)

	ctx = context.WithValue(ctx, kitGRPC.ContextKeyRequestMethod, "FullDuplex")

	md := &metadata.MD{}
	for _, f := range c.before {
		ctx = f(ctx, md)
	}
	ctx = metadata.NewOutgoingContext(ctx, *md)

	var header, trailer metadata.MD

	stream, err := c.client.FullDuplex(ctx, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		c.log.Log("client_error", err)
		return err
	}

	// sending messages loop
	streamTicker := time.NewTicker(1500 * time.Millisecond)
	go func() {
		for {
			select {
			case tick := <-streamTicker.C:
				request := &InternalRequest{Msg: fmt.Sprintf("sending a new message to server @ %v", tick)}
				c.log.Log("client", "sending new message via stream")
				req, err := c.enc(ctx, request)
				if err != nil {
					c.log.Log("client_error", err)
					streamTicker.Stop()
					return
				}

				if err := stream.Send(req.(*pb.Request)); err != nil {
					c.log.Log("client_error", err)
					streamTicker.Stop()
					return
				}

			case <-ctx.Done():
				c.log.Log("client_done", "Context done.")
				streamTicker.Stop()
				return
			}
		}
	}()

	var closing = func() {
		cancel()
		stream.CloseSend()
		for _, f := range c.after {
			ctx = f(ctx, header, trailer)
		}
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			closing()
			// read done.
			return nil
		}
		if err != nil {
			if err == CloseCommunication {
				closing()
				return nil
			}
			c.log.Log("client_error", err)
			closing()
			return err
		}

		response, err := c.dec(ctx, message)
		if err != nil {
			c.log.Log("client_error", err)
			closing()
			return err
		}
		domResp, ok := response.(*InternalReply)
		if ok {
			c.log.Log("client", "---------------------------")
			c.log.Log("FULLDUPLEX", domResp.Msg)
			c.log.Log("client", "---------------------------")
		} else {
			c.log.Log("client_error", "expecting *InternalReply")
		}
	}

	return nil
}

func NewDirectClient(
	log log.Logger,
	conn *grpc.ClientConn,
	enc kitGRPC.EncodeRequestFunc,
	dec kitGRPC.DecodeResponseFunc) *Client {

	c := &Client{
		log:    log,
		client: pb.NewStreamsClient(conn),
		enc:    enc,
		dec:    dec,
	}

	return c
}

func (c *Client) ApplyExtraOptions(options ...ClientOption) {
	for _, option := range options {
		option(c)
	}
}

type ClientOption func(*Client)

func ClientBefore(before ...kitGRPC.ClientRequestFunc) ClientOption {
	return func(c *Client) { c.before = append(c.before, before...) }
}

func ClientAfter(after ...kitGRPC.ClientResponseFunc) ClientOption {
	return func(c *Client) { c.after = append(c.after, after...) }
}
