package streetmap

import (
	"math"
)

type NaiveNodeSet []*Node

func (nns *NaiveNodeSet) Insert(n *Node) {
	*nns = append(*nns, n)
}

func (nns *NaiveNodeSet) Nearest(n *Node) (best *Node) {
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
