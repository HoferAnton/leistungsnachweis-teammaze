package solver

import (
	"fmt"
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type LabSolver interface {
	SolveLabyrinth(labyrinth common.Labyrinth, from common.Location, to common.Location)
}

// solver 1 depth first search (recursive) not paralysed
// still (suboptimal) i guess
func DFS(lab common.Labyrinth, from common.Location, to common.Location) []common.Location {
	return rdfs(lab, from, to, []common.Location{})
}

func rdfs(lab common.Labyrinth, from common.Location, to common.Location, dontTouch []common.Location) []common.Location {

	fmt.Printf("From %v | To: %v \n", from, to)

	if from.Compare(to) {
		return append([]common.Location{to})
	}

	for _, neighbor := range lab.GetConnected(from) {
		if neighbor.Compare(from) || contains(dontTouch, neighbor) {
			continue
		}
		dontTouch = append(dontTouch, neighbor)
		if result := rdfs(lab, neighbor, to, dontTouch); result != nil {
			return append(result, from)
		}
	}

	return nil
}

func contains(l []common.Location, e common.Location) bool {
	for _, s := range l {
		if s.Compare(e) {
			return true
		}
	}
	return false
}
