package display

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Constant Cube Vertices. Same for all cubes.
func getVertices() []mgl32.Vec3 {
	return []mgl32.Vec3{
		{0.5, 0.5, 0.5},
		{0.5, 0.5, -0.5},
		{0.5, -0.5, 0.5},
		{0.5, -0.5, -0.5},
		{-0.5, 0.5, 0.5},
		{-0.5, 0.5, -0.5},
		{-0.5, -0.5, 0.5},
		{-0.5, -0.5, -0.5},
	}
}

// Constant Drawing Indices, defines a number of triangles that are used to render cubes
func getDrawingIndices() []uint8 {
	return []uint8{
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
	}
}

type Cube struct {
	Transform Transform

	info renderInfo
}

type CubeConstructor = func(x, y, z, xSize, ySize, zSize float32) Cube

type renderInfo struct {
	vao                    uint32
	vbo                    uint32
	vboIndices             uint32
	mvpUniformID           int32
	viewMatUniformID       int32
	modelMatUniformID      int32
	lightPositionUniformID int32
	materialColorUniformID int32

	color mgl32.Vec4

	shaderProgram uint32
}

func GetCubeConstructor(shaderProgram uint32) CubeConstructor {
	vao, vbo, indicesVbo := generateAndInitializeBuffers()

	info := renderInfo{
		vao:                    vao,
		vbo:                    vbo,
		vboIndices:             indicesVbo,
		mvpUniformID:           gl.GetUniformLocation(shaderProgram, gl.Str("MVP\x00")),
		viewMatUniformID:       gl.GetUniformLocation(shaderProgram, gl.Str("V\x00")),
		modelMatUniformID:      gl.GetUniformLocation(shaderProgram, gl.Str("M\x00")),
		lightPositionUniformID: gl.GetUniformLocation(shaderProgram, gl.Str("lightPosition_worldSpace\x00")),
		materialColorUniformID: gl.GetUniformLocation(shaderProgram, gl.Str("materialDiffuseColor\x00")),
		shaderProgram:          shaderProgram,
		color:                  mgl32.Vec4{},
	}

	positionAttrib := uint32(gl.GetAttribLocation(shaderProgram, gl.Str("position_modelSpace\x00")))
	gl.EnableVertexAttribArray(positionAttrib)

	stride := int32(unsafe.Sizeof(mgl32.Vec3{}))
	gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))

	gl.BindVertexArray(0)

	return func(x, y, z, xSize, ySize, zSize float32) Cube {
		return newCube(x, y, z, xSize, ySize, zSize, &info)
	}
}

// Create a new cube.
func newCube(x, y, z, xSize, ySize, zSize float32, info *renderInfo) Cube {
	if xSize < 0 || ySize < 0 || zSize < 0 {
		panic("Negative Size given")
	}

	if info == nil {
		panic("render info may not be nil")
	}

	cube := Cube{
		Transform: Transform{
			translation: mgl32.Translate3D(x, y, z),
			scale:       mgl32.Scale3D(xSize, ySize, zSize),
			rotation:    mgl32.QuatIdent(),
		},
		info: *info,
	}

	return cube
}

// Creates a VAO (Vertex Array Object)
// Creates two VBO (Vertex Buffer Objects) and binds them into ARRAY and ELEMENT_ARRAY Buffers of the created vao
// loads vertex data into ARRAY_BUFFER and index data into ELEMENT_ARRAY_BUFFER
// this could be optimized because every cube has the same vertex and index data, we don't have to keep it in
// 		 memory for every cube.
func generateAndInitializeBuffers() (uint32, uint32, uint32) {
	var vao, vbo, indicesVbo uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.GenBuffers(1, &indicesVbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indicesVbo)

	numVertexBufferBytes := len(getVertices()) * int(unsafe.Sizeof(mgl32.Vec3{}))
	gl.BufferData(gl.ARRAY_BUFFER, numVertexBufferBytes, gl.Ptr(getVertices()), gl.STATIC_READ)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(getDrawingIndices()), gl.Ptr(getDrawingIndices()), gl.STATIC_READ)

	return vao, vbo, indicesVbo
}

// Draws the cube into the default framebuffer with the specified view, projection matrices and an arbitrary transform
// lightPosition is a shader parameter
func (cube Cube) draw(view, projection, transform *mgl32.Mat4, lightPosition mgl32.Vec3) {
	gl.BindVertexArray(cube.info.vao)
	checkForGLError(fmt.Sprintf("glGetError not zero after BindVertexArray(%v)", cube.info.vao))

	gl.UseProgram(cube.info.shaderProgram)

	model := transform.Mul4(cube.Transform.AsMatrix())
	mvp := projection.Mul4(view.Mul4(model))

	gl.UniformMatrix4fv(cube.info.mvpUniformID, 1, false, &mvp[0])
	gl.UniformMatrix4fv(cube.info.modelMatUniformID, 1, false, &model[0])
	gl.UniformMatrix4fv(cube.info.viewMatUniformID, 1, false, &view[0])

	gl.Uniform4fv(cube.info.materialColorUniformID, 1, &cube.info.color[0])
	gl.Uniform3fv(cube.info.lightPositionUniformID, 1, &lightPosition[0])

	gl.DrawElements(gl.TRIANGLES, int32(len(getDrawingIndices())), gl.UNSIGNED_BYTE, gl.PtrOffset(0))

	gl.BindVertexArray(0)
}

// Pretty prints location and vao validness
func (cube Cube) String() string {
	position := cube.Transform.translation.Mul4x1(mgl32.Vec4{0, 0, 0, 1}).Vec3()

	return fmt.Sprintf("cube at %v", position)
}
