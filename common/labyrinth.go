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

func GetIndex(x uint, y uint, z uint, maxLoc Location) uint {
	maxX, maxY, _ := maxLoc.As3DCoordinates()
	return x + y*(maxX+1) + z*(maxX+1)*(maxY+1)
}
