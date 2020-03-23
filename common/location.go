package common

type Location3D struct {
	x, y, z int
}

type Location interface {
	As3DCoordinates() (int, int, int)
}
