package poc

import (
	"context"
	"errors"
	"fmt"
	"github.com/badu/gokit-gen/example/poc/pb"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io"
	"time"
)

var CloseCommunication = errors.New("communication end")

type Service interface {
	HalfDuplex(request *pb.Request, stream pb.Streams_HalfDuplexServer) error
	FullDuplex(stream pb.Streams_FullDuplexServer) error
	Logger() log.Logger
}

type InternalRequest struct {
	Msg string
}

type InternalReply struct {
	Ctx context.Context
	Msg string
}

type service struct {
	log       log.Logger
	ticker    *time.Ticker
	byeTicker *time.Ticker
}

func (s service) Logger() log.Logger {
	return s.log
}

func (s service) HalfDuplex(request *pb.Request, stream pb.Streams_HalfDuplexServer) error {
	waitc := make(chan struct{})
	// server sending loop. exit conditions : client closes or sending error
	var sendErr error
	go func(ctx context.Context, waitChannel chan struct{}) {
		for {
			select {
			case <-s.byeTicker.C:
				s.log.Log("service", "half duplex sending stop")
				s.ticker.Stop()
				s.byeTicker.Stop()
				close(waitc)
				return
			case tick := <-s.ticker.C:
				s.log.Log("service", "half duplex sending a new message")
				payload := &pb.Reply{Msg: fmt.Sprintf("A new message @ %v", tick)} // send a message to the client
				// send a message to the client
				if sendErr = stream.Send(payload); sendErr != nil {
					s.log.Log("service_error", sendErr)
					s.ticker.Stop()
					close(waitChannel)
					return
				}
			case <-ctx.Done():
				level.Debug(s.log).Log("service", "half duplex context done (client closed)")
				s.ticker.Stop()
				close(waitChannel)
				return
			}
		}
	}(stream.Context(), waitc)
	// send response to client : POC for using broadcast channel
	s.log.Log("service", "half duplex initial request : "+request.Msg)
	<-waitc // wait for channel to be closed
	return CloseCommunication
}

func (s service) FullDuplex(stream pb.Streams_FullDuplexServer) error {

	// server sending loop. exit conditions : client closes or sending error
	var sendErr error
	go func(ctx context.Context) {
		for {
			select {
			case <-s.byeTicker.C:
				s.log.Log("service", "full duplex sending stop")
				s.ticker.Stop()
				s.byeTicker.Stop()
				sendErr = CloseCommunication
				return
			case tick := <-s.ticker.C:
				s.log.Log("service", "full duplex sending a message")
				payload := &pb.Reply{Msg: fmt.Sprintf("A new message @ %v", tick)} // send a message to the client
				// send a message to the client
				if sendErr = stream.Send(payload); sendErr != nil {
					s.log.Log("service_error", sendErr)
					s.ticker.Stop()
					return
				}
			case <-ctx.Done():
				level.Debug(s.log).Log("service", "full duplex context done (client closed)")
				s.ticker.Stop()
				return
			}
		}
	}(stream.Context())

	// client receive loop : blocks until client closes or we have an error other than sending
	for {
		pbReq, err := stream.Recv()
		if err == io.EOF {
			level.Debug(s.log).Log("service", "full duplex client closed")
			return nil // no error, client just closed
		}
		if err != nil {
			s.log.Log("service_error", err)
			return err // receive error
		}
		if sendErr != nil {
			s.log.Log("service_error_dup", sendErr)
			return sendErr
		}

		s.log.Log("service", "---------------------------")
		s.log.Log("FULLDUPLEX", pbReq.Msg)
		s.log.Log("service", "---------------------------")
		// send response to client : POC for using broadcast channel
	}

	return sendErr
}

func NewService(log log.Logger) Service {
	return service{
		log:       log,
		ticker:    time.NewTicker(2 * time.Second),
		byeTicker: time.NewTicker(10 * time.Second),
	}
}
