package display

import (
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type Cube struct {
	Center   mgl32.Vec3
	Vertices []mgl32.Vec3

	vbo uint32
}

func NewCube(centerLoc common.Location, xSize float32, ySize float32, zSize float32) Cube {
	xInt, yInt, zInt := centerLoc.As3DCoordinates()
	x := float32(xInt)
	y := float32(yInt)
	z := float32(zInt)

	xSize = xSize / 2
	ySize = ySize / 2
	zSize = zSize / 2

	center := mgl32.Vec3{
		x, y, z,
	}
	cube := Cube{
		Center: center,
		Vertices: []mgl32.Vec3{
			{x + xSize, y + ySize, z - zSize},
			{x - xSize, y + ySize, z - zSize},
			{x + xSize, y + ySize, z + zSize},
			{x - xSize, y + ySize, z + zSize},
			{x + xSize, y - ySize, z - zSize},
			{x - xSize, y - ySize, z - zSize},
			{x - xSize, y - ySize, z + zSize},
			{x + xSize, y - ySize, z + zSize},
		},
	}

	gl.GenBuffers(1, &cube.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, cube.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cube.Vertices)*3*4, gl.Ptr(cube.Vertices), gl.STATIC_READ)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return cube
}

func (cube Cube) Draw() {
	gl.BindBuffer(gl.ARRAY_BUFFER, cube.vbo)
	gl.EnableVertexAttribArray(0)

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 1)

	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
