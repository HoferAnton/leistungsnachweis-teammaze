package common

type Labyrinth interface {
	GetNeighborsByNode(Node) []Node
	GetNeighborsByLocation(Location) []Node
	GetNode(Location) Node
	Compare(Labyrinth) bool
}
