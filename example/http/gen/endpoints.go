package factory

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log/level"
)

type Endpoints struct {
	MakeBoxEndpoint  endpoint.Endpoint
	GetBoxesEndpoint endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
}

func wrapEndpoint(ep endpoint.Endpoint, mwares []endpoint.Middleware) endpoint.Endpoint {
	for i := range mwares {
		ep = mwares[i](ep)
	}
	return ep
}

func MakeEndpoints(svc Service, mwares []endpoint.Middleware) Endpoints {
	return Endpoints{
		MakeBoxEndpoint:  wrapEndpoint(makeMakeBoxEndpoint(svc), mwares),
		GetBoxesEndpoint: wrapEndpoint(makeGetBoxesEndpoint(svc), mwares),
		StatusEndpoint:   wrapEndpoint(makeStatusEndpoint(svc), mwares),
	}
}
func makeMakeBoxEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		domReq, ok := req.(*BoxSpecification)
		if !ok {
			err := fmt.Errorf("expecting *MakeBoxRequest received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		page, perPage, err := svc.MakeBox(ctx, domReq.Name, domReq.Height, domReq.Width, domReq.Depth)
		if err != nil {
			return nil, err
		}
		return &Empty{Page: page, PerPage: perPage}, nil
	}
}

func makeGetBoxesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		domReq, ok := req.(*Empty)
		if !ok {
			err := fmt.Errorf("expecting *GetBoxesRequest received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		boxes, err := svc.GetBoxes(ctx, domReq.Page, domReq.PerPage)
		if err != nil {
			return nil, err
		}
		return &Boxes{Boxes: boxes}, nil
	}
}

func makeStatusEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		domReq, ok := req.(*Empty)
		if !ok {
			err := fmt.Errorf("expecting *StatusRequest received %T", req)
			level.Error(svc.Log()).Log("endpoint_error", err)
			return nil, err
		}
		serviceName, ok, err := svc.Status(ctx, domReq.Page, domReq.PerPage)
		if err != nil {
			return nil, err
		}
		return &StatusMessage{ServiceName: serviceName, Ok: ok}, nil
	}
}
