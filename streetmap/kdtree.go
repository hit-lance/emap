package streetmap

import (
	"fmt"
	"math"
)

type KDTree struct {
	root *treeNode
}

type treeNode struct {
	*Node
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

func (t *KDTree) Insert(n *Node) {
	t.root = insertHelper(t.root, n, true)
}

func insertHelper(t *treeNode, n *Node, flag bool) *treeNode {
	if t == nil {
		return &treeNode{n, nil, nil}
	}
	if (flag && n.lat < t.lat) || (!flag && n.lon < t.lon) {
		t.left = insertHelper(t.left, n, !flag)
	} else {
		t.right = insertHelper(t.right, n, !flag)
	}
	return t
}

func (t *KDTree) Nearest(n *Node) (best *Node) {
	if t.root == nil {
		return
	}

	best = t.root.Node
	min := math.MaxFloat64
	nearestHelper(t.root, n, &best, &min, true)
	return
}

func nearestHelper(t *treeNode, n *Node, best **Node, min *float64, flag bool) {
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

	if (flag && n.lat < t.lat) || (!flag && n.lon < t.lon) {
		goodSide, badSide = t.left, t.right
	} else {
		goodSide, badSide = t.right, t.left
	}

	if flag {
		lat, lon = t.lat, n.lon
	} else {
		lat, lon = n.lat, t.lon
	}

	nearestHelper(goodSide, n, best, min, !flag)
	if n.Distance(&Node{lat: lat, lon: lon}) < *min {
		nearestHelper(badSide, n, best, min, !flag)
	}
}
