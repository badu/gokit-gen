package example_test

import (
	"testing"

	. "github.com/badu/gokit-gen/example/poc/gen"
	pb "github.com/badu/gokit-gen/example/poc/pb"
	"github.com/stretchr/testify/assert"
)

// test transformer *pb.Request to *Request
func TestRequestFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Request{

		Msg: "",
	}
	result := NewRequestFromPB(&payload)
	assert.Equal(t, payload.Msg, result.Msg)

}

// test transformer *Request to *pb.Request
func TestPBFromRequest(t *testing.T) {
	// TODO : fill me up
	payload := Request{

		Msg: "",
	}
	result := NewPBFromRequest(&payload)
	assert.Equal(t, payload.Msg, result.Msg)

}

// test transformer *pb.Reply to *Reply
func TestReplyFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Reply{

		Msg: "",
	}
	result := NewReplyFromPB(&payload)
	assert.Equal(t, payload.Msg, result.Msg)

}

// test transformer *Reply to *pb.Reply
func TestPBFromReply(t *testing.T) {
	// TODO : fill me up
	payload := Reply{

		Msg: "",
	}
	result := NewPBFromReply(&payload)
	assert.Equal(t, payload.Msg, result.Msg)

}
