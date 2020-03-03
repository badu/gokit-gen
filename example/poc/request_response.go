package poc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/badu/gokit-gen/example/poc/pb"
)

func DecodeRequestFromClient(log log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		return req, nil
	}
}

func EncodeResponseToClient(log log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(_ context.Context, resp interface{}) (interface{}, error) {
		return resp, nil
	}
}

func EncodeRequestToServer(log log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		r, ok := req.(*InternalRequest)
		if !ok {
			err := errors.New("expecting *InternalRequest")
			level.Error(log).Log("CLIENT_ENCODE", err)
			return nil, err
		}
		level.Debug(log).Log("CLIENT_ENCODE", "", "message", r.Msg)
		return &pb.Request{Msg: r.Msg}, nil
	}
}
func DecodeResponseFromServer(log log.Logger) func(context.Context, interface{}) (interface{}, error) {
	return func(ctx context.Context, resp interface{}) (interface{}, error) {
		r, ok := resp.(*pb.Reply)
		if !ok {
			err := errors.New("expecting *pb.Reply")
			level.Error(log).Log("CLIENT_DECODE", err)
			return nil, err
		}
		level.Debug(log).Log("CLIENT_DECODE", "", "message", r.Msg)
		return &InternalReply{Msg: r.Msg, Ctx: ctx}, nil
	}
}
