package display

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type Cube struct {
	Center   mgl32.Vec3
	Vertices []mgl32.Vec3

	drawingIndices []uint8

	vao          uint32
	vbo          uint32
	vboIndices   uint32
	matUniformID int32

	projectionMatrix mgl32.Mat4

	shaderProgram uint32
}

func NewCube(centerLoc common.Location, xSize float32, ySize float32, zSize float32, projection mgl32.Mat4) Cube {
	xInt, yInt, zInt := centerLoc.As3DCoordinates()
	x := float32(xInt)
	y := float32(yInt)
	z := float32(zInt)

	xSize /= 2
	ySize /= 2
	zSize /= 2

	center := mgl32.Vec3{
		x, y, z,
	}

	program, err := CreateProgram("display/shaders/simple_vertex.glsl", "display/shaders/simple_fragment.glsl")

	FatalIfError("Could not create shader program", err)

	cube := Cube{
		Center: center,
		Vertices: []mgl32.Vec3{
			{x + xSize, y + ySize, z + zSize},
			{x + xSize, y + ySize, z - zSize},
			{x + xSize, y - ySize, z + zSize},
			{x + xSize, y - ySize, z - zSize},
			{x - xSize, y + ySize, z + zSize},
			{x - xSize, y + ySize, z - zSize},
			{x - xSize, y - ySize, z + zSize},
			{x - xSize, y - ySize, z - zSize},
		},
		drawingIndices: []uint8{
			0, 1, 3,
			0, 3, 2,
			6, 0, 2,
			6, 4, 0,
			7, 4, 6,
			5, 7, 4,
			5, 7, 3,
			5, 3, 1,
		},
		shaderProgram:    program,
		matUniformID:     gl.GetUniformLocation(program, gl.Str("mvp\x00")),
		projectionMatrix: projection,
	}

	generateAndInitializeBuffers(&cube)
	gl.BindFragDataLocation(program, 0, gl.Str("colorOut\x00"))

	positionAttrib := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(positionAttrib)

	stride := int32(unsafe.Sizeof(mgl32.Vec3{}))
	gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))

	gl.BindVertexArray(0)
	return cube
}

func generateAndInitializeBuffers(cube *Cube) {
	gl.GenVertexArrays(1, &cube.vao)
	gl.BindVertexArray(cube.vao)

	gl.GenBuffers(1, &cube.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, cube.vbo)

	gl.GenBuffers(1, &cube.vboIndices)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, cube.vboIndices)

	numVertexBufferBytes := len(cube.Vertices) * int(unsafe.Sizeof(mgl32.Vec3{}))
	gl.BufferData(gl.ARRAY_BUFFER, numVertexBufferBytes, gl.Ptr(cube.Vertices), gl.STATIC_READ)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(cube.drawingIndices), gl.Ptr(cube.drawingIndices), gl.STATIC_READ)
}

func (cube Cube) draw(view mgl32.Mat4, timeSinceStart time.Duration) {
	gl.BindVertexArray(cube.vao)
	checkForGLError(fmt.Sprintf("glGetError not zero after BindVertexArray(%v)", cube.vao))

	gl.UseProgram(cube.shaderProgram)
	model := mgl32.HomogRotate3DY(float32(timeSinceStart.Seconds()))

	mvp := cube.projectionMatrix.Mul4(view.Mul4(model))
	gl.UniformMatrix4fv(cube.matUniformID, 1, false, &mvp[0])

	gl.DrawElements(gl.TRIANGLES, int32(len(cube.drawingIndices)), gl.UNSIGNED_BYTE, gl.PtrOffset(0))

	gl.BindVertexArray(0)
}
