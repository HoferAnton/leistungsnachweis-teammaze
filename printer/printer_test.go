package printer

import "testing"
import . "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

func TestPrint2D_nil(t *testing.T) {
	// arrange
	want := ""
	// act
	have := Print2D(nil)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_000(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_002_empty(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_002_Tower(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	lab.Connect(NewLocation(0, 0, 1), NewLocation(0, 0, 0))
	lab.Connect(NewLocation(0, 0, 1), NewLocation(0, 0, 2))

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellDown + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + cellTower + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + cellUp + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_100_empty(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_100_connected(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	lab.Connect(NewLocation(0, 0, 0), NewLocation(1, 0, 0))

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + noWall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_010_empty(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(1)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + wall + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_010_connected(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(1)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	lab.Connect(NewLocation(0, 0, 0), NewLocation(0, 1, 0))

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + noWall + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_222_empty(t *testing.T) {
	// arrange
	maxX := uint(2)
	maxY := uint(2)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_222_star(t *testing.T) {
	// arrange
	maxX := uint(2)
	maxY := uint(2)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	lab.Connect(NewLocation(1, 1, 1), NewLocation(1, 1, 0))
	lab.Connect(NewLocation(1, 1, 1), NewLocation(1, 1, 2))
	lab.Connect(NewLocation(1, 1, 1), NewLocation(1, 0, 1))
	lab.Connect(NewLocation(1, 1, 1), NewLocation(1, 2, 1))
	lab.Connect(NewLocation(1, 1, 1), NewLocation(0, 1, 1))
	lab.Connect(NewLocation(1, 1, 1), NewLocation(2, 1, 1))

	want := perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellDown + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + noWall + post + wall + perimeter + nl
	want += perimeter + cellNormal + noWall + cellTower + noWall + cellNormal + perimeter + nl
	want += perimeter + wall + post + noWall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellUp + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}
