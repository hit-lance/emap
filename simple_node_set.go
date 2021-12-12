package tinymap

import (
	"math"
)

type SimpleNodeSet []*Node

func (sns *SimpleNodeSet) Add(n *Node) {
	*sns = append(*sns, n)
}

func (sns *SimpleNodeSet) Nearest(n *Node) (best *Node) {
	min := math.MaxFloat64
	for _, v := range *sns {
		dis := distance(v, n)
		if dis < min {
			min = dis
			best = v
		}
	}
	return best
}
