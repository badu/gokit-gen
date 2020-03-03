package poc

import (
	"context"
	"errors"
	"github.com/badu/gokit-gen/example/poc/pb"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	kitGRPC "github.com/go-kit/kit/transport/grpc"
)

func makeHalfDuplexEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*RequestAndStreamHalfDuplex)
		if !ok {
			err := errors.New("expecting *RequestAndStreamHalfDuplex")
			level.Error(svc.Logger()).Log("SRV_ERROR", err)
			return nil, err
		}
		err := svc.HalfDuplex(req.Request, req.Stream)
		if err != nil {
			level.Error(svc.Logger()).Log("SRV_ERROR", err)
			return nil, err
		}
		return nil, err
	}
}

func makeFullDuplexEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(pb.Streams_FullDuplexServer)
		if !ok {
			err := errors.New("expecting pb.Streams_FullDuplexServer")
			level.Error(svc.Logger()).Log("SRV_ERROR", err)
			return nil, err
		}
		err := svc.FullDuplex(req)
		if err != nil {
			level.Error(svc.Logger()).Log("SRV_ERROR", err)
			return nil, err
		}
		return nil, err
	}
}

type serverBinding struct {
	log        log.Logger
	halfDuplex kitGRPC.Handler
	fullDuplex kitGRPC.Handler
}

// request and stream for half duplex
type RequestAndStreamHalfDuplex struct {
	Request *pb.Request
	Stream  pb.Streams_HalfDuplexServer
}

func (b *serverBinding) HalfDuplex(req *pb.Request, stream pb.Streams_HalfDuplexServer) error {
	b.log.Log("SRV_HALFDUPLEX", "making internal object")
	reqNStream := &RequestAndStreamHalfDuplex{Request: req, Stream: stream}
	_, _, err := b.halfDuplex.ServeGRPC(stream.Context(), reqNStream)
	if err != nil {
		return err
	}
	return nil
}

func (b *serverBinding) FullDuplex(stream pb.Streams_FullDuplexServer) error {
	b.log.Log("SRV_HALFDUPLEX", "passing stream")
	_, _, err := b.fullDuplex.ServeGRPC(stream.Context(), stream)
	if err != nil {
		return err
	}
	return nil
}

func NewBinding(svc Service) *serverBinding {
	options := []kitGRPC.ServerOption{
		kitGRPC.ServerErrorHandler(
			transport.NewLogErrorHandler(svc.Logger()),
		),
		kitGRPC.ServerBefore(
			extractCorrelationID(svc.Logger()),
			displayServerRequestHeaders(svc.Logger()),
		),
		kitGRPC.ServerAfter(
			injectResponseHeader(svc.Logger()),
			injectResponseTrailer(svc.Logger()),
			injectConsumedCorrelationID(svc.Logger()),
			displayServerResponseHeaders(svc.Logger()),
			displayServerResponseTrailers(svc.Logger()),
		),
	}
	return &serverBinding{
		log: svc.Logger(),
		halfDuplex: kitGRPC.NewServer(
			makeHalfDuplexEndpoint(svc),
			DecodeRequestFromClient(svc.Logger()),
			EncodeResponseToClient(svc.Logger()),
			options...,
		),
		fullDuplex: kitGRPC.NewServer(
			makeFullDuplexEndpoint(svc),
			DecodeRequestFromClient(svc.Logger()),
			EncodeResponseToClient(svc.Logger()),
			options...,
		),
	}
}
