package graph

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

func NewNode(lat, lon float64) *Node {
	return &Node{id: INVALID_NODE_ID, lat: lat, lon: lon}
}

func (n *Node) String() string {
	return fmt.Sprintf("%+v\n", *n)
}

func (n *Node) Id() int64 {
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

// Returns the great-circle distance between vertices v and w in kilometres.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func (n1 *Node) Distance(n2 *Node) float64 {
	return Distance(n1.lat, n1.lon, n2.lat, n2.lon)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	const EARTH_RADIUS = 6371

	phi1 := (lat2 - lat1) * (math.Pi / 180.0)
	phi2 := (lon2 - lon1) * (math.Pi / 180.0)
	dphi := lat1 * (math.Pi / 180.0)
	dlamda := lat2 * (math.Pi / 180.0)

	a1 := math.Sin(phi1/2) * math.Sin(phi1/2)
	a2 := math.Sin(phi2/2) * math.Sin(phi2/2) * math.Cos(dphi) * math.Cos(dlamda)
	a := a1 + a2
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}

// func Distance(n1, n2 *Node) float64 {
// 	const EARTH_RADIUS = 6371

// 	phi1 := (n2.lat - n1.lat) * (math.Pi / 180.0)
// 	phi2 := (n2.lon - n1.lon) * (math.Pi / 180.0)
// 	dphi := n1.lat * (math.Pi / 180.0)
// 	dlamda := n2.lat * (math.Pi / 180.0)

// 	a1 := math.Sin(phi1/2) * math.Sin(phi1/2)
// 	a2 := math.Sin(phi2/2) * math.Sin(phi2/2) * math.Cos(dphi) * math.Cos(dlamda)
// 	a := a1 + a2
// 	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

// 	return EARTH_RADIUS * c
// }

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
