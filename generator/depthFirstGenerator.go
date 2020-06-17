package generator

import (
	"math/rand"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type DepthFirstGenerator struct {
	visited []bool
	lab     common.Labyrinth
}

func NewDepthFirstGenerator() DepthFirstGenerator {
	dfg := DepthFirstGenerator{}
	dfg.visited = nil
	dfg.lab = nil

	return dfg
}

func (d DepthFirstGenerator) GenerateLabyrinth(furthestPoint common.Location) (common.Labyrinth, []common.Pair) {
	if furthestPoint == nil {
		return nil, nil
	}

	d.lab = common.NewLabyrinth(furthestPoint)
	maxX, maxY, maxZ := furthestPoint.As3DCoordinates()
	d.visited = make([]bool, (maxX+dCoordinate)*(maxY+dCoordinate)*(maxZ+dCoordinate))
	s := make([]common.Pair, 0)
	steps := &s
	startLocation := getRandomizedStart(furthestPoint)
	*steps = append(*steps, common.NewPair(startLocation, Start))

	d.backtrack(startLocation, steps)

	return d.lab, *steps
}

func (d DepthFirstGenerator) backtrack(location common.Location, steps *[]common.Pair) {
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
		*steps = append(*steps, common.NewPair(next, Discover))

		d.lab.Connect(location, next)
		d.backtrack(next, steps)

		*steps = append(*steps, common.NewPair(next, Backtrack))
		unvisited = getUnvisited(location, &d.lab, &d.visited)
	}
}
