package display

import (
	"github.com/go-gl/gl/v4.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
)

type LabyrinthVisualizer struct {
	cubes []Cube
}

func NewLabyrinthVisualizer(lab *common.Labyrinth) LabyrinthVisualizer {
	//maxX, maxY, maxZ := (*lab).GetMaxLocation().As3DCoordinates()
	//numNodes := (maxX + 1) * (maxY + 1) * (maxZ + 1)
	//visited := make([]common.Location, 0)
	//accumulator := make([]Cube, 0)

	shaderProgram, err := CreateProgram("display/shaders/simple_vertex.glsl", "display/shaders/simple_fragment.glsl")

	gl.BindFragDataLocation(shaderProgram, 0, gl.Str("colorOut\x00"))

	FatalIfError("Could not create shader program", err)

	//exploreLabyrinth((*lab).GetMaxLocation(), lab, visited, accumulator, shaderProgram)

	cubes := exploreLabyrinth(lab, shaderProgram)

	return LabyrinthVisualizer{
		cubes: cubes,
	}
}

func exploreRecursively(loc common.Location, lab *common.Labyrinth, visited []common.Location, accumulator []Cube, cubeShader uint32) {
	visited = append(visited, loc)

	locX, locY, locZ := loc.As3DCoordinates()

	locV := mgl32.Vec3{
		float32(locX),
		float32(locY),
		float32(locZ),
	}

	cube := NewCube(locV.X(), locV.Y(), locV.Z(), 0.75, 0.75, 0.75, cubeShader)
	accumulator = append(accumulator, cube)

outer:
	for _, node := range (*lab).GetNeighbors(loc) {
		for _, other := range visited {
			if other.Compare(node) {
				continue outer
			}
		}

		nodeX, nodeY, nodeZ := node.As3DCoordinates()
		nodeV := mgl32.Vec3{
			float32(nodeX),
			float32(nodeY),
			float32(nodeZ),
		}

		diffV := nodeV.Sub(locV)
		connV := locV.Add(diffV)

		diffV = diffV.Mul(0.5).Add(mgl32.Vec3{1, 1, 1}.Mul(0.125))

		newCube := NewCube(connV.X(), connV.Y(), connV.Z(), diffV.X(), diffV.Y(), diffV.Z(), cubeShader)
		accumulator = append(accumulator, newCube)

		exploreRecursively(node, lab, visited, accumulator, cubeShader)
	}
}

func exploreLabyrinth(lab *common.Labyrinth, cubeShader uint32) []Cube {
	cubes := make([]Cube, 0)
	maxX, maxY, maxZ := (*lab).GetMaxLocation().As3DCoordinates()

	isConnected := func(loc common.Location, connected []common.Location) bool {
		for _, conn := range connected {
			if loc.Compare(conn) {
				return true
			}
		}

		return false
	}

	makeConnection := func(loc common.Location, other common.Location, cubeShader uint32) Cube {
		locX, locY, locZ := loc.As3DCoordinates()
		locV := mgl32.Vec3{float32(locX), float32(locY), float32(locZ)}
		otherX, otherY, otherZ := other.As3DCoordinates()
		otherV := mgl32.Vec3{float32(otherX), float32(otherY), float32(otherZ)}

		diffV := otherV.Sub(locV).Mul(0.5)

		centerV := locV.Add(diffV)

		diffV = diffV.Mul(0.5).Add(mgl32.Vec3{0.25, 0.25, 0.25})

		return NewCube(centerV.X(), centerV.Y(), centerV.Z(), diffV.X(), diffV.Y(), diffV.Z(), cubeShader)
	}

	for x := uint(0); x <= maxX; x++ {
		for y := uint(0); y <= maxY; y++ {
			for z := uint(0); z <= maxZ; z++ {
				loc := common.NewLocation(x, y, z)
				log.Printf("Checking %v", loc)

				cubes = append(cubes, NewCube(float32(x), float32(y), float32(z), 0.5, 0.5, 0.5, cubeShader))

				connected := (*lab).GetConnected(loc)

				var other common.Location

				if x < maxX {
					other = common.NewLocation(x+1, y, z)

					if isConnected(other, connected) {
						cubes = append(cubes, makeConnection(loc, other, cubeShader))
					}
				}

				if y < maxY {
					other = common.NewLocation(x, y+1, z)

					if isConnected(other, connected) {
						cubes = append(cubes, makeConnection(loc, other, cubeShader))
					}
				}

				if z < maxZ {
					other = common.NewLocation(x, y, z+1)

					if isConnected(other, connected) {
						cubes = append(cubes, makeConnection(loc, other, cubeShader))
					}
				}
			}
		}
	}

	return cubes
}
