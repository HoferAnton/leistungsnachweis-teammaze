package display

import (
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
	cubes           []Cube
	highlightedPath []*Cube
	steps           []common.Pair
	colorConverter  StepColorConverter
	currentStep     int
}

func NewLabyrinthVisualizer(lab *common.Labyrinth, steps []common.Pair, stepToColor StepColorConverter, constructor CubeConstructor) LabyrinthVisualizer {
	if lab == nil {
		panic("passed labyrinth has to be valid")
	}

	cubes := exploreLabyrinth(lab, constructor)

	return LabyrinthVisualizer{
		cubes:           cubes,
		highlightedPath: nil,
		steps:           steps,
		currentStep:     0,
		colorConverter:  stepToColor,
	}
}

func (vis *LabyrinthVisualizer) DoStep() {
	if vis.currentStep == len(vis.steps) {
		for _, cube := range vis.cubes {
			cube.info.color = defaultCubeColor()
		}

		vis.currentStep = 0
	}

	cube, color := vis.colorConverter.StepToColor(vis.steps[vis.currentStep], vis.cubes)

	cube.info.color = color
}

func (vis *LabyrinthVisualizer) SetPath(path []common.Location) {
	if vis.highlightedPath != nil {
		for _, cube := range vis.highlightedPath {
			cube.info.color = defaultCubeColor()
		}
	}

	vis.highlightedPath = make([]*Cube, 2*len(path)-1)

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

		for i, cube := range vis.cubes {
			if cube.Transform.GetTranslation() == translation || cube.Transform.GetTranslation() == connectorTranslation {
				vis.cubes[i].info.color = pathCubeColor()
				vis.highlightedPath = append(vis.highlightedPath, &vis.cubes[i])
			}
		}
	}
}

// This has to be 0.5 due to the way our grid works at the moment.
const cubeSize float32 = 0.5

func (vis *LabyrinthVisualizer) IsValid() bool {
	return len(vis.cubes) > 0
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
