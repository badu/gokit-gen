package factory

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	pb "github.com/badu/gokit-gen/example/http/pb"
	kitHTTP "github.com/go-kit/kit/transport/http"
)

type httpClient struct {
	log              log.Logger
	MakeBoxEndpoint  endpoint.Endpoint
	GetBoxesEndpoint endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
}

func (c *httpClient) Log() log.Logger {
	return c.log
}

func (c *httpClient) MakeBox(ctx context.Context, name string, height int32, width int32, depth int32) (int32, int32, error) {
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

func (c *httpClient) GetBoxes(ctx context.Context, page int32, perPage int32) ([]BoxSpecification, error) {
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

func (c *httpClient) Status(ctx context.Context, page int32, perPage int32) (string, bool, error) {
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

func NewHTTPClient(serverAddr string, logger log.Logger, options ...kitHTTP.ClientOption) (Service, error) {
	urlMakeBox, err := url.Parse(serverAddr + "/v1/make-box")
	if err != nil {
		return nil, err
	}
	urlGetBoxes, err := url.Parse(serverAddr + "/v1/boxes")
	if err != nil {
		return nil, err
	}
	urlStatus, err := url.Parse(serverAddr + "/v1/status")
	if err != nil {
		return nil, err
	}
	result := &httpClient{
		log: logger,
		MakeBoxEndpoint: kitHTTP.NewClient(
			http.MethodPost,
			urlMakeBox,
			encodeMakeBoxHTTPRequest(logger),
			decodeMakeBoxHTTPResponse(logger),
			options...,
		).Endpoint(),
		GetBoxesEndpoint: kitHTTP.NewClient(
			http.MethodGet,
			urlGetBoxes,
			encodeGetBoxesHTTPRequest(logger),
			decodeGetBoxesHTTPResponse(logger),
			options...,
		).Endpoint(),
		StatusEndpoint: kitHTTP.NewClient(
			http.MethodGet,
			urlStatus,
			encodeStatusHTTPRequest(logger),
			decodeStatusHTTPResponse(logger),
			options...,
		).Endpoint(),
	}
	return result, nil
}

func encodeMakeBoxHTTPRequest(logger log.Logger) func(context.Context, *http.Request, interface{}) error {
	return func(ctx context.Context, req *http.Request, request interface{}) error {
		req.URL.Path = "/v1/make-box"
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(request); err != nil {
			level.Error(logger).Log("encodeMakeBoxHTTPRequest", err)
			return err
		}
		req.Body = ioutil.NopCloser(&buf)
		return nil
	}
}

func decodeMakeBoxHTTPResponse(logger log.Logger) func(context.Context, *http.Response) (interface{}, error) {
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, errors.New(string(body))
		}
		var response *pb.Empty
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			level.Error(logger).Log("decodeMakeBoxHTTPResponse", err)
			return nil, err
		}
		return NewEmptyFromPB(response), nil
	}
}

func encodeGetBoxesHTTPRequest(logger log.Logger) func(context.Context, *http.Request, interface{}) error {
	return func(ctx context.Context, req *http.Request, request interface{}) error {
		req.URL.Path = "/v1/boxes"
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(request); err != nil {
			level.Error(logger).Log("encodeGetBoxesHTTPRequest", err)
			return err
		}
		req.Body = ioutil.NopCloser(&buf)
		return nil
	}
}

func decodeGetBoxesHTTPResponse(logger log.Logger) func(context.Context, *http.Response) (interface{}, error) {
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, errors.New(string(body))
		}
		var response *pb.Boxes
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			level.Error(logger).Log("decodeGetBoxesHTTPResponse", err)
			return nil, err
		}
		return NewBoxesFromPB(response), nil
	}
}

func encodeStatusHTTPRequest(logger log.Logger) func(context.Context, *http.Request, interface{}) error {
	return func(ctx context.Context, req *http.Request, request interface{}) error {
		req.URL.Path = "/v1/status"
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(request); err != nil {
			level.Error(logger).Log("encodeStatusHTTPRequest", err)
			return err
		}
		req.Body = ioutil.NopCloser(&buf)
		return nil
	}
}

func decodeStatusHTTPResponse(logger log.Logger) func(context.Context, *http.Response) (interface{}, error) {
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, errors.New(string(body))
		}
		var response *pb.StatusMessage
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			level.Error(logger).Log("decodeStatusHTTPResponse", err)
			return nil, err
		}
		return NewStatusMessageFromPB(response), nil
	}
}
