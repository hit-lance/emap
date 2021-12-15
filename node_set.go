package etaxi

type NodeSet interface {
	Insert(n *Node)
	Nearest(n *Node) *Node
}
