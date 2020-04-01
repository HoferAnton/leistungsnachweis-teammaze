package generator

import (
	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
	"log"
	"math/rand"
	"testing"
)

func TestDepthFirstGenerator_GenerateLabyrinth(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	var want common.Labyrinth = nil
	// act
	have := sut.GenerateLabyrinth(nil)
	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_GenerateLabyrinth2(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxLoc := common.NewLocation(0, 0, 0)
	want := common.NewLabyrinth(maxLoc)
	// act
	have := sut.GenerateLabyrinth(maxLoc)
	// assert
	if !want.Compare(have) {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_BackTrack(t *testing.T) {
	// arrange
	rand.Seed(0)
	sut := NewDepthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}
	// act
	sut.backtrack(maxLoc)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_BackTrack2(t *testing.T) {
	// arrange
	rand.Seed(0)
	sut := NewDepthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)

	wantLab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0))
	wantLab.Connect(common.NewLocation(1, 0, 0), common.NewLocation(1, 1, 0))
	wantLab.Connect(common.NewLocation(1, 1, 0), common.NewLocation(0, 1, 0))

	wantLab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(0, 0, 0))
	wantLab.Connect(common.NewLocation(1, 0, 1), common.NewLocation(1, 0, 0))

	wantLab.Connect(common.NewLocation(0, 1, 1), common.NewLocation(0, 0, 1))
	wantLab.Connect(common.NewLocation(0, 1, 1), common.NewLocation(1, 1, 1))

	wantVisited := []common.Location{
		common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0), common.NewLocation(1, 1, 0),
		common.NewLocation(0, 0, 1), common.NewLocation(1, 0, 1),
		common.NewLocation(0, 1, 1), common.NewLocation(1, 1, 1),
	}
	// act
	sut.backtrack(maxLoc)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("%v should equal to %v\n", wantVisited, sut.visited)
	}

	for i, visited := range wantVisited {
		if !visited.Compare(sut.visited[i]) {
			t.Errorf("%v should equal to %v\n", visited, sut.visited[i])
		}
	}
}

func TestDepthFirstGenerator_BackTrack3(t *testing.T) {
	// arrange
	rand.Seed(0)
	sut := NewDepthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	// act
	sut.backtrack(nil)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_BackTrack4(t *testing.T) {
	// arrange
	rand.Seed(0)
	sut := NewDepthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	// act
	sut.backtrack(common.NewLocation(2, 2, 2))
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_BackTrack5(t *testing.T) {
	// arrange
	rand.Seed(0)
	sut := NewDepthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[0] = maxLoc
	wantLab := common.NewLabyrinth(maxLoc)
	wantVisited := []common.Location{maxLoc}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	// act
	sut.backtrack(maxLoc)
	// assert
	if !wantLab.Compare(sut.lab) {
		t.Errorf("")
	}

	if len(wantVisited) != len(sut.visited) {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_GetUnvisited(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited = append(sut.visited, maxLoc)
	// act
	have := sut.getUnvisited(nil)
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_GetUnvisited2(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited = append(sut.visited, maxLoc)
	// act
	have := sut.getUnvisited(common.NewLocation(1, 1, 1))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestDepthFirstGenerator_GetUnvisited3(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(0, 0, 1)}
	// act
	have := sut.getUnvisited(common.NewLocation(0, 0, 0))
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

func TestDepthFirstGenerator_GetUnvisited4(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[7] = maxLoc
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(0, 0, 1)}
	// act
	have := sut.getUnvisited(common.NewLocation(0, 0, 0))
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

func TestDepthFirstGenerator_GetUnvisited5(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[7] = maxLoc
	sut.visited[0] = common.NewLocation(0, 0, 0)
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0),
		common.NewLocation(0, 0, 1)}
	// act
	have := sut.getUnvisited(common.NewLocation(0, 0, 0))
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

func TestDepthFirstGenerator_GetUnvisited6(t *testing.T) {
	// arrange
	sut := NewDepthFirstGenerator()
	maxX := uint(1)
	maxY := uint(1)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	sut.lab = common.NewLabyrinth(maxLoc)
	sut.visited = make([]common.Location, (maxX+1)*(maxY+1)*(maxZ+1))
	sut.visited[0] = common.NewLocation(0, 0, 0)
	sut.visited[4] = common.NewLocation(0, 0, 1)
	want := []common.Location{
		common.NewLocation(1, 0, 0),
		common.NewLocation(0, 1, 0)}
	// act
	have := sut.getUnvisited(common.NewLocation(0, 0, 0))
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
