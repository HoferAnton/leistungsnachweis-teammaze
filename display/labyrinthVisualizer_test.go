package display

import (
	"testing"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func testingCubeConstructor(x, y, z, xSize, ySize, zSize float32) Cube {
	return newCube(x, y, z, xSize, ySize, zSize, nil)
}

func TestMakeConnectionFailsOnNonAdjacent(t *testing.T) {
	loc1 := common.NewLocation(0, 0, 0)
	loc2 := common.NewLocation(1, 1, 1)

	defer func() {
		want := "tried to make connection between non-adjacent locations"
		if got := recover(); got != nil {
			if got != want {
				t.Errorf("expected \"%s\" but got \"%s\"", want, got)
			}
		} else {
			t.Error("Expected panic, got none")
		}
	}()

	makeConnection(loc1, loc2, testingCubeConstructor)
}

func TestMakeConnectionHasCorrectCenter(t *testing.T) {
	loc1 := common.NewLocation(1, 2, 1)
	loc2 := common.NewLocation(1, 1, 1)
	cube := makeConnection(loc1, loc2, testingCubeConstructor)

	got := cube.Transform.translation.Mul4x1(mgl32.Vec4{0, 0, 0, 1}).Vec3()
	want := mgl32.Vec3{1, 1.5, 1}

	if got != want {
		t.Errorf("got: %v\nexpected: %v", got, want)
	}
}

func TestMakeConnectionHasCorrectScale(t *testing.T) {
	loc1 := common.NewLocation(3, 3, 3)
	loc2 := common.NewLocation(4, 3, 3)
	cube := makeConnection(loc1, loc2, testingCubeConstructor)

	got := cube.Transform.scale.Mul4x1(mgl32.Vec4{1, 1, 1, 1}).Vec3()
	want := mgl32.Vec3{0.5, 0.25, 0.25}

	if got != want {
		t.Errorf("got: %v\nexpected: %v", got, want)
	}
}

func TestCheckAndMake(t *testing.T) {
	maxLoc := common.NewLocation(2, 2, 2)
	lab := common.NewLabyrinth(maxLoc)

	baseLoc := common.NewLocation(1, 1, 1)
	lab.Connect(baseLoc, common.NewLocation(1, 2, 1))
	lab.Connect(baseLoc, common.NewLocation(1, 1, 2))

	wantedCubes := []Cube{
		newCube(1, 1.5, 1, 0.25, 0.5, 0.25, nil),
		newCube(1, 1, 1.5, 0.25, 0.25, 0.5, nil),
	}
	cubes := checkAndMakeConnections(&lab, baseLoc, testingCubeConstructor)

	compareCubeSlices(t, cubes, wantedCubes)
}

func TestExploreLabyrinth(t *testing.T) {
	maxLoc := common.NewLocation(1, 1, 2)
	lab := common.NewLabyrinth(maxLoc)

	wantedCubes := []Cube{
		// Nodes
		newCube(0, 0, 0, 0.5, 0.5, 0.5, nil),
		newCube(0, 0, 1, 0.5, 0.5, 0.5, nil),
		newCube(0, 0, 2, 0.5, 0.5, 0.5, nil),
		newCube(0, 1, 0, 0.5, 0.5, 0.5, nil),
		newCube(0, 1, 1, 0.5, 0.5, 0.5, nil),
		newCube(0, 1, 2, 0.5, 0.5, 0.5, nil),
		newCube(1, 0, 0, 0.5, 0.5, 0.5, nil),
		newCube(1, 0, 1, 0.5, 0.5, 0.5, nil),
		newCube(1, 0, 2, 0.5, 0.5, 0.5, nil),
		newCube(1, 1, 0, 0.5, 0.5, 0.5, nil),
		newCube(1, 1, 1, 0.5, 0.5, 0.5, nil),
		newCube(1, 1, 2, 0.5, 0.5, 0.5, nil),
		// Connections
		newCube(0, 0.5, 0, 0.25, 0.5, 0.25, nil),
		newCube(0, 0, 0.5, 0.25, 0.25, 0.5, nil),
		newCube(1, 0, 1.5, 0.25, 0.25, 0.5, nil),
		newCube(1, 0.5, 1, 0.25, 0.5, 0.25, nil),
		newCube(0, 0.5, 1, 0.25, 0.5, 0.25, nil),
		newCube(1, 1, 1.5, 0.25, 0.25, 0.5, nil),
	}

	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(0, 1, 0))
	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(0, 0, 1))
	lab.Connect(common.NewLocation(1, 0, 1), common.NewLocation(1, 0, 2))
	lab.Connect(common.NewLocation(1, 0, 1), common.NewLocation(1, 1, 1))
	lab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(0, 1, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 1, 2))

	cubes := exploreLabyrinth(&lab, testingCubeConstructor)

	compareCubeSlices(t, cubes, wantedCubes)
}

func TestNewLabyrinthVisualizerPanicsOnNil(t *testing.T) {
	defer func() {
		want := "passed labyrinth has to be valid"
		if got := recover(); got != nil {
			if got != want {
				t.Errorf("Unexpected panic: \"%s\", expected: \"%s\"", got, want)
			}
		} else {
			t.Errorf("Expected panic, got none")
		}
	}()

	NewLabyrinthVisualizer(nil, testingCubeConstructor)
}

// Helpers:

func compareCubeSlices(t *testing.T, cubes []Cube, wantedCubes []Cube) {
	if len(cubes) < len(wantedCubes) {
		t.Errorf("Not enough cubes")
	} else if len(cubes) > len(wantedCubes) {
		t.Errorf("Too many cubes")
	}

	for _, cube := range cubes {
		isInWanted := false

		for i, wantedCube := range wantedCubes {
			if cube == wantedCube {
				wantedCubes = append(wantedCubes[0:i], wantedCubes[i+1:]...)
				isInWanted = true

				break
			}
		}

		if !isInWanted {
			t.Errorf("Unexpected Cube at %v", cube.Transform.GetTranslation())
		}
	}

	if len(wantedCubes) != 0 {
		for _, wantedCube := range wantedCubes {
			t.Errorf("Failed to make cube at %v\n", wantedCube.Transform.GetTranslation())
		}
	}
}
