package generator

import (
	"math/rand"
	"time"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type DepthFirstGenerator struct {
	visited []common.Location
	lab     common.Labyrinth
}

const dCoordinate = 1

func NewDepthFirstGenerator() DepthFirstGenerator {
	dfg := DepthFirstGenerator{}
	dfg.visited = nil
	dfg.lab = nil

	return dfg
}

func (d DepthFirstGenerator) GenerateLabyrinth(furthestPoint common.Location) common.Labyrinth {
	if furthestPoint == nil {
		return nil
	}

	d.lab = common.NewLabyrinth(furthestPoint)
	maxX, maxY, maxZ := furthestPoint.As3DCoordinates()
	d.visited = make([]common.Location, (maxX+dCoordinate)*(maxY+dCoordinate)*(maxZ+dCoordinate))

	rand.Seed(time.Now().UnixNano())
	d.backtrack(
		common.NewLocation(
			uint(
				rand.Intn(
					int(maxX+dCoordinate))),
			uint(
				rand.Intn(
					int(maxY+dCoordinate))),
			uint(
				rand.Intn(
					int(maxZ+dCoordinate)))))

	return d.lab
}

func (d DepthFirstGenerator) backtrack(location common.Location) {
	if location == nil || !d.lab.CheckLocation(location) {
		panic("got nil")
	}

	thisX, thisY, thisZ := location.As3DCoordinates()
	thisIndex := common.GetIndex(thisX, thisY, thisZ, d.lab.GetMaxLocation())

	if d.visited[thisIndex] != nil {
		panic("got visited location")
	}

	d.visited[thisIndex] = location
	unvisited := d.getUnvisited(location)

	for len(unvisited) > 0 {
		next := unvisited[rand.Intn(len(unvisited))]

		d.lab.Connect(location, next)
		d.backtrack(next)

		unvisited = d.getUnvisited(location)
	}
}

func (d DepthFirstGenerator) getUnvisited(location common.Location) []common.Location {
	if location == nil || !d.lab.CheckLocation(location) {
		return nil
	}

	neighbors := d.lab.GetNeighbors(location)
	available := make([]common.Location, 0)

	for _, neighbor := range neighbors {
		x, y, z := neighbor.As3DCoordinates()

		if d.visited[common.GetIndex(x, y, z, d.lab.GetMaxLocation())] == nil {
			available = append(available, neighbor)
		}
	}

	return available
}
