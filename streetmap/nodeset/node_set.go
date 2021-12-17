package streetmap

import "etaxi/streetmap/graph"

type NodeSet interface {
	Insert(n *graph.Node)
	Nearest(n *graph.Node) *graph.Node
}
