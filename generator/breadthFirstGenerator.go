package generator

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"math/rand"
)

type BreadthFirstGenerator struct {
	visited []bool
	lab     common.Labyrinth
	steps   []common.Pair
}

func NewBreadthFirstGenerator() BreadthFirstGenerator {
	b := BreadthFirstGenerator{}
	b.visited = nil
	b.lab = nil
	b.steps = nil

	return b
}

func (b BreadthFirstGenerator) GenerateLabyrinth(furthestPoint common.Location) (common.Labyrinth, []common.Pair) {
	if furthestPoint == nil {
		return nil, nil
	}

	b.lab = common.NewLabyrinth(furthestPoint)
	maxX, maxY, maxZ := furthestPoint.As3DCoordinates()
	b.visited = make([]bool, (maxX+dCoordinate)*(maxY+dCoordinate)*(maxZ+dCoordinate))
	b.steps = make([]common.Pair, 0)
	startLocation := getRandomizedStart(furthestPoint)
	b.steps = append(b.steps, common.NewPair(startLocation, Start))

	b.iterate(startLocation)

	return b.lab, b.steps
}

func (b BreadthFirstGenerator) iterate(startLocation common.Location) {
	workList := []common.Location{startLocation}

	for len(workList) != 0 {
		i := rand.Intn(len(workList))
		e := workList[i]
		workList = append(workList[:i], workList[i+1:]...)
		eX, eY, eZ := e.As3DCoordinates()
		eIndex := common.GetIndex(eX, eY, eZ, b.lab.GetMaxLocation())
		b.visited[eIndex] = true
		b.steps = append(b.steps, common.NewPair(e, Select))

		for _, n := range getUnvisited(e, &b.lab, &b.visited) {
			nX, nY, nZ := n.As3DCoordinates()
			nIndex := common.GetIndex(nX, nY, nZ, b.lab.GetMaxLocation())
			b.visited[nIndex] = true
			b.lab.Connect(e, n)
			workList = append(workList, n)
			b.steps = append(b.steps, common.NewPair(n, Add))
		}
	}
}
