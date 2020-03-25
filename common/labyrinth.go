package common

type Labyrinth interface {
	GetMaxLocation() Location
	GetNeighbors(Location) []Location
	GetConnected(Location) []Location
	Connect(Location, Location) bool
	Disconnect(Location, Location) bool
	Compare(Labyrinth) bool

	checkLocation(Location) bool
	getNode(Location) Node
}
