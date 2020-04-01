package generator

import (
	. "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"math/rand"
	"time"
)

type DepthFirstGenerator struct {
	visited []Location
	lab     Labyrinth
}

const dCoordinate = 1

func NewDepthFirstGenerator() DepthFirstGenerator {
	dfg := DepthFirstGenerator{}
	dfg.visited = nil
	dfg.lab = nil

	return dfg
}

func (d DepthFirstGenerator) GenerateLabyrinth(furthestPoint Location) Labyrinth {
	if furthestPoint == nil {
		return nil
	}

	d.lab = NewLabyrinth(furthestPoint)
	maxX, maxY, maxZ := furthestPoint.As3DCoordinates()
	d.visited = make([]Location, (maxX+1)*(maxY+1)*(maxZ+1))

	rand.Seed(time.Now().UnixNano())
	d.backtrack(
		NewLocation(
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

func (d DepthFirstGenerator) backtrack(location Location) {
	if location == nil || !d.lab.CheckLocation(location) {
		panic("got nil")
	}

	thisX, thisY, thisZ := location.As3DCoordinates()
	thisIndex := GetIndex(thisX, thisY, thisZ, d.lab.GetMaxLocation())

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

func (d DepthFirstGenerator) getUnvisited(location Location) []Location {
	if location == nil || !d.lab.CheckLocation(location) {
		return nil
	}

	neighbors := d.lab.GetNeighbors(location)
	available := make([]Location, 0)

	for _, neighbor := range neighbors {
		x, y, z := neighbor.As3DCoordinates()

		if d.visited[GetIndex(x, y, z, d.lab.GetMaxLocation())] == nil {
			available = append(available, neighbor)
		}
	}

	return available
}
