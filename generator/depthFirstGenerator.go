package generator

import (
	"math/rand"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type DepthFirstGenerator struct {
	visited []bool
	lab     common.Labyrinth
	steps   []common.Pair
}

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
	startLocation := getRandomizedStart(furthestPoint)
	d.steps = append(d.steps, common.NewPair(startLocation, Start))

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
	unvisited := getUnvisited(location, &d.lab, &d.visited)

	for len(unvisited) > 0 {
		next := unvisited[rand.Intn(len(unvisited))]
		d.steps = append(d.steps, common.NewPair(next, Discover))

		d.lab.Connect(location, next)
		d.backtrack(next)

		d.steps = append(d.steps, common.NewPair(next, Backtrack))
		unvisited = getUnvisited(location, &d.lab, &d.visited)
	}
}
