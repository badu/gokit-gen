package example

import (
	"context"
	"io"

	pb "github.com/badu/gokit-gen/example/poc/pb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	BroadcastHalfDuplex() chan Reply
	HalfDuplex(req *pb.Request, stream pb.Streams_HalfDuplexServer) error // half duplex (client request, server streams)
	BroadcastFullDuplex() chan Reply
	FullDuplex(stream pb.Streams_FullDuplexServer) error // full duplex (both client and server streams)
	Log() log.Logger
}

type serviceImpl struct {
	log                 log.Logger
	repo                Repository
	broadcastHalfDuplex chan Reply // channel to stream to clients
	broadcastFullDuplex chan Reply // channel to stream to clients
}

func NewService(log log.Logger, repo Repository) Service {
	return &serviceImpl{log: log, repo: repo, broadcastHalfDuplex: make(chan Reply), broadcastFullDuplex: make(chan Reply)}
}

func (s *serviceImpl) Log() log.Logger {
	return s.log
}

// getter for broadcastHalfDuplex chan Reply
func (s *serviceImpl) BroadcastHalfDuplex() chan Reply {
	return s.broadcastHalfDuplex
}

// getter for broadcastFullDuplex chan Reply
func (s *serviceImpl) BroadcastFullDuplex() chan Reply {
	return s.broadcastFullDuplex
}

// half duplex for HalfDuplex
func (s *serviceImpl) HalfDuplex(pbReq *pb.Request, stream pb.Streams_HalfDuplexServer) error {
	waitc := make(chan struct{})
	// server sending loop. exit conditions : client closes or sending error
	var sendErr error
	go func(ctx context.Context, waitChannel chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				level.Debug(s.log).Log("HalfDuplexService", "Context done (client closed)")
				close(waitChannel)
				return
			case message := <-s.broadcastHalfDuplex:
				level.Debug(s.log).Log("HalfDuplexService", "Sending payload")
				payload := NewPBFromReply(&message)
				// send a message to the client
				if sendErr = stream.Send(payload); sendErr != nil {
					level.Error(s.log).Log("send_error", sendErr)
					close(waitChannel)
					return
				}
			}
		}
	}(stream.Context(), waitc)
	// send response to client : POC for using broadcast channel
	req := NewRequestFromPB(pbReq)
	resMsg, err := s.repo.HalfDuplex(stream.Context(), req.Msg)
	if err != nil {
		return err
	}
	resp := Reply{Msg: resMsg}
	s.broadcastHalfDuplex <- resp
	// wait until client has closed or we're having a sending error
	<-waitc
	return sendErr
}

// full duplex for FullDuplex
func (s *serviceImpl) FullDuplex(stream pb.Streams_FullDuplexServer) error {
	// server sending loop. exit conditions : client closes or sending error
	var sendErr error
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				level.Debug(s.log).Log("FullDuplexService", "Context done (client closed)")
				return
			case message := <-s.broadcastFullDuplex:
				level.Debug(s.log).Log("FullDuplexService", "Sending payload")
				payload := NewPBFromReply(&message)
				// send a message to the client
				if sendErr = stream.Send(payload); sendErr != nil {
					level.Error(s.log).Log("send_error", sendErr)
					return
				}
			}
		}
	}(stream.Context())
	// client receive loop : blocks until client closes or we have an error other than sending
	for {
		pbReq, err := stream.Recv()
		if err == io.EOF {
			level.Debug(s.log).Log("FullDuplexService", "client closed")
			return nil // no error, client just closed
		}
		if err != nil {
			level.Error(s.log).Log("receive_error", err)
			return err // receive error
		}
		// checking if sending error has happen previously
		if sendErr != nil {
			return sendErr
		}

		req := NewRequestFromPB(pbReq)
		resMsg, err := s.repo.FullDuplex(stream.Context(), req.Msg)
		if err != nil {
			return err
		}
		resp := Reply{Msg: resMsg}
		s.broadcastFullDuplex <- resp
	}
	// if we had a previous sending error, we're returning it
	return sendErr
}
