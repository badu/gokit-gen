package route_guide

import (
	"context"
	"io"

	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	GetFeature(ctx context.Context, latitude int32, longitude int32) (string, *Point, error)
	BroadcastHalfDuplex() chan Feature
	HalfDuplex(req *pb.Rectangle, stream pb.RouteGuide_HalfDuplexServer) error // half duplex (client request, server streams)
	ReverseHalfDuplex(stream pb.RouteGuide_ReverseHalfDuplexServer) error      // reverse half duplex (client streams, server responds)
	BroadcastFullDuplex() chan RouteNote
	FullDuplex(stream pb.RouteGuide_FullDuplexServer) error // full duplex (both client and server streams)
	Log() log.Logger
}

type serviceImpl struct {
	log                 log.Logger
	repo                Repository
	broadcastHalfDuplex chan Feature   // channel to stream to clients
	broadcastFullDuplex chan RouteNote // channel to stream to clients
}

func NewService(log log.Logger, repo Repository) Service {
	return &serviceImpl{log: log, repo: repo, broadcastHalfDuplex: make(chan Feature), broadcastFullDuplex: make(chan RouteNote)}
}

func (s *serviceImpl) Log() log.Logger {
	return s.log
}

// getter for broadcastHalfDuplex chan Feature
func (s *serviceImpl) BroadcastHalfDuplex() chan Feature {
	return s.broadcastHalfDuplex
}

// getter for broadcastFullDuplex chan RouteNote
func (s *serviceImpl) BroadcastFullDuplex() chan RouteNote {
	return s.broadcastFullDuplex
}

func (s *serviceImpl) GetFeature(ctx context.Context, latitude int32, longitude int32) (string, *Point, error) {
	resName, resLocation, err := s.repo.GetFeature(ctx, latitude, longitude)
	if err != nil {
		return "", nil, err
	}
	return resName, resLocation, nil
}

// half duplex for HalfDuplex
func (s *serviceImpl) HalfDuplex(pbReq *pb.Rectangle, stream pb.RouteGuide_HalfDuplexServer) error {
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
				payload := NewPBFromFeature(&message)
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
	req := NewRectangleFromPB(pbReq)
	resName, resLocation, err := s.repo.HalfDuplex(stream.Context(), req.Lo, req.Hi)
	if err != nil {
		return err
	}
	resp := Feature{Name: resName, Location: resLocation}
	s.broadcastHalfDuplex <- resp
	// wait until client has closed or we're having a sending error
	<-waitc
	return sendErr
}

// reverse half duplex for ReverseHalfDuplex
func (s *serviceImpl) ReverseHalfDuplex(stream pb.RouteGuide_ReverseHalfDuplexServer) error {
	// client receive loop : blocks until client closes or we have an error other than sending
	for {
		pbReq, err := stream.Recv()
		if err == io.EOF {
			level.Debug(s.log).Log("ReverseHalfDuplexService", "client closed")
			// TODO : make payload to be returned
			sendErr := stream.SendAndClose(&pb.RouteSummary{
				PointCount: 1, FeatureCount: 2, Distance: 3, ElapsedTime: 4,
			})
			if sendErr != nil {
				level.Error(s.log).Log("send_error", sendErr)
			}
			return nil // no error, client just closed
		}
		if err != nil {
			level.Error(s.log).Log("receive_error", err)
			return err // receive error
		}

		req := NewPointFromPB(pbReq)
		// store in repository
		_, _, _, _, err = s.repo.ReverseHalfDuplex(stream.Context(), req.Latitude, req.Longitude)
		if err != nil {
			return err
		}
	}

	return nil
}

// full duplex for FullDuplex
func (s *serviceImpl) FullDuplex(stream pb.RouteGuide_FullDuplexServer) error {
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
				payload := NewPBFromRouteNote(&message)
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

		req := NewRouteNoteFromPB(pbReq)
		resLocation, resMessage, err := s.repo.FullDuplex(stream.Context(), req.Location, req.Message)
		if err != nil {
			return err
		}
		resp := RouteNote{Location: resLocation, Message: resMessage}
		s.broadcastFullDuplex <- resp
	}
	// if we had a previous sending error, we're returning it
	return sendErr
}
