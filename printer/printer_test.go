package printer

import (
	"errors"
	"testing"
)
import "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

func TestPrint2D_nil(t *testing.T) {
	// arrange
	wantS := ""
	wantE := errors.New("got nil")
	// act
	haveS, haveE := Print2D(nil)

	// assert
	if wantS != haveS {
		t.Errorf("\n%v should be equal to \n%v", haveS, wantS)
	}

	if haveE == nil || haveE.Error() != wantE.Error() {
		t.Errorf("\n%v should be equal to \n%v", haveE, wantE)
	}
}

func TestPrint2D_interpretFloor_nil(t *testing.T) {
	// arrange
	wantS := ""
	wantE := errors.New("got nil")
	// act
	haveS, haveE := interpretFloor(nil, 0)

	// assert
	if wantS != haveS {
		t.Errorf("\n%v should be equal to \n%v", haveS, wantS)
	}

	if haveE == nil || haveE.Error() != wantE.Error() {
		t.Errorf("\n%v should be equal to \n%v", haveE, wantE)
	}
}

func TestPrint2D_interpretFloor_zOverflow(t *testing.T) {
	// arrange
	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))
	wantS := ""
	wantE := errors.New("z out of range")
	// act
	haveS, haveE := interpretFloor(lab, 1)

	// assert
	if wantS != haveS {
		t.Errorf("\n%v should be equal to \n%v", haveS, wantS)
	}

	if haveE == nil || haveE.Error() != wantE.Error() {
		t.Errorf("\n%v should be equal to \n%v", haveE, wantE)
	}
}

func TestPrint2D_interpretLine_nil(t *testing.T) {
	// arrange
	wantS := ""
	wantE := errors.New("got nil")
	// act
	haveS, haveE := interpretLine(nil, 0, 0)

	// assert
	if wantS != haveS {
		t.Errorf("\n%v should be equal to \n%v", haveS, wantS)
	}

	if haveE == nil || haveE.Error() != wantE.Error() {
		t.Errorf("\n%v should be equal to \n%v", haveE, wantE)
	}
}

func TestPrint2D_interpretFloorLine_zOverflow(t *testing.T) {
	// arrange
	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))
	wantS := ""
	wantE := errors.New("z out of range")
	// act
	haveS, haveE := interpretLine(lab, 0, 1)

	// assert
	if wantS != haveS {
		t.Errorf("\n%v should be equal to \n%v", haveS, wantS)
	}

	if haveE == nil || haveE.Error() != wantE.Error() {
		t.Errorf("\n%v should be equal to \n%v", haveE, wantE)
	}
}

func TestPrint2D_interpretLine_yOverflow(t *testing.T) {
	// arrange
	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))
	wantS := ""
	wantE := errors.New("y out of range")
	// act
	haveS, haveE := interpretLine(lab, 1, 0)

	// assert
	if wantS != haveS {
		t.Errorf("\n%v should be equal to \n%v", haveS, wantS)
	}

	if haveE == nil || haveE.Error() != wantE.Error() {
		t.Errorf("\n%v should be equal to \n%v", haveE, wantE)
	}
}

func TestPrint2D_000(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)

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
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(0, 0, 0))
	lab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(0, 0, 2))

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
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0))

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + noWall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + wall + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(0, 1, 0))

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + noWall + perimeter + nl
	want += perimeter + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)

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
	have, _ := Print2D(lab)

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
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 1, 0))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 1, 2))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 0, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(1, 2, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(0, 1, 1))
	lab.Connect(common.NewLocation(1, 1, 1), common.NewLocation(2, 1, 1))

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
	have, _ := Print2D(lab)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}
