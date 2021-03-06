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

const gridStep = 1

func GetIndex(x uint, y uint, z uint, maxLoc Location) uint {
	maxX, maxY, _ := maxLoc.As3DCoordinates()
	return x + y*(maxX+1) + z*(maxX+1)*(maxY+1)
}

func GetLocation(index uint, maxLoc Location) Location {
	maxX, maxY, maxZ := maxLoc.As3DCoordinates()

	if index >= (maxX+gridStep)*(maxY+gridStep)*(maxZ+gridStep) {
		return nil
	}

	upperBoundOfX := maxX + gridStep
	upperBoundOfY := maxY + gridStep
	factorOfZ := upperBoundOfX * upperBoundOfY
	z := index / factorOfZ
	uintIndexWithoutZ := index % factorOfZ
	y := uintIndexWithoutZ / upperBoundOfX
	x := uintIndexWithoutZ % upperBoundOfX

	return NewLocation(x, y, z)
}
