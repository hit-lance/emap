package streetmap

import (
	"etaxi/streetmap/graph"
	"fmt"
	"math"
)

type KDTree struct {
	root *treeNode
}

type treeNode struct {
	*graph.Node
	left  *treeNode
	right *treeNode
}

func (t *KDTree) String() (s string) {
	stringHelper(t.root, &s)
	return
}

func stringHelper(t *treeNode, s *string) {
	if t == nil {
		return
	}
	*s += fmt.Sprintln(*t.Node)
	stringHelper(t.left, s)
	stringHelper(t.right, s)
}

func (t *KDTree) Insert(n *graph.Node) {
	t.root = insertHelper(t.root, n, true)
}

func insertHelper(t *treeNode, n *graph.Node, flag bool) *treeNode {
	if t == nil {
		return &treeNode{n, nil, nil}
	}
	if (flag && n.Lat() < t.Lat()) || (!flag && n.Lon() < t.Lon()) {
		t.left = insertHelper(t.left, n, !flag)
	} else {
		t.right = insertHelper(t.right, n, !flag)
	}
	return t
}

func (t *KDTree) Nearest(lat, lon float64) (best *graph.Node) {
	if t.root == nil {
		return
	}

	best = t.root.Node
	min := math.MaxFloat64
	nearestHelper(t.root, lat, lon, &best, &min, true)
	return
}

func nearestHelper(t *treeNode, lat, lon float64, best **graph.Node, min *float64, flag bool) {
	if t == nil {
		return
	}

	dis := graph.Distance(lat, lon, t.Lat(), t.Lon())
	if dis < *min {
		*min = dis
		*best = t.Node
	}

	var goodSide, badSide *treeNode
	var x, y float64

	if (flag && lat < t.Lat()) || (!flag && lon < t.Lon()) {
		goodSide, badSide = t.left, t.right
	} else {
		goodSide, badSide = t.right, t.left
	}

	if flag {
		x, y = t.Lat(), lon
	} else {
		x, y = lat, t.Lon()
	}

	nearestHelper(goodSide, lat, lon, best, min, !flag)
	if graph.Distance(x, y, lat, lon) < *min {
		nearestHelper(badSide, lat, lon, best, min, !flag)
	}
}
