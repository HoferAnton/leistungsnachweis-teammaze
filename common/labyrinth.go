package common

type Labyrinth interface {
	GetMaxLocation() Location
	GetNeighbors(Location) []Location
	GetConnected(Location) []Location
	IsConnected(Location, Location) bool

	Connect(Location, Location) bool
	Disconnect(Location, Location) bool

	Compare(Labyrinth) bool

	CheckLocation(Location) bool
	getNode(Location) Node
}
