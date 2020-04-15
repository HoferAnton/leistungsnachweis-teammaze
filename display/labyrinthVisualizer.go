package display

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type LabyrinthVisualizer struct {
	cubes []Cube
}

func NewLabyrinthVisualizer(lab *common.Labyrinth) LabyrinthVisualizer {
	if lab == nil {
		panic("passed labyrinth has to be valid")
	}

	shaderProgram, err := CreateProgram("display/shaders/simple_vertex.glsl", "display/shaders/simple_fragment.glsl")

	gl.BindFragDataLocation(shaderProgram, 0, gl.Str("colorOut\x00"))

	FatalIfError("Could not create shader program", err)

	cubes := exploreLabyrinth(lab, shaderProgram, true)

	return LabyrinthVisualizer{
		cubes: cubes,
	}
}

// This has to be 0.5 due to the way our grid works at the moment.
const cubeSize float32 = 0.5

func makeConnection(loc common.Location, other common.Location, cubeShader uint32, genBuffers bool) Cube {
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

	return NewCube(centerV.X(), centerV.Y(), centerV.Z(), diffV.X(), diffV.Y(), diffV.Z(), cubeShader, genBuffers)
}

func exploreLabyrinth(lab *common.Labyrinth, cubeShader uint32, genBuffers bool) []Cube {
	cubes := make([]Cube, 0)
	maxX, maxY, maxZ := (*lab).GetMaxLocation().As3DCoordinates()

	for x := uint(0); x <= maxX; x++ {
		for y := uint(0); y <= maxY; y++ {
			for z := uint(0); z <= maxZ; z++ {
				loc := common.NewLocation(x, y, z)
				cubes = append(cubes,
					NewCube(float32(x), float32(y), float32(z), cubeSize, cubeSize, cubeSize, cubeShader, genBuffers))
				cubes = append(cubes,
					checkAndMakeConnectionsForward(lab, loc, cubeShader, genBuffers)...)
			}
		}
	}

	return cubes
}

func checkAndMakeConnectionsForward(lab *common.Labyrinth,
	loc common.Location,
	cubeShader uint32, genBuffers bool) []Cube {
	x, y, z := loc.As3DCoordinates()
	maxX, maxY, maxZ := (*lab).GetMaxLocation().As3DCoordinates()
	cubes := make([]Cube, 0)

	var other common.Location

	if x < maxX {
		other = common.NewLocation(x+1, y, z)

		if (*lab).IsConnected(loc, other) {
			cubes = append(cubes, makeConnection(loc, other, cubeShader, genBuffers))
		}
	}

	if y < maxY {
		other = common.NewLocation(x, y+1, z)

		if (*lab).IsConnected(loc, other) {
			cubes = append(cubes, makeConnection(loc, other, cubeShader, genBuffers))
		}
	}

	if z < maxZ {
		other = common.NewLocation(x, y, z+1)

		if (*lab).IsConnected(loc, other) {
			cubes = append(cubes, makeConnection(loc, other, cubeShader, genBuffers))
		}
	}

	return cubes
}
