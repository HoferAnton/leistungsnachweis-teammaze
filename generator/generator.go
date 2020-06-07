package generator

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"math/rand"
	"time"
)

type LabGenerator interface {
	GenerateLabyrinth(furthestPoint common.Location) (common.Labyrinth, []common.Pair)
}

const (
	dCoordinate = 1
	Start       = "START"
	Discover    = "DISCOVER"
	Backtrack   = "BACKTRACK"
	Select      = "SELECT"
	Add         = "ADD"
)

func getRandomizedStart(maxLoc common.Location) common.Location {
	maxX, maxY, maxZ := maxLoc.As3DCoordinates()

	rand.Seed(time.Now().UnixNano())
	return common.NewLocation(
		uint(
			rand.Intn(
				int(maxX+dCoordinate))),
		uint(
			rand.Intn(
				int(maxY+dCoordinate))),
		uint(
			rand.Intn(
				int(maxZ+dCoordinate))))
}

func getUnvisited(location common.Location, lab *common.Labyrinth, visited *[]bool) []common.Location {
	if location == nil || lab == nil || visited == nil || !(*lab).CheckLocation(location) {
		return nil
	}

	neighbors := (*lab).GetNeighbors(location)
	available := make([]common.Location, 0)

	for _, neighbor := range neighbors {
		x, y, z := neighbor.As3DCoordinates()

		if !(*visited)[common.GetIndex(x, y, z, (*lab).GetMaxLocation())] {
			available = append(available, neighbor)
		}
	}

	return available
}
