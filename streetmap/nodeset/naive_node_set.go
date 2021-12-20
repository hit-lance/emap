package streetmap

import (
	"etaxi/streetmap/graph"
	"math"
)

type NaiveNodeSet []*graph.Node

func (nns *NaiveNodeSet) Insert(n *graph.Node) {
	*nns = append(*nns, n)
}

func (nns *NaiveNodeSet) Nearest(lat, lon float64) (best *graph.Node) {
	min := math.MaxFloat64
	for _, v := range *nns {
		dis := graph.Distance(lat, lon, v.Lat(), v.Lon())
		if dis < min {
			min = dis
			best = v
		}
	}
	return best
}
