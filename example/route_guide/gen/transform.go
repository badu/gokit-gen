package route_guide

import (
	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
)

// transformer *pb.Point to *Point
func NewPointFromPB(msg *pb.Point) *Point {
	if msg == nil {
		return nil
	}
	var result Point
	result.Latitude = msg.Latitude
	result.Longitude = msg.Longitude
	return &result
}

// transformer *Point to *pb.Point
func NewPBFromPoint(msg *Point) *pb.Point {
	if msg == nil {
		return nil
	}
	var result pb.Point
	result.Latitude = msg.Latitude
	result.Longitude = msg.Longitude
	return &result
}

// transformer *pb.Rectangle to *Rectangle
func NewRectangleFromPB(msg *pb.Rectangle) *Rectangle {
	if msg == nil {
		return nil
	}
	var result Rectangle
	result.Lo = NewPointFromPB(msg.Lo)
	result.Hi = NewPointFromPB(msg.Hi)
	return &result
}

// transformer *Rectangle to *pb.Rectangle
func NewPBFromRectangle(msg *Rectangle) *pb.Rectangle {
	if msg == nil {
		return nil
	}
	var result pb.Rectangle
	result.Lo = NewPBFromPoint(msg.Lo)
	result.Hi = NewPBFromPoint(msg.Hi)
	return &result
}

// transformer *pb.Feature to *Feature
func NewFeatureFromPB(msg *pb.Feature) *Feature {
	if msg == nil {
		return nil
	}
	var result Feature
	result.Name = msg.Name
	result.Location = NewPointFromPB(msg.Location)
	return &result
}

// transformer *Feature to *pb.Feature
func NewPBFromFeature(msg *Feature) *pb.Feature {
	if msg == nil {
		return nil
	}
	var result pb.Feature
	result.Name = msg.Name
	result.Location = NewPBFromPoint(msg.Location)
	return &result
}

// transformer *pb.RouteNote to *RouteNote
func NewRouteNoteFromPB(msg *pb.RouteNote) *RouteNote {
	if msg == nil {
		return nil
	}
	var result RouteNote
	result.Location = NewPointFromPB(msg.Location)
	return &result
}

// transformer *RouteNote to *pb.RouteNote
func NewPBFromRouteNote(msg *RouteNote) *pb.RouteNote {
	if msg == nil {
		return nil
	}
	var result pb.RouteNote
	result.Location = NewPBFromPoint(msg.Location)
	return &result
}

// transformer *pb.RouteSummary to *RouteSummary
func NewRouteSummaryFromPB(msg *pb.RouteSummary) *RouteSummary {
	if msg == nil {
		return nil
	}
	var result RouteSummary
	result.PointCount = msg.PointCount
	result.FeatureCount = msg.FeatureCount
	result.Distance = msg.Distance
	result.ElapsedTime = msg.ElapsedTime
	return &result
}

// transformer *RouteSummary to *pb.RouteSummary
func NewPBFromRouteSummary(msg *RouteSummary) *pb.RouteSummary {
	if msg == nil {
		return nil
	}
	var result pb.RouteSummary
	result.PointCount = msg.PointCount
	result.FeatureCount = msg.FeatureCount
	result.Distance = msg.Distance
	result.ElapsedTime = msg.ElapsedTime
	return &result
}
