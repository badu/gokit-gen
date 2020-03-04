package example

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log/level"

	pb "github.com/badu/gokit-gen/example/poc/pb"
)

type Endpoints struct {
	HalfDuplexEndpoint endpoint.Endpoint // half duplex
	FullDuplexEndpoint endpoint.Endpoint //  full duplex
}

func wrapEndpoint(ep endpoint.Endpoint, mwares []endpoint.Middleware) endpoint.Endpoint {
	for i := range mwares {
		ep = mwares[i](ep)
	}
	return ep
}

func MakeEndpoints(svc Service, mwares []endpoint.Middleware) Endpoints {
	return Endpoints{
		HalfDuplexEndpoint: wrapEndpoint(makeHalfDuplexEndpoint(svc), mwares),
		FullDuplexEndpoint: wrapEndpoint(makeFullDuplexEndpoint(svc), mwares),
	}
}

// half duplex
func makeHalfDuplexEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqAndStream, ok := req.(*RequestAndStreamHalfDuplex)
		if !ok {
			err := fmt.Errorf("expecting *RequestAndStreamHalfDuplex received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		err := svc.HalfDuplex(reqAndStream.Request, reqAndStream.Stream)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// full duplex
func makeFullDuplexEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		stream, ok := req.(pb.Streams_FullDuplexServer)
		if !ok {
			err := fmt.Errorf("expecting pb.Streams_FullDuplexServer received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		err := svc.FullDuplex(stream)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
