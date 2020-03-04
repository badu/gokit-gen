package factory_test

import (
	"testing"

	. "github.com/badu/gokit-gen/example/http/gen"
	"github.com/badu/gokit-gen/example/http/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

// test transformer *pb.Empty to *Empty
func TestEmptyFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Empty{

		Page:    0,
		PerPage: 0,
	}
	result := NewEmptyFromPB(&payload)
	assert.Equal(t, payload.Page, result.Page)
	assert.Equal(t, payload.PerPage, result.PerPage)

}

// test transformer *Empty to *pb.Empty
func TestPBFromEmpty(t *testing.T) {
	// TODO : fill me up
	payload := Empty{

		Page:    0,
		PerPage: 0,
	}
	result := NewPBFromEmpty(&payload)
	assert.Equal(t, payload.Page, result.Page)
	assert.Equal(t, payload.PerPage, result.PerPage)

}

// test transformer *pb.BoolValue to *BoolValue
func TestBoolValueFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.BoolValue{

		Value: false,
	}
	result := NewBoolValueFromPB(&payload)
	assert.Equal(t, payload.Value, result.Value)

}

// test transformer *BoolValue to *pb.BoolValue
func TestPBFromBoolValue(t *testing.T) {
	// TODO : fill me up
	payload := BoolValue{

		Value: false,
	}
	result := NewPBFromBoolValue(&payload)
	assert.Equal(t, payload.Value, result.Value)

}

// test transformer *pb.StringValue to *StringValue
func TestStringValueFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.StringValue{

		Value:     "",
		Timestamp: &timestamp.Timestamp{},
	}
	result := NewStringValueFromPB(&payload)
	assert.Equal(t, payload.Value, result.Value)
	assert.Equal(t, payload.Timestamp, result.Timestamp)

}

// test transformer *StringValue to *pb.StringValue
func TestPBFromStringValue(t *testing.T) {
	// TODO : fill me up
	payload := StringValue{

		Value:     "",
		Timestamp: timestamp.Timestamp{},
	}
	result := NewPBFromStringValue(&payload)
	assert.Equal(t, payload.Value, result.Value)
	assert.Equal(t, payload.Timestamp, result.Timestamp)

}

// test transformer *pb.BoxSpecification to *BoxSpecification
func TestBoxSpecificationFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.BoxSpecification{

		Name:   "",
		Height: 0,
		Width:  0,
		Depth:  0,
	}
	result := NewBoxSpecificationFromPB(&payload)
	assert.Equal(t, payload.Name, result.Name)
	assert.Equal(t, payload.Height, result.Height)
	assert.Equal(t, payload.Width, result.Width)
	assert.Equal(t, payload.Depth, result.Depth)

}

// test transformer *BoxSpecification to *pb.BoxSpecification
func TestPBFromBoxSpecification(t *testing.T) {
	// TODO : fill me up
	payload := BoxSpecification{

		Name:   "",
		Height: 0,
		Width:  0,
		Depth:  0,
	}
	result := NewPBFromBoxSpecification(&payload)
	assert.Equal(t, payload.Name, result.Name)
	assert.Equal(t, payload.Height, result.Height)
	assert.Equal(t, payload.Width, result.Width)
	assert.Equal(t, payload.Depth, result.Depth)

}

// test transformer *pb.Boxes to *Boxes
func TestBoxesFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Boxes{
		// repeated Boxes,
	}
	result := NewBoxesFromPB(&payload)
	// repeated Boxes,
	_ = result
}

// test transformer *Boxes to *pb.Boxes
func TestPBFromBoxes(t *testing.T) {
	// TODO : fill me up
	payload := Boxes{
		// repeated Boxes,
	}
	result := NewPBFromBoxes(&payload)
	// repeated Boxes,
	_ = result
}

// test transformer *pb.StatusMessage to *StatusMessage
func TestStatusMessageFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.StatusMessage{

		ServiceName: "",
		Ok:          false,
	}
	result := NewStatusMessageFromPB(&payload)
	assert.Equal(t, payload.ServiceName, result.ServiceName)
	assert.Equal(t, payload.Ok, result.Ok)

}

// test transformer *StatusMessage to *pb.StatusMessage
func TestPBFromStatusMessage(t *testing.T) {
	// TODO : fill me up
	payload := StatusMessage{

		ServiceName: "",
		Ok:          false,
	}
	result := NewPBFromStatusMessage(&payload)
	assert.Equal(t, payload.ServiceName, result.ServiceName)
	assert.Equal(t, payload.Ok, result.Ok)

}
