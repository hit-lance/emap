package graph

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

func (g *Graph) String() (s string) {
	s += fmt.Sprintln("nodes: ")
	s += fmt.Sprintln(g.nodes)
	s += fmt.Sprintln("neighbors lists: ")
	s += fmt.Sprintln(g.neighbors)
	return s
}

func (g *Graph) NodeIds() []int64 {
	nids := make([]int64, len(g.nodes))
	i := 0

	for k := range g.nodes {
		nids[i] = k
		i++
	}
	return nids
}

func (g *Graph) GetNode(nid int64) *Node {
	return g.nodes[nid]
}

func (g *Graph) GetEdge(from, to int64) *Edge {
	for _, e := range g.Neighbors(from) {
		if e.to == to {
			return e
		}
	}
	return nil
}

func (g *Graph) AddNode(n *Node) {
	if _, ok := g.nodes[n.id]; !ok {
		g.nodes[n.id] = n
	}
}

func (g *Graph) AddEdge(nidFrom, nidTo int64, name string) {
	n1, ok1 := g.nodes[nidFrom]
	n2, ok2 := g.nodes[nidTo]
	if ok1 && ok2 {
		e := Edge{nidFrom, nidTo, n1.Distance(n2), name}
		g.neighbors[nidFrom] = append(g.neighbors[nidFrom], &e)
	}
}

func (g *Graph) Neighbors(nid int64) []*Edge {
	return g.neighbors[nid]
}

func (g *Graph) Contains(nid int64) bool {
	_, ok := g.nodes[nid]
	return ok
}

func (g *Graph) clean() {
	for nid, n := range g.nodes {
		if g.Neighbors(nid) == nil && n.name == "" {
			delete(g.nodes, nid)
		}
	}
}
