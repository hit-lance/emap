package graph

import (
	"fmt"
	"math"
)

const InvalidNodeID = -1

type Node struct {
	id       int64
	lat, lon float64
	name     string
}

func (n *Node) String() string {
	return fmt.Sprintf("%+v\n", *n)
}

func (n *Node) ID() int64 {
	return n.id
}

func (n *Node) Lat() float64 {
	return n.lat
}

func (n *Node) Lon() float64 {
	return n.lon
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Distance(n1 *Node) float64 {
	return Distance(n.lat, n.lon, n1.lat, n1.lon)
}

// Distance returns the great-circle distance between vertices v and w in kilometres.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371

	phi1 := (lat2 - lat1) * (math.Pi / 180.0)
	phi2 := (lon2 - lon1) * (math.Pi / 180.0)
	dphi := lat1 * (math.Pi / 180.0)
	dlamda := lat2 * (math.Pi / 180.0)

	a1 := math.Sin(phi1/2) * math.Sin(phi1/2)
	a2 := math.Sin(phi2/2) * math.Sin(phi2/2) * math.Cos(dphi) * math.Cos(dlamda)
	a := a1 + a2
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func (n *Node) Bearing(n1 *Node) float64 {
	return Bearing(n.lat, n.lon, n1.lat, n1.lon)
}

// Returns the initial bearing (angle) between vertices v and w in degrees.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func Bearing(lat1, lon1, lat2, lon2 float64) float64 {
	dlambda := (lon2 - lon1) * math.Pi / 180.0
	phi1 := lat1 * math.Pi / 180.0
	phi2 := lat2 * math.Pi / 180.0

	y := math.Sin(dlambda) * math.Cos(phi2)
	x := math.Cos(phi1)*math.Sin(phi2) -
		math.Sin(phi1)*math.Cos(phi2)*math.Cos(dlambda)
	return math.Atan2(y, x) * 180.0 / math.Pi
}
