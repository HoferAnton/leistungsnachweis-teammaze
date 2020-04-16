package printer

import (
	"errors"
	"testing"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func NoPanicFail(t *testing.T, message string) {
	if r := recover(); r == nil {
		t.Errorf(message)
	}
}

func TestPrint2D_nil(t *testing.T) {
	// arrange
	wantS := ""
	wantE := errors.New("no maze given")
	// act
	haveS, haveE := Print2D(nil, nil)

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
	defer NoPanicFail(t, "missing panic for lab = nil")
	// act
	interpretFloor(nil, 0, nil)
}

func TestPrint2D_interpretFloor_zOverflow(t *testing.T) {
	// arrange
	defer NoPanicFail(t, "missing panic for z out of bounds")

	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))

	// act
	interpretFloor(lab, 1, nil)
}

func TestPrint2D_interpretLine_nil(t *testing.T) {
	// arrange
	defer NoPanicFail(t, "missing panic for lab = nil")

	// act
	interpretLine(nil, 0, 0, nil)
}

func TestPrint2D_interpretFloorLine_zOverflow(t *testing.T) {
	// arrange
	defer NoPanicFail(t, "missing panic for z out of bounds")

	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))

	// act
	interpretLine(lab, 0, 1, nil)
}

func TestPrint2D_interpretLine_yOverflow(t *testing.T) {
	// arrange
	defer NoPanicFail(t, "missing panic for y out of bounds")

	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))

	// act
	interpretLine(lab, 1, 0, nil)
}

func TestPrint2D_interpretCell_labNil(t *testing.T) {
	// arrange
	defer NoPanicFail(t, "missing panic for nil")

	loc := common.NewLocation(0, 0, 0)

	// act
	interpretCell(nil, loc, nil)
}

func TestPrint2D_interpretCell_locNil(t *testing.T) {
	// arrange
	defer NoPanicFail(t, "missing panic for nil")

	lab := common.NewLabyrinth(common.NewLocation(0, 0, 0))

	// act
	interpretCell(lab, nil, nil)
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
	have, _ := Print2D(lab, nil)

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
	have, _ := Print2D(lab, nil)

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
	want += perimeter + cellUpDown + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + cellUp + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab, nil)

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
	have, _ := Print2D(lab, nil)

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
	have, _ := Print2D(lab, nil)

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
	have, _ := Print2D(lab, nil)

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
	have, _ := Print2D(lab, nil)

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
	have, _ := Print2D(lab, nil)

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
	want += perimeter + cellNormal + noWall + cellUpDown + noWall + cellNormal + perimeter + nl
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
	have, _ := Print2D(lab, nil)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////// with paths

func TestPrint2D_path_normal(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0))
	path := []common.Location{common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0)}

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + pathNormal + noWallPath + pathNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_path_downCould(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(1, 0, 1))
	lab.Connect(common.NewLocation(1, 0, 1), common.NewLocation(1, 0, 0))
	path := []common.Location{common.NewLocation(0, 0, 1), common.NewLocation(1, 0, 1)}

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + pathNormal + noWallPath + pathDownCould + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellUp + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_path_downGo_upGo_upDownGO(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(2)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(0, 0, 0))
	lab.Connect(common.NewLocation(0, 0, 1), common.NewLocation(0, 0, 2))
	path := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(0, 0, 1),
		common.NewLocation(0, 0, 2),
	}

	want := perimeter + perimeter + perimeter + nl
	want += perimeter + pathDownGo + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + pathUpDownGo + perimeter + nl
	want += perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + nl
	want += perimeter + pathUpGo + perimeter + nl
	want += perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_path_upCloud(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(1)
	maxLoc := common.NewLocation(maxX, maxY, maxZ)
	lab := common.NewLabyrinth(maxLoc)
	lab.Connect(common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0))
	lab.Connect(common.NewLocation(1, 0, 1), common.NewLocation(1, 0, 0))
	path := []common.Location{common.NewLocation(0, 0, 0), common.NewLocation(1, 0, 0)}

	want := perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellDown + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + pathNormal + noWallPath + pathUpCould + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_path_upDownCould(t *testing.T) {
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
	path := []common.Location{
		common.NewLocation(1, 1, 1),
		common.NewLocation(0, 1, 1),
		common.NewLocation(2, 1, 1),
	}

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
	want += perimeter + pathNormal + noWallPath + pathUpDownCould + noWallPath + pathNormal + perimeter + nl
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
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_path_upCouldDownGo(t *testing.T) { // arrange
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
	path := []common.Location{
		common.NewLocation(1, 1, 1),
		common.NewLocation(2, 1, 1),
		common.NewLocation(1, 1, 0),
	}

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
	want += perimeter + cellNormal + noWall + pathUpCouldDownGo + noWallPath + pathNormal + perimeter + nl
	want += perimeter + wall + post + noWall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + pathUpGo + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl

	// act
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}

func TestPrint2D_path_upGoDownCould(t *testing.T) { // arrange
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
	path := []common.Location{
		common.NewLocation(1, 1, 1),
		common.NewLocation(1, 2, 1),
		common.NewLocation(1, 1, 2),
	}

	want := perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + pathDownGo + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + wall + post + wall + perimeter + nl
	want += perimeter + cellNormal + wall + cellNormal + wall + cellNormal + perimeter + nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += nl
	want += perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + perimeter + nl
	want += perimeter + cellNormal + wall + pathNormal + wall + cellNormal + perimeter + nl
	want += perimeter + wall + post + noWallPath + post + wall + perimeter + nl
	want += perimeter + cellNormal + noWall + pathUpGoDownCould + noWall + cellNormal + perimeter + nl
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
	have, _ := Print2D(lab, path)

	// assert
	if want != have {
		t.Errorf("\n%v should be equal to \n%v", have, want)
	}
}
