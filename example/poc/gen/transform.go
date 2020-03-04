package example

import (
	pb "github.com/badu/gokit-gen/example/poc/pb"
)

// transformer *pb.Request to *Request
func NewRequestFromPB(msg *pb.Request) *Request {
	if msg == nil {
		return nil
	}
	var result Request
	result.Msg = msg.Msg
	return &result
}

// transformer *Request to *pb.Request
func NewPBFromRequest(msg *Request) *pb.Request {
	if msg == nil {
		return nil
	}
	var result pb.Request
	result.Msg = msg.Msg
	return &result
}

// transformer *pb.Reply to *Reply
func NewReplyFromPB(msg *pb.Reply) *Reply {
	if msg == nil {
		return nil
	}
	var result Reply
	result.Msg = msg.Msg
	return &result
}

// transformer *Reply to *pb.Reply
func NewPBFromReply(msg *Reply) *pb.Reply {
	if msg == nil {
		return nil
	}
	var result pb.Reply
	result.Msg = msg.Msg
	return &result
}
