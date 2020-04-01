package solver

import (
	"fmt"
	"sync"

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

func rdfs(lab common.Labyrinth, from common.Location, to common.Location,
	dontTouch []common.Location) []common.Location {
	fmt.Printf("From %v | To: %v \n", from, to)

	if from.Compare(to) {
		return []common.Location{to}
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

// solver 2
func PDS(lab common.Labyrinth, from common.Location, to common.Location) []common.Location {
	var (
		result []common.Location
		wg     sync.WaitGroup
	)

	wg.Add(1)
	pd(lab, from, to, []common.Location{}, &wg, &result)

	wg.Wait()

	return result
}

func pd(lab common.Labyrinth, from common.Location, to common.Location,
	way []common.Location, wg *sync.WaitGroup, result *[]common.Location) {
	if *result != nil {
		wg.Add(-1)
		return
	}

	fmt.Printf("From %v | To: %v | Way: %v \n", from, to, way)

	way = append(way, from)

	if from.Compare(to) {
		*result = way

		wg.Add(-1)

		return
	}

	for _, neighbor := range lab.GetConnected(from) {
		if contains(way, neighbor) {
			continue
		}

		wg.Add(1)

		go pd(lab, neighbor, to, way, wg, result)
	}

	wg.Add(-1)
}

func contains(l []common.Location, e common.Location) bool {
	for _, s := range l {
		if s.Compare(e) {
			return true
		}
	}

	return false
}
