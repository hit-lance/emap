package etaxi

import (
	"fmt"
	"math"
)

const INVALID_NODE_ID = -1

type Node struct {
	id       int64
	lat, lon float64
	name     string
}

func (n *Node) String() string {
	return fmt.Sprintf("%+v\n", *n)
}

// Returns the great-circle distance between vertices v and w in kilometres.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func distance(n1, n2 *Node) float64 {
	const EARTH_RADIUS = 6371

	phi1 := (n2.lat - n1.lat) * (math.Pi / 180.0)
	phi2 := (n2.lon - n1.lon) * (math.Pi / 180.0)
	dphi := n1.lat * (math.Pi / 180.0)
	dlamda := n2.lat * (math.Pi / 180.0)

	a1 := math.Sin(phi1/2) * math.Sin(phi1/2)
	a2 := math.Sin(phi2/2) * math.Sin(phi2/2) * math.Cos(dphi) * math.Cos(dlamda)
	a := a1 + a2
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}

// // Returns the initial bearing (angle) between vertices v and w in degrees.
// // Refer from https://www.movable-type.co.uk/scripts/latlong.html
// func bearing(lat1, lon1, lat2, lon2 float64) float64 {
// 	dlambda := (lon2 - lon1) * math.Pi / 180.0
// 	phi1 := lat1 * math.Pi / 180.0
// 	phi2 := lat2 * math.Pi / 180.0

// 	y := math.Sin(dlambda) * math.Cos(phi2)
// 	x := math.Cos(phi1)*math.Sin(phi2) -
// 		math.Sin(phi1)*math.Cos(phi2)*math.Cos(dlambda)
// 	return math.Atan2(y, x) * 180.0 / math.Pi
// }
