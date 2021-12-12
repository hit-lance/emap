package tinymap

type NodeSet interface {
	Add(n *Node)
	Nearest(n *Node) *Node
}
