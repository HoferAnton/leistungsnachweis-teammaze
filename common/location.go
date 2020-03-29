package common

type Location3D struct {
	x, y, z uint
}

type Location interface {
	As3DCoordinates() (uint, uint, uint)
	Compare(Location) bool
}

type LocationImpl struct {
	location3D Location3D
}

func NewLocation(xCoordinate uint, yCoordinate uint, zCoordinate uint) Location {
	return LocationImpl{Location3D{xCoordinate, yCoordinate, zCoordinate}}
}

func (l LocationImpl) As3DCoordinates() (uint, uint, uint) {
	return l.location3D.x, l.location3D.y, l.location3D.z
}

func (l LocationImpl) Compare(that Location) bool {
	if that == nil {
		return false
	}

	thisX, thisY, thisZ := l.As3DCoordinates()
	thatX, thatY, thatZ := that.As3DCoordinates()

	return thisX == thatX && thisY == thatY && thisZ == thatZ
}
