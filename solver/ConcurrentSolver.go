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

	go runner(&lab, from, to, from, []common.Location{}, &wg, &result, trust, nil)

	wg.Wait()

	return result
}

func ConcurrentSolverSteps(lab common.Labyrinth, from common.Location, to common.Location,
	trust bool) ([]common.Location, []common.Pair) {
	var (
		result []common.Location
		steps  []common.Pair
		wgw    sync.WaitGroup
		wgs    sync.WaitGroup
	)

	c := make(chan common.Location)

	wgw.Add(functionReady)

	go runner(&lab, from, to, from, []common.Location{}, &wgw, &result, trust, c)

	wgs.Add(functionReady)

	go func(c chan common.Location, steps *[]common.Pair, wg *sync.WaitGroup) {
		defer (*wg).Add(functionDone)

		for loc := range c {
			*steps = append(*steps, common.NewPair(loc, Visited))
		}
	}(c, &steps, &wgs)

	wgw.Wait()
	close(c)
	wgs.Wait()

	return result, steps
}

func runner(lab *common.Labyrinth, from common.Location, to common.Location, previous common.Location,
	path []common.Location, wg *sync.WaitGroup, result *[]common.Location, trust bool, c chan common.Location) {
	defer (*wg).Add(functionDone)

	if *result != nil {
		return
	}

	newPath := make([]common.Location, len(path)+1)
	copy(newPath, path)
	newPath[len(newPath)-1] = from

	if c != nil {
		c <- from
	}

	if from.Compare(to) {
		*result = newPath

		return
	}

	for _, neighbor := range (*lab).GetConnected(from) {
		if neighbor.Compare(previous) || (!trust && contains(newPath, neighbor)) {
			continue
		}

		wg.Add(functionReady)

		go runner(lab, neighbor, to, from, newPath, wg, result, trust, c)
	}
}
