package route_guide

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log/level"

	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
)

type Endpoints struct {
	GetFeatureEndpoint        endpoint.Endpoint
	HalfDuplexEndpoint        endpoint.Endpoint // half duplex
	ReverseHalfDuplexEndpoint endpoint.Endpoint // reverse half duplex
	FullDuplexEndpoint        endpoint.Endpoint //  full duplex
}

func wrapEndpoint(ep endpoint.Endpoint, mwares []endpoint.Middleware) endpoint.Endpoint {
	for i := range mwares {
		ep = mwares[i](ep)
	}
	return ep
}

func MakeEndpoints(svc Service, mwares []endpoint.Middleware) Endpoints {
	return Endpoints{
		GetFeatureEndpoint:        wrapEndpoint(makeGetFeatureEndpoint(svc), mwares),
		HalfDuplexEndpoint:        wrapEndpoint(makeHalfDuplexEndpoint(svc), mwares),
		ReverseHalfDuplexEndpoint: wrapEndpoint(makeReverseHalfDuplexEndpoint(svc), mwares),
		FullDuplexEndpoint:        wrapEndpoint(makeFullDuplexEndpoint(svc), mwares),
	}
}
func makeGetFeatureEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		domReq, ok := req.(*Point)
		if !ok {
			err := fmt.Errorf("expecting *GetFeatureRequest received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		name, location, err := svc.GetFeature(ctx, domReq.Latitude, domReq.Longitude)
		if err != nil {
			return nil, err
		}
		return &Feature{Name: name, Location: location}, nil
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

// reverse half duplex
func makeReverseHalfDuplexEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		stream, ok := req.(pb.RouteGuide_ReverseHalfDuplexServer)
		if !ok {
			err := fmt.Errorf("expecting pb.RouteGuide_ReverseHalfDuplexServer received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		err := svc.ReverseHalfDuplex(stream)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

// full duplex
func makeFullDuplexEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		stream, ok := req.(pb.RouteGuide_FullDuplexServer)
		if !ok {
			err := fmt.Errorf("expecting pb.RouteGuide_FullDuplexServer received %T", req)
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
