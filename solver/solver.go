package solver

import (
	"sync"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

type LabSolver interface {
	SolveLabyrinth(labyrinth common.Labyrinth, from common.Location, to common.Location)
}

// Uses recursive depth first search (not Concurrent)
func RecursiveSolver(lab common.Labyrinth, from common.Location, to common.Location) []common.Location {
	return rdfs(lab, from, to, []common.Location{})
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

const (
	functionReady = 1
	functionDone  = -1
)

// It starts a new goroutine for each new potential sub-path
func ConcurrentSolver(lab common.Labyrinth, from common.Location, to common.Location) []common.Location {
	var (
		result []common.Location
		wg     sync.WaitGroup
	)

	wg.Add(functionReady)
	pd(lab, from, to, []common.Location{}, &wg, &result)

	wg.Wait()

	return result
}

func pd(lab common.Labyrinth, from common.Location, to common.Location,
	way []common.Location, wg *sync.WaitGroup, result *[]common.Location) {

	defer func(wg *sync.WaitGroup) {
		wg.Add(functionDone)
	}(wg)

	if *result != nil {
		return
	}

	way = append(way, from)

	if from.Compare(to) {
		*result = way

		return
	}

	for _, neighbor := range lab.GetConnected(from) {
		if contains(way, neighbor) {
			continue
		}

		wg.Add(functionReady)

		go pd(lab, neighbor, to, way, wg, result)
	}
}

func contains(l []common.Location, e common.Location) bool {
	for _, s := range l {
		if s.Compare(e) {
			return true
		}
	}

	return false
}
