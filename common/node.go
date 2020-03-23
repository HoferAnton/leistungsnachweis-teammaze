package common

type Labyrinth struct {
	nodes []Node
}

type Node interface {
	GetNeighbors() []Node
	GetConnected() []Node

	Connect(Node) bool
	Disconnect(Node) bool

	GetLocation() Location

	Equals(Node) bool
}
