package solver

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

const (
	Add    = "ADD"
	Remove = "REMOVE"
)

// Uses recursive depth first search (not Concurrent)
func RecursiveSolver(lab common.Labyrinth, from common.Location, to common.Location, trust bool) []common.Location {
	var steps []common.Pair
	var dontTouch []common.Location
	if !trust {
		dontTouch = []common.Location{from}
	}

	return rdfs(&lab, from, to, from, dontTouch, &steps)
}

// Uses recursive depth first search (not Concurrent)
func RecursiveSolverSteps(lab common.Labyrinth, from common.Location,
	to common.Location, trust bool) ([]common.Location, []common.Pair) {
	steps := []common.Pair{common.NewPair(from, Add)}
	var dontTouch []common.Location
	if !trust {
		dontTouch = []common.Location{from}
	}

	path := rdfs(&lab, from, to, from, dontTouch, &steps)
	return path, steps
}

func rdfs(lab *common.Labyrinth, from common.Location, to common.Location, previous common.Location,
	dontTouch []common.Location, steps *[]common.Pair) []common.Location {
	if from.Compare(to) {
		return []common.Location{to}
	}

	for _, neighbor := range (*lab).GetConnected(from) {
		if neighbor.Compare(previous) || contains(dontTouch, neighbor) {
			continue
		}

		if dontTouch != nil {
			dontTouch = append(dontTouch, neighbor)
		}

		if len(*steps) > 0 {
			*steps = append(*steps, common.NewPair(neighbor, Add))
		}

		if result := rdfs(lab, neighbor, to, from, dontTouch, steps); result != nil {
			return append([]common.Location{from}, result...)
		}

		if len(*steps) > 0 {
			*steps = append(*steps, common.NewPair(neighbor, Remove))
		}
	}

	return nil
}
