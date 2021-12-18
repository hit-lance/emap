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

func NewNode(lat, lon float64) *Node {
	return &Node{id: InvalidNodeID, lat: lat, lon: lon}
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

// Distance returns the great-circle distance between vertices v and w in kilometres.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func (n *Node) Distance(n1 *Node) float64 {
	return Distance(n.lat, n.lon, n1.lat, n1.lon)
}

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
