package solver

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

// Uses recursive depth first search (not Concurrent)
func RecursiveSolver(lab common.Labyrinth, from common.Location, to common.Location) []common.Location {
	return rdfs(lab, from, to, []common.Location{from})
}

func rdfs(lab common.Labyrinth, from common.Location, to common.Location,
	dontTouch []common.Location) []common.Location {
	if from.Compare(to) {
		return []common.Location{to}
	}

	for _, neighbor := range lab.GetConnected(from) {
		if neighbor.Compare(from) || contains(dontTouch, neighbor) {
			continue
		}

		dontTouch = append(dontTouch, neighbor)
		if result := rdfs(lab, neighbor, to, dontTouch); result != nil {
			//return append(result, from)
			return append([]common.Location{from}, result...)
		}
	}

	return nil
}
