package solver

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type LabSolver interface {
	SolveLabyrinth(labyrinth common.Labyrinth, from common.Location, to common.Location)
}

func contains(l []common.Location, e common.Location) bool {
	for _, s := range l {
		if s.Compare(e) {
			return true
		}
	}
	return false
}

func removeFirstOccurrence(l []common.Location, e common.Location) []common.Location {
	for i, s := range l {
		if s.Compare(e) {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
