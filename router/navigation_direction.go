package router

import "math"

type DirectionType int

const (
	Start DirectionType = iota
	Straight
	SlightLeft
	SlightRight
	Right
	Left
	SharpLeft
	SharpRight
)

type NavigationDirection struct {
	direction DirectionType
	way       string
	distance  int64
}

func getDirection(prevBearing, curBearing float64) (dt DirectionType) {
	return
}

// Returns the initial bearing (angle) between vertices v and w in degrees.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func bearing(lat1, lon1, lat2, lon2 float64) float64 {
	dlambda := (lon2 - lon1) * math.Pi / 180.0
	phi1 := lat1 * math.Pi / 180.0
	phi2 := lat2 * math.Pi / 180.0

	y := math.Sin(dlambda) * math.Cos(phi2)
	x := math.Cos(phi1)*math.Sin(phi2) -
		math.Sin(phi1)*math.Cos(phi2)*math.Cos(dlambda)
	return math.Atan2(y, x) * 180.0 / math.Pi
}
