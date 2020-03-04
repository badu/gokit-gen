package factory

import "github.com/golang/protobuf/ptypes/timestamp"

type Empty struct {
	Page    int32
	PerPage int32
}

type BoolValue struct {
	Value bool
}

type StringValue struct {
	Value     string
	Timestamp timestamp.Timestamp
}

type BoxSpecification struct {
	Name   string
	Height int32
	Width  int32
	Depth  int32
}

type Boxes struct {
	Boxes []BoxSpecification
}

type StatusMessage struct {
	ServiceName string
	Ok          bool
}
