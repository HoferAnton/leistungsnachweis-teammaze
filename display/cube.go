package display

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Cube struct {
	Center   mgl32.Vec3
	Vertices []mgl32.Vec3

	drawingIndices []uint8

	vao                    uint32
	vbo                    uint32
	vboIndices             uint32
	mvpUniformID           int32
	viewMatUniformID       int32
	modelMatUniformID      int32
	lightPositionUniformID int32

	shaderProgram uint32
}

func NewCube(x, y, z, xSize, ySize, zSize float32, shaderProgram uint32) Cube {
	if xSize < 0 || ySize < 0 || zSize < 0 {
		panic("Negative Size given")
	}

	xSize /= 2
	ySize /= 2
	zSize /= 2

	center := mgl32.Vec3{
		x, y, z,
	}

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
			5, 3, 1,
			5, 7, 3,
			5, 7, 4,
			7, 4, 6,
			6, 4, 0,
			6, 0, 2,
			0, 3, 2,
			0, 1, 3,
			4, 5, 0,
			5, 1, 0,
			6, 7, 2,
			7, 3, 2,
		},
		shaderProgram:          shaderProgram,
		mvpUniformID:           gl.GetUniformLocation(shaderProgram, gl.Str("MVP\x00")),
		viewMatUniformID:       gl.GetUniformLocation(shaderProgram, gl.Str("V\x00")),
		modelMatUniformID:      gl.GetUniformLocation(shaderProgram, gl.Str("M\x00")),
		lightPositionUniformID: gl.GetUniformLocation(shaderProgram, gl.Str("lightPosition_worldSpace\x00")),
	}

	generateAndInitializeBuffers(&cube)

	positionAttrib := uint32(gl.GetAttribLocation(shaderProgram, gl.Str("position_modelSpace\x00")))
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

func (cube Cube) draw(view mgl32.Mat4, projection mgl32.Mat4, labCenter mgl32.Vec3, lightPosition mgl32.Vec3, timeSinceStart time.Duration) {
	gl.BindVertexArray(cube.vao)
	checkForGLError(fmt.Sprintf("glGetError not zero after BindVertexArray(%v)", cube.vao))

	gl.UseProgram(cube.shaderProgram)

	model := mgl32.HomogRotate3DY(float32(timeSinceStart.Seconds())).
		Mul4(mgl32.Translate3D(-labCenter.X(), -labCenter.Y(), -labCenter.Z()))
	mvp := projection.Mul4(view.Mul4(model))

	gl.UniformMatrix4fv(cube.mvpUniformID, 1, false, &mvp[0])
	gl.UniformMatrix4fv(cube.modelMatUniformID, 1, false, &model[0])
	gl.UniformMatrix4fv(cube.viewMatUniformID, 1, false, &view[0])

	gl.Uniform3fv(cube.lightPositionUniformID, 1, &lightPosition[0])

	gl.DrawElements(gl.TRIANGLES, int32(len(cube.drawingIndices)), gl.UNSIGNED_BYTE, gl.PtrOffset(0))

	gl.BindVertexArray(0)
}

func (cube Cube) String() string {
	var vaoValidness string

	if cube.vao == 0 {
		vaoValidness = "invalid"
	} else {
		vaoValidness = "valid"
	}

	return fmt.Sprintf("Cube at (%v) with %s vao.", cube.Center, vaoValidness)
}
