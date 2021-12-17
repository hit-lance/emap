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

func (t *KDTree) Nearest(n *graph.Node) (best *graph.Node) {
	if t.root == nil {
		return
	}

	best = t.root.Node
	min := math.MaxFloat64
	nearestHelper(t.root, n, &best, &min, true)
	return
}

func nearestHelper(t *treeNode, n *graph.Node, best **graph.Node, min *float64, flag bool) {
	if t == nil {
		return
	}

	dis := n.Distance(t.Node)
	if dis < *min {
		*min = dis
		*best = t.Node
	}

	var goodSide, badSide *treeNode
	var lat, lon float64

	if (flag && n.Lat() < t.Lat()) || (!flag && n.Lon() < t.Lon()) {
		goodSide, badSide = t.left, t.right
	} else {
		goodSide, badSide = t.right, t.left
	}

	if flag {
		lat, lon = t.Lat(), n.Lon()
	} else {
		lat, lon = n.Lat(), t.Lon()
	}

	nearestHelper(goodSide, n, best, min, !flag)
	if graph.Distance(lat, lon, n.Lat(), n.Lon()) < *min {
		nearestHelper(badSide, n, best, min, !flag)
	}
}
