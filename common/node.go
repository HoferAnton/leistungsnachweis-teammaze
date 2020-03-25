package common

type Node interface {
	getConnected() []Node

	getLocation() Location
	isNeighbor(Node) bool

	connect(Node) (bool, Node, Node)
	disconnect(Node) (bool, Node, Node)
	hardCompare(Node) bool
	compare(Node) bool
}
