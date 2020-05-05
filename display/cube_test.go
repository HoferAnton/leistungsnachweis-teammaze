package display

import (
	"fmt"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func TestNewCubePanicsOnNil(t *testing.T) {
	defer func() {
		if got := recover(); got != nil {
			want := "render info may not be nil"
			if got != want {
				t.Errorf("unexpected panic: %v", got)
			}
		} else {
			t.Errorf("expected panic, got none")
		}
	}()

	newCube(0, 0, 0, 0, 0, 0, nil)
}

func TestCube_String(t *testing.T) {
	x, y, z := randomVec3().Elem()
	sizeX, sizeY, sizeZ := randomVec3().Elem()

	info := renderInfo{}

	sut := newCube(x, y, z, sizeX, sizeY, sizeZ, &info)

	got := sut.String()
	want := fmt.Sprintf("cube at [%v %v %v] with rendering data: %v", x, y, z, info)

	if got != want {
		t.Errorf("got: %v\nexpected: %v", got, want)
	}
}

func TestNewCubeAppliesTransform(t *testing.T) {
	x, y, z := randomVec3().Elem()
	sizeX, sizeY, sizeZ := randomVec3().Elem()
	baseVec := mgl32.Vec4{2, 2, 2, 1}

	sut := newCube(x, y, z, sizeX, sizeY, sizeZ, &renderInfo{})
	got := sut.Transform.AsMatrix().Mul4x1(baseVec)
	want := mgl32.Vec4{baseVec.X()*sizeX + x, baseVec.Y()*sizeY + y, baseVec.Z()*sizeZ + z, 1}

	if !got.ApproxEqual(want) {
		t.Errorf("expected: %v\ngot: %v", want, got)
	}
}
