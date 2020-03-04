package factory

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"

	pb "github.com/badu/gokit-gen/example/http/pb"
	"github.com/go-kit/kit/endpoint"
	kitHTTP "github.com/go-kit/kit/transport/http"
)

func MakeMakeBoxHTTPHandler(ep endpoint.Endpoint, logger log.Logger, options ...kitHTTP.ServerOption) *kitHTTP.Server {
	return kitHTTP.NewServer(
		ep,
		decodeMakeBoxHTTPRequest(logger),
		encodeMakeBoxHTTPResponse(logger),
		options...,
	)
}

func decodeMakeBoxHTTPRequest(logger log.Logger) func(context.Context, *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var request pb.BoxSpecification
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			level.Debug(logger).Log("json_decode_error", err)
			return nil, err
		}
		return NewBoxSpecificationFromPB(&request), nil
	}
}

func encodeMakeBoxHTTPResponse(logger log.Logger) func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
		result, ok := resp.(*Empty)
		if !ok {
			err := fmt.Errorf("expecting *Empty received %T", resp)
			level.Error(logger).Log("encodeMakeBoxHTTPResponse", err)
			return err
		}
		payload := NewPBFromEmpty(result)
		return json.NewEncoder(w).Encode(payload)
	}
}

func MakeGetBoxesHTTPHandler(ep endpoint.Endpoint, logger log.Logger, options ...kitHTTP.ServerOption) *kitHTTP.Server {
	return kitHTTP.NewServer(
		ep,
		decodeGetBoxesHTTPRequest(logger),
		encodeGetBoxesHTTPResponse(logger),
		options...,
	)
}

func decodeGetBoxesHTTPRequest(logger log.Logger) func(context.Context, *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var request pb.Empty
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			level.Debug(logger).Log("json_decode_error", err)
			return nil, err
		}
		return NewEmptyFromPB(&request), nil
	}
}

func encodeGetBoxesHTTPResponse(logger log.Logger) func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
		result, ok := resp.(*Boxes)
		if !ok {
			err := fmt.Errorf("expecting *Boxes received %T", resp)
			level.Error(logger).Log("encodeGetBoxesHTTPResponse", err)
			return err
		}
		payload := NewPBFromBoxes(result)
		return json.NewEncoder(w).Encode(payload)
	}
}

func MakeStatusHTTPHandler(ep endpoint.Endpoint, logger log.Logger, options ...kitHTTP.ServerOption) *kitHTTP.Server {
	return kitHTTP.NewServer(
		ep,
		decodeStatusHTTPRequest(logger),
		encodeStatusHTTPResponse(logger),
		options...,
	)
}

func decodeStatusHTTPRequest(logger log.Logger) func(context.Context, *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var request pb.Empty
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			level.Debug(logger).Log("json_decode_error", err)
			return nil, err
		}
		return NewEmptyFromPB(&request), nil
	}
}

func encodeStatusHTTPResponse(logger log.Logger) func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
		result, ok := resp.(*StatusMessage)
		if !ok {
			err := fmt.Errorf("expecting *StatusMessage received %T", resp)
			level.Error(logger).Log("encodeStatusHTTPResponse", err)
			return err
		}
		payload := NewPBFromStatusMessage(result)
		return json.NewEncoder(w).Encode(payload)
	}
}

func RegisterHandlers(router *mux.Router, endpoints Endpoints, logger log.Logger, options ...kitHTTP.ServerOption) {
	// MakeBox
	router.Methods(http.MethodPost).Path("/v1/make-box").Handler(MakeMakeBoxHTTPHandler(endpoints.MakeBoxEndpoint, logger, options...))
	router.Methods(http.MethodOptions).Path("/v1/make-box").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // CORS Options
	})

	// GetBoxes
	router.Methods(http.MethodGet).Path("/v1/boxes").Handler(MakeGetBoxesHTTPHandler(endpoints.GetBoxesEndpoint, logger, options...))
	router.Methods(http.MethodOptions).Path("/v1/boxes").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // CORS Options
	})

	// Status
	router.Methods(http.MethodGet).Path("/v1/status").Handler(MakeStatusHTTPHandler(endpoints.StatusEndpoint, logger, options...))
	router.Methods(http.MethodOptions).Path("/v1/status").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // CORS Options
	})

}
