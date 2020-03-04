package factory

import (
	pb "github.com/badu/gokit-gen/example/http/pb"
)

// transformer *pb.Empty to *Empty
func NewEmptyFromPB(msg *pb.Empty) *Empty {
	if msg == nil {
		return nil
	}
	var result Empty
	result.Page = msg.Page
	result.PerPage = msg.PerPage
	return &result
}

// transformer *Empty to *pb.Empty
func NewPBFromEmpty(msg *Empty) *pb.Empty {
	if msg == nil {
		return nil
	}
	var result pb.Empty
	result.Page = msg.Page
	result.PerPage = msg.PerPage
	return &result
}

// transformer *pb.BoolValue to *BoolValue
func NewBoolValueFromPB(msg *pb.BoolValue) *BoolValue {
	if msg == nil {
		return nil
	}
	var result BoolValue
	result.Value = msg.Value
	return &result
}

// transformer *BoolValue to *pb.BoolValue
func NewPBFromBoolValue(msg *BoolValue) *pb.BoolValue {
	if msg == nil {
		return nil
	}
	var result pb.BoolValue
	result.Value = msg.Value
	return &result
}

// transformer *pb.StringValue to *StringValue
func NewStringValueFromPB(msg *pb.StringValue) *StringValue {
	if msg == nil {
		return nil
	}
	var result StringValue
	result.Value = msg.Value
	result.Timestamp = *msg.Timestamp
	return &result
}

// transformer *StringValue to *pb.StringValue
func NewPBFromStringValue(msg *StringValue) *pb.StringValue {
	if msg == nil {
		return nil
	}
	var result pb.StringValue
	result.Value = msg.Value
	result.Timestamp = &msg.Timestamp
	return &result
}

// transformer *pb.BoxSpecification to *BoxSpecification
func NewBoxSpecificationFromPB(msg *pb.BoxSpecification) *BoxSpecification {
	if msg == nil {
		return nil
	}
	var result BoxSpecification
	result.Name = msg.Name
	result.Height = msg.Height
	result.Width = msg.Width
	result.Depth = msg.Depth
	return &result
}

// transformer *BoxSpecification to *pb.BoxSpecification
func NewPBFromBoxSpecification(msg *BoxSpecification) *pb.BoxSpecification {
	if msg == nil {
		return nil
	}
	var result pb.BoxSpecification
	result.Name = msg.Name
	result.Height = msg.Height
	result.Width = msg.Width
	result.Depth = msg.Depth
	return &result
}

// transformer *pb.Boxes to *Boxes
func NewBoxesFromPB(msg *pb.Boxes) *Boxes {
	if msg == nil {
		return nil
	}
	var result Boxes
	for _, Boxes := range msg.Boxes {
		elem := NewBoxSpecificationFromPB(Boxes)
		if elem != nil {
			result.Boxes = append(result.Boxes, *elem)
		}
	}
	return &result
}

// transformer *Boxes to *pb.Boxes
func NewPBFromBoxes(msg *Boxes) *pb.Boxes {
	if msg == nil {
		return nil
	}
	var result pb.Boxes
	for _, Boxes := range msg.Boxes {
		elem := NewPBFromBoxSpecification(&Boxes)
		if elem != nil {
			result.Boxes = append(result.Boxes, elem)
		}
	}
	return &result
}

// transformer *pb.StatusMessage to *StatusMessage
func NewStatusMessageFromPB(msg *pb.StatusMessage) *StatusMessage {
	if msg == nil {
		return nil
	}
	var result StatusMessage
	result.ServiceName = msg.ServiceName
	result.Ok = msg.Ok
	return &result
}

// transformer *StatusMessage to *pb.StatusMessage
func NewPBFromStatusMessage(msg *StatusMessage) *pb.StatusMessage {
	if msg == nil {
		return nil
	}
	var result pb.StatusMessage
	result.ServiceName = msg.ServiceName
	result.Ok = msg.Ok
	return &result
}
