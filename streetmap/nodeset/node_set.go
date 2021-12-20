package streetmap

import "etaxi/streetmap/graph"

type NodeSet interface {
	Insert(n *graph.Node)
	Nearest(lat, lon float64) *graph.Node
}
