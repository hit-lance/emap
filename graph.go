package tinymap

import (
	"fmt"
	"math"
	"os"
)

type Node struct {
	id       int64
	lat, lon float64
	name     string
}

func (n *Node) String() string {
	s := fmt.Sprintf("%+v \n", *n)
	return s
}

type Edge struct {
	u, v   int64
	weight float64
	name   string
}

func (e *Edge) String() string {
	s := fmt.Sprintf("%+v \n", *e)
	return s
}

type Graph struct {
	nodes map[int64]*Node
	adj   map[int64][]*Edge
}

func NewGraph() *Graph {
	var g Graph
	g.nodes = make(map[int64]*Node)
	g.adj = make(map[int64][]*Edge)
	return &g
}

func NewGraphFrom(fn string) *Graph {
	var g Graph
	g.nodes = make(map[int64]*Node)
	g.adj = make(map[int64][]*Edge)
	f, _ := os.Open(fn)
	parseXML(&g, f)
	return &g
}

func (g *Graph) String() string {
	var s string
	s += fmt.Sprintln("nodes: ")
	s += fmt.Sprintln(g.nodes)
	s += fmt.Sprintln("adjacency lists: ")
	s += fmt.Sprintln(g.adj)
	return s
}

func (g *Graph) AddNode(n *Node) {
	if _, ok := g.nodes[n.id]; !ok {
		g.nodes[n.id] = n
	}
}

func (g *Graph) AddEdge(nid1, nid2 int64, name string) {
	node1, ok1 := g.nodes[nid1]
	node2, ok2 := g.nodes[nid2]
	if ok1 && ok2 {
		e := Edge{nid1, nid2, distance(node1.lat, node1.lon, node2.lat, node2.lon), name}
		g.adj[nid1] = append(g.adj[nid1], &e)
		g.adj[nid2] = append(g.adj[nid2], &e)
	}
}

// Returns the great-circle distance between vertices v and w in kilometres.
// Refer from https://www.movable-type.co.uk/scripts/latlong.html
func distance(lat1, lon1, lat2, lon2 float64) float64 {
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
