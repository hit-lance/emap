package streetmap

import (
	"math"
	"etaxi/streetmap/graph"
)

type NaiveNodeSet []*graph.Node

func (nns *NaiveNodeSet) Insert(n *graph.Node) {
	*nns = append(*nns, n)
}

func (nns *NaiveNodeSet) Nearest(n *graph.Node) (best *graph.Node) {
	min := math.MaxFloat64
	for _, v := range *nns {
		dis := v.Distance(n)
		if dis < min {
			min = dis
			best = v
		}
	}
	return best
}
