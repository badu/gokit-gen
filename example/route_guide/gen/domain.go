package route_guide

type Point struct {
	Latitude  int32
	Longitude int32
}

type Rectangle struct {
	Lo *Point
	Hi *Point
}

type Feature struct {
	Name     string
	Location *Point
}

type RouteNote struct {
	Location *Point
}

type RouteSummary struct {
	PointCount   int32
	FeatureCount int32
	Distance     int32
	ElapsedTime  int32
}
