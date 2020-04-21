package display

import (
	"fmt"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func TestCube_String(t *testing.T) {
	x, y, z := randomVec3().Elem()
	sizeX, sizeY, sizeZ := randomVec3().Elem()

	sut := newCube(x, y, z, sizeX, sizeY, sizeZ, nil)

	got := sut.String()
	want := fmt.Sprintf("cube at [%v %v %v] with rendering data: <nil>", x, y, z)

	if got != want {
		t.Errorf("got: %v\nexpected: %v", got, want)
	}
}

func TestNewCubeAppliesTransform(t *testing.T) {
	x, y, z := randomVec3().Elem()
	sizeX, sizeY, sizeZ := randomVec3().Elem()
	baseVec := mgl32.Vec4{2, 2, 2, 1}

	sut := newCube(x, y, z, sizeX, sizeY, sizeZ, nil)
	got := sut.Transform.AsMatrix().Mul4x1(baseVec)
	want := mgl32.Vec4{baseVec.X()*sizeX + x, baseVec.Y()*sizeY + y, baseVec.Z()*sizeZ + z, 1}

	if !got.ApproxEqual(want) {
		t.Errorf("expected: %v\ngot: %v", want, got)
	}
}
