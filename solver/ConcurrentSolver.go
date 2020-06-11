package solver

import (
	"sync"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

const (
	functionReady = 1
	functionDone  = -1
)

func ConcurrentSolver(lab common.Labyrinth, from common.Location, to common.Location, trust bool) []common.Location {

	var (
		result []common.Location
		wg     sync.WaitGroup
	)

	wg.Add(functionReady)
	go runner(&lab, from, to, from, []common.Location{}, &wg, &result, trust)

	wg.Wait()
	return result
}

func runner(lab *common.Labyrinth, from common.Location, to common.Location, previous common.Location,
	path []common.Location, wg *sync.WaitGroup, result *[]common.Location, trust bool) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Add(functionDone)
	}(wg)

	if *result != nil {
		return
	}

	newPath := append(path, from)

	if from.Compare(to) {
		*result = newPath

		return
	}

	for _, neighbor := range (*lab).GetConnected(from) {
		if neighbor.Compare(previous) {
			continue
		}

		if !trust {
			if contains(newPath, neighbor) {
				continue
			}
		}

		wg.Add(functionReady)
		go runner(lab, neighbor, to, from, newPath, wg, result, trust)
	}

}
