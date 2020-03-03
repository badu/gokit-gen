package route_guide_test

import (
	"testing"

	. "github.com/badu/gokit-gen/example/route_guide/gen"
	pb "github.com/badu/gokit-gen/example/route_guide/routeguide"
	"github.com/stretchr/testify/assert"
)

// test transformer *pb.Point to *Point
func TestPointFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Point{

		Latitude:  0,
		Longitude: 0,
	}
	result := NewPointFromPB(&payload)
	assert.Equal(t, payload.Latitude, result.Latitude)
	assert.Equal(t, payload.Longitude, result.Longitude)

}

// test transformer *Point to *pb.Point
func TestPBFromPoint(t *testing.T) {
	// TODO : fill me up
	payload := Point{

		Latitude:  0,
		Longitude: 0,
	}
	result := NewPBFromPoint(&payload)
	assert.Equal(t, payload.Latitude, result.Latitude)
	assert.Equal(t, payload.Longitude, result.Longitude)

}

// test transformer *pb.Rectangle to *Rectangle
func TestRectangleFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Rectangle{
		Lo: &pb.Point{

			Latitude:  0,
			Longitude: 0,
		}, Hi: &pb.Point{

			Latitude:  0,
			Longitude: 0,
		},
	}
	result := NewRectangleFromPB(&payload)
	assert.Equal(t, payload.Lo.Latitude, result.Lo.Latitude)
	assert.Equal(t, payload.Lo.Longitude, result.Lo.Longitude)
	assert.Equal(t, payload.Hi.Latitude, result.Hi.Latitude)
	assert.Equal(t, payload.Hi.Longitude, result.Hi.Longitude)

}

// test transformer *Rectangle to *pb.Rectangle
func TestPBFromRectangle(t *testing.T) {
	// TODO : fill me up
	payload := Rectangle{
		Lo: &Point{

			Latitude:  0,
			Longitude: 0,
		}, Hi: &Point{

			Latitude:  0,
			Longitude: 0,
		},
	}
	result := NewPBFromRectangle(&payload)
	assert.Equal(t, payload.Lo.Latitude, result.Lo.Latitude)
	assert.Equal(t, payload.Lo.Longitude, result.Lo.Longitude)
	assert.Equal(t, payload.Hi.Latitude, result.Hi.Latitude)
	assert.Equal(t, payload.Hi.Longitude, result.Hi.Longitude)

}

// test transformer *pb.Feature to *Feature
func TestFeatureFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.Feature{

		Name: "", Location: &pb.Point{

			Latitude:  0,
			Longitude: 0,
		},
	}
	result := NewFeatureFromPB(&payload)
	assert.Equal(t, payload.Name, result.Name)
	assert.Equal(t, payload.Location.Latitude, result.Location.Latitude)
	assert.Equal(t, payload.Location.Longitude, result.Location.Longitude)

}

// test transformer *Feature to *pb.Feature
func TestPBFromFeature(t *testing.T) {
	// TODO : fill me up
	payload := Feature{

		Name: "", Location: &Point{

			Latitude:  0,
			Longitude: 0,
		},
	}
	result := NewPBFromFeature(&payload)
	assert.Equal(t, payload.Name, result.Name)
	assert.Equal(t, payload.Location.Latitude, result.Location.Latitude)
	assert.Equal(t, payload.Location.Longitude, result.Location.Longitude)

}

// test transformer *pb.RouteNote to *RouteNote
func TestRouteNoteFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.RouteNote{
		Location: &pb.Point{

			Latitude:  0,
			Longitude: 0,
		},
		Message: "",
	}
	result := NewRouteNoteFromPB(&payload)
	assert.Equal(t, payload.Location.Latitude, result.Location.Latitude)
	assert.Equal(t, payload.Location.Longitude, result.Location.Longitude)
	assert.Equal(t, payload.Message, result.Message)

}

// test transformer *RouteNote to *pb.RouteNote
func TestPBFromRouteNote(t *testing.T) {
	// TODO : fill me up
	payload := RouteNote{
		Location: &Point{

			Latitude:  0,
			Longitude: 0,
		},
		Message: "",
	}
	result := NewPBFromRouteNote(&payload)
	assert.Equal(t, payload.Location.Latitude, result.Location.Latitude)
	assert.Equal(t, payload.Location.Longitude, result.Location.Longitude)
	assert.Equal(t, payload.Message, result.Message)

}

// test transformer *pb.RouteSummary to *RouteSummary
func TestRouteSummaryFromPB(t *testing.T) {
	// TODO : fill me up
	payload := pb.RouteSummary{

		PointCount:   0,
		FeatureCount: 0,
		Distance:     0,
		ElapsedTime:  0,
	}
	result := NewRouteSummaryFromPB(&payload)
	assert.Equal(t, payload.PointCount, result.PointCount)
	assert.Equal(t, payload.FeatureCount, result.FeatureCount)
	assert.Equal(t, payload.Distance, result.Distance)
	assert.Equal(t, payload.ElapsedTime, result.ElapsedTime)

}

// test transformer *RouteSummary to *pb.RouteSummary
func TestPBFromRouteSummary(t *testing.T) {
	// TODO : fill me up
	payload := RouteSummary{

		PointCount:   0,
		FeatureCount: 0,
		Distance:     0,
		ElapsedTime:  0,
	}
	result := NewPBFromRouteSummary(&payload)
	assert.Equal(t, payload.PointCount, result.PointCount)
	assert.Equal(t, payload.FeatureCount, result.FeatureCount)
	assert.Equal(t, payload.Distance, result.Distance)
	assert.Equal(t, payload.ElapsedTime, result.ElapsedTime)

}
