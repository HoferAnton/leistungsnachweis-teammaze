package generator

import (
	"log"
	"math/rand"
	"testing"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func TestBreadthFirstGenerator_GenerateLabyrinth(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()

	var want common.Labyrinth = nil
	// act
	have, haveSteps := sut.GenerateLabyrinth(nil)
	// assert
	if want != have || haveSteps != nil {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_GenerateLabyrinth2(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxLoc := common.NewLocation(0, 0, 0)
	want := common.NewLabyrinth(maxLoc)
	wantSteps := []common.Pair{common.NewPair(maxLoc, Start), common.NewPair(maxLoc, Select)}
	// act
	have, haveSteps := sut.GenerateLabyrinth(maxLoc)
	// assert
	if !want.Compare(have) ||
		len(wantSteps) != len(haveSteps) ||
		!wantSteps[0].Compare(haveSteps[0]) {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_GenerateLabyrinth3(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxLoc := common.NewLocation(0, 0, 2)
	want := common.NewLabyrinth(maxLoc)
	have, haveSteps := sut.GenerateLabyrinth(maxLoc)
	// assert
	if want.Compare(have) ||
		len(haveSteps) == 3 {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_BackTrack(t *testing.T) {
	// arrange
	rand.Seed(0)

	sut := NewBreadthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}
	s := make([]common.Pair, 0)
	// act
	sut.iterate(maxLoc, &s)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_BackTrack3(t *testing.T) {
	// arrange
	rand.Seed(0)

	sut := NewBreadthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	// act
	s := make([]common.Pair, 0)
	sut.iterate(nil, &s)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_BackTrack4(t *testing.T) {
	// arrange
	rand.Seed(0)

	sut := NewBreadthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	s := make([]common.Pair, 0)
	// act
	sut.iterate(common.NewLocation(2, 2, 2), &s)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_BackTrack5(t *testing.T) {
	// arrange
	rand.Seed(0)

	sut := NewBreadthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[0] = true
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	s := make([]common.Pair, 0)
	// act
	sut.iterate(maxLoc, &s)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_GetUnvisited(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()

	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[0] = true
	// act
	have := getUnvisited(nil, &sut.lab, &sut.visited)
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_GetUnvisited2(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[0] = true
	// act
	have := getUnvisited(common.NewLocation(1, 1, 1), &sut.lab, &sut.visited)
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestBreadthFirstGenerator_GetUnvisited3(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(0, 0, 1)}
	// act
	have := getUnvisited(common.NewLocation(0, 0, 0), &sut.lab, &sut.visited)
	// assert
	if have == nil || len(want) != len(have) {
		t.Errorf("")
	}

	for i, haveLoc := range have {
		if !haveLoc.Compare(want[i]) {
			t.Errorf("")
		}
	}
}

func TestBreadthFirstGenerator_GetUnvisited4(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[7] = true
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(0, 0, 1)}
	// act
	have := getUnvisited(common.NewLocation(0, 0, 0), &sut.lab, &sut.visited)
	// assert
	if have == nil || len(want) != len(have) {
		t.Errorf("")
	}

	for i, haveLoc := range have {
		if !haveLoc.Compare(want[i]) {
			t.Errorf("")
		}
	}
}

func TestBreadthFirstGenerator_GetUnvisited5(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[7] = true
	sut.visited[0] = true
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(0, 0, 1)}
	// act
	have := getUnvisited(common.NewLocation(0, 0, 0), &sut.lab, &sut.visited)
	// assert
	if have == nil || len(want) != len(have) {
		t.Errorf("")
	}

	for i, haveLoc := range have {
		if !haveLoc.Compare(want[i]) {
			t.Errorf("")
		}
	}
}

func TestBreadthFirstGenerator_GetUnvisited6(t *testing.T) {
	// arrange
	sut := NewBreadthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]bool, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[0] = true
	sut.visited[4] = true
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0)}
	// act
	have := getUnvisited(common.NewLocation(0, 0, 0), &sut.lab, &sut.visited)
	// assert
	if have == nil || len(want) != len(have) {
		t.Errorf("\n%v should equal to %v\n in %v", want, have, sut.visited)
	}

	for i, haveLoc := range have {
		if !haveLoc.Compare(want[i]) {
			t.Errorf("")
		}
	}
}
