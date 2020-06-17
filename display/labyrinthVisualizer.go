package display

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func defaultCubeColor() mgl32.Vec4 {
	return mgl32.Vec4{0, 0.75, 0.75, 1}
}

func pathCubeColor() mgl32.Vec4 {
	return mgl32.Vec4{1, 0, 0, 1}
}

type LabyrinthVisualizer struct {
	mapView         map[mgl32.Vec3]*Cube
	cubes           []Cube
	highlightedPath []*Cube
	steps           []common.Pair
	colorConverter  StepColorConverter
	currentStep     int
}

func NewLabyrinthVisualizer(lab *common.Labyrinth, constructor CubeConstructor) LabyrinthVisualizer {
	if lab == nil {
		panic("passed labyrinth has to be valid")
	}

	view := map[mgl32.Vec3]*Cube{}

	cubes := exploreLabyrinth(lab, constructor)

	vis := LabyrinthVisualizer{
		cubes:           cubes,
		highlightedPath: nil,
		currentStep:     0,
		mapView:         view,
	}

	for i := range vis.cubes {
		view[vis.cubes[i].Transform.GetTranslation()] = &vis.cubes[i]
	}

	vis.mapView = view

	return vis
}

func (vis *LabyrinthVisualizer) GetCubeAt(vec3 mgl32.Vec3) *Cube {
	return vis.mapView[vec3]
}

func (vis *LabyrinthVisualizer) SetSteps(steps []common.Pair, converter StepColorConverter) {
	if vis.steps != nil && vis.colorConverter != nil {
		for _, step := range vis.steps {
			cube, _ := vis.colorConverter.StepToColor(step, vis)
			cube.info.color = defaultCubeColor()
		}
	}

	if steps == nil && converter == nil {
		return
	}

	vis.steps = steps
	vis.colorConverter = converter
	vis.currentStep = 0
}

func (vis *LabyrinthVisualizer) DoStep() {
	if vis.steps == nil {
		panic("cannot do step: steps is nil")
	}

	if vis.colorConverter == nil {
		panic("cannot do step: color converter is nil")
	}

	if vis.currentStep > len(vis.steps) {
		for i := range vis.cubes {
			vis.cubes[i].info.color = defaultCubeColor()
		}

		vis.currentStep = 0
	} else if vis.currentStep == len(vis.steps) {
		vis.currentStep++
		return
	}

	cube, color := vis.colorConverter.StepToColor(vis.steps[vis.currentStep], vis)

	axes := []mgl32.Vec3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	}

	for _, axis := range axes {
		neighbor := vis.GetCubeAt(cube.Transform.GetTranslation().Add(axis))

		if neighbor != nil && neighbor.info.color == color {
			connection := vis.GetCubeAt(cube.Transform.GetTranslation().Add(axis.Mul(0.5))) //nolint:gomnd

			if connection != nil {
				connection.info.color = neighbor.info.color
			}
		}
	}

	cube.info.color = color
	vis.currentStep++
}

func (vis *LabyrinthVisualizer) SetPath(path []common.Location) {
	for i := range vis.cubes {
		vis.cubes[i].info.color = defaultCubeColor()
	}

	if path == nil {
		return
	}

	for locationIndex, location := range path {
		x, y, z := location.As3DCoordinates()

		var nextLocation common.Location

		if locationIndex+1 < len(path) {
			nextLocation = path[locationIndex+1]
		} else {
			nextLocation = location
		}

		xNext, yNext, zNext := nextLocation.As3DCoordinates()
		connectorTranslation := mgl32.Vec3{float32(x+xNext) / 2, float32(y+yNext) / 2, float32(z+zNext) / 2}
		translation := mgl32.Vec3{float32(x), float32(y), float32(z)}

		cube := vis.GetCubeAt(translation)

		if cube != nil {
			cube.info.color = pathCubeColor()
		}

		cube = vis.GetCubeAt(connectorTranslation)

		if cube != nil {
			cube.info.color = pathCubeColor()
		}
	}
}

// This has to be 0.5 due to the way our grid works at the moment.
const cubeSize float32 = 0.5

func (vis *LabyrinthVisualizer) IsValid() bool {
	return len(vis.cubes) > 01
}

func makeConnection(loc common.Location, other common.Location, cubeConstructor CubeConstructor) Cube {
	locX, locY, locZ := loc.As3DCoordinates()
	locV := mgl32.Vec3{float32(locX), float32(locY), float32(locZ)}
	otherX, otherY, otherZ := other.As3DCoordinates()
	otherV := mgl32.Vec3{float32(otherX), float32(otherY), float32(otherZ)}

	// Disable anti-magic number linting here. This Expression calculates the middle point between two Locations by
	// adding half the distance to one of them.
	diffFactor := float32(0.5) //nolint:gomnd
	diffV := otherV.Sub(locV).Mul(diffFactor)

	if diffV.LenSqr() > diffFactor*diffFactor {
		panic("tried to make connection between non-adjacent locations")
	}

	centerV := locV.Add(diffV)

	diffV = diffV.Mul(cubeSize).Add(mgl32.Vec3{cubeSize / 2, cubeSize / 2, cubeSize / 2})

	cube := cubeConstructor(centerV.X(), centerV.Y(), centerV.Z(), diffV.X(), diffV.Y(), diffV.Z())
	cube.info.color = defaultCubeColor()

	return cube
}

func exploreLabyrinth(lab *common.Labyrinth, cubeConstructor CubeConstructor) []Cube {
	cubes := make([]Cube, 0)
	maxX, maxY, maxZ := (*lab).GetMaxLocation().As3DCoordinates()

	for x := uint(0); x <= maxX; x++ {
		for y := uint(0); y <= maxY; y++ {
			for z := uint(0); z <= maxZ; z++ {
				loc := common.NewLocation(x, y, z)
				cube := cubeConstructor(float32(x), float32(y), float32(z), cubeSize, cubeSize, cubeSize)
				cube.info.color = defaultCubeColor()
				cubes = append(cubes, cube)
				cubes = append(cubes, makeConnections(lab, loc, cubeConstructor)...)
			}
		}
	}

	return cubes
}

func makeConnections(lab *common.Labyrinth, loc common.Location, cubeConstructor CubeConstructor) []Cube {
	cubes := make([]Cube, 0)
	connected := (*lab).GetConnected(loc)
	x, y, z := loc.As3DCoordinates()

	for _, other := range connected {
		otherX, otherY, otherZ := other.As3DCoordinates()

		if otherX < x || otherY < y || otherZ < z {
			continue
		}

		cubes = append(cubes, makeConnection(loc, other, cubeConstructor))
	}

	return cubes
}

func (vis LabyrinthVisualizer) String() string {
	return fmt.Sprintf("LabyrinthVisualizer {Cubes: %v, currentStep: %v}", vis.cubes, vis.currentStep)
}
