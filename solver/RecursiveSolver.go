package solver

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

// Uses recursive depth first search (not Concurrent)
func RecursiveSolver(lab common.Labyrinth, from common.Location, to common.Location, trust bool) []common.Location {
	var dontTouch []common.Location
	if !trust {
		dontTouch = []common.Location{from}
	}

	return rdfs(&lab, from, to, from, dontTouch)
}

func rdfs(lab *common.Labyrinth, from common.Location, to common.Location,
	previous common.Location, dontTouch []common.Location) []common.Location {
	if from.Compare(to) {
		return []common.Location{to}
	}

	for _, neighbor := range (*lab).GetConnected(from) {
		if neighbor.Compare(previous) || contains(dontTouch, neighbor) {
			continue
		}

		dontTouch = append(dontTouch, neighbor)
		if result := rdfs(lab, neighbor, to, from, dontTouch); result != nil {
			//return append(result, from)
			return append([]common.Location{from}, result...)
		}
	}

	return nil
}
