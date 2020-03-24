package common

type Node interface {
	GetConnected() []Node

	Connect(Node) bool
	Disconnect(Node) bool

	GetLocation() Location

	HardCompare(Node) bool
	Compare(Node) bool
}
