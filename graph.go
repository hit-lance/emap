package tinymap

import (
	"fmt"
	"os"
)

type Graph struct {
	nodes     map[int64]*Node
	neighbors map[int64][]*Edge
}

func NewGraph() *Graph {
	var g Graph
	g.nodes = make(map[int64]*Node)
	g.neighbors = make(map[int64][]*Edge)
	return &g
}

func NewGraphFrom(fn string) *Graph {
	var g Graph
	g.nodes = make(map[int64]*Node)
	g.neighbors = make(map[int64][]*Edge)
	f, _ := os.Open(fn)
	parseXML(&g, f)
	g.clean()
	return &g
}

func (g *Graph) String() string {
	var s string
	s += fmt.Sprintln("nodes: ")
	s += fmt.Sprintln(g.nodes)
	s += fmt.Sprintln("neighbors lists: ")
	s += fmt.Sprintln(g.neighbors)
	return s
}

func (g *Graph) AddNode(n *Node) {
	if _, ok := g.nodes[n.id]; !ok {
		g.nodes[n.id] = n
	}
}

func (g *Graph) AddEdge(nid1, nid2 int64, name string) {
	n1, ok1 := g.nodes[nid1]
	n2, ok2 := g.nodes[nid2]
	if ok1 && ok2 {
		e := Edge{nid1, nid2, distance(n1, n2), name}
		g.neighbors[nid1] = append(g.neighbors[nid1], &e)
		g.neighbors[nid2] = append(g.neighbors[nid2], &e)
	}
}

func (g *Graph) Neighbors(nid int64) []*Edge {
	return g.neighbors[nid]
}

func (g *Graph) clean() {
	for nid, n := range g.nodes {
		if g.Neighbors(nid) == nil && n.name == "" {
			delete(g.nodes, nid)
		}
	}
}

// // Returns the great-circle distance between vertices v and w in kilometres.
// // Refer from https://www.movable-type.co.uk/scripts/latlong.html
// func distance(lat1, lon1, lat2, lon2 float64) float64 {
// 	const EARTH_RADIUS = 6371

// 	phi1 := (lat2 - lat1) * (math.Pi / 180.0)
// 	phi2 := (lon2 - lon1) * (math.Pi / 180.0)
// 	dphi := lat1 * (math.Pi / 180.0)
// 	dlamda := lat2 * (math.Pi / 180.0)

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
