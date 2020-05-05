package generator

import (
	"math/rand"
	"time"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type DepthFirstGenerator struct {
	visited []bool
	lab     common.Labyrinth
	steps   []common.Pair
}

const (
	dCoordinate = 1
	start       = "START"
	discover    = "DISCOVER"
	backtrack   = "BACKTRACK"
)

func NewDepthFirstGenerator() DepthFirstGenerator {
	dfg := DepthFirstGenerator{}
	dfg.visited = nil
	dfg.lab = nil
	dfg.steps = nil

	return dfg
}

func (d DepthFirstGenerator) GenerateLabyrinth(furthestPoint common.Location) (common.Labyrinth, []common.Pair) {
	if furthestPoint == nil {
		return nil, nil
	}

	d.lab = common.NewLabyrinth(furthestPoint)
	maxX, maxY, maxZ := furthestPoint.As3DCoordinates()
	d.visited = make([]bool, (maxX+dCoordinate)*(maxY+dCoordinate)*(maxZ+dCoordinate))
	d.steps = make([]common.Pair, 0)

	rand.Seed(time.Now().UnixNano())

	startLocation := common.NewLocation(
		uint(
			rand.Intn(
				int(maxX+dCoordinate))),
		uint(
			rand.Intn(
				int(maxY+dCoordinate))),
		uint(
			rand.Intn(
				int(maxZ+dCoordinate))))
	d.steps = append(d.steps, common.NewPair(startLocation, start))

	d.backtrack(startLocation)

	return d.lab, d.steps
}

func (d DepthFirstGenerator) backtrack(location common.Location) {
	if location == nil || !d.lab.CheckLocation(location) {
		panic("got nil")
	}

	thisX, thisY, thisZ := location.As3DCoordinates()
	thisIndex := common.GetIndex(thisX, thisY, thisZ, d.lab.GetMaxLocation())

	if d.visited[thisIndex] {
		panic("got visited location")
	}

	d.visited[thisIndex] = true
	unvisited := d.getUnvisited(location)

	for len(unvisited) > 0 {
		next := unvisited[rand.Intn(len(unvisited))]
		d.steps = append(d.steps, common.NewPair(next, discover))

		d.lab.Connect(location, next)
		d.backtrack(next)

		d.steps = append(d.steps, common.NewPair(next, backtrack))
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

		if !d.visited[common.GetIndex(x, y, z, d.lab.GetMaxLocation())] {
			available = append(available, neighbor)
		}
	}

	return available
}
