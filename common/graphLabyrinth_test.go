package common

import (
	"testing"
)

func TestGraphLabyrinth_ctorSmall(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	// act
	lab := NewLabyrinth(maxLoc)
	// assert
	if !newNode(maxLoc).hardCompare(lab.getNode(maxLoc)) {
		t.Errorf("%v should be equal to %v", newNode(maxLoc), lab.getNode(maxLoc))
	}
}

func TestGraphLabyrinth_ctorNormal(t *testing.T) {
	// arrange
	maxX := uint(64)
	maxY := uint(64)
	maxZ := uint(64)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	// act
	lab := NewLabyrinth(maxLoc)
	// assert
	for z := uint(0); z <= maxZ; z++ {
		for y := uint(0); y <= maxY; y++ {
			for x := uint(0); x <= maxX; x++ {
				loc := NewLocation(x, y, z)
				if !newNode(loc).hardCompare(lab.getNode(loc)) {
					t.Errorf("%v should be equal to %v by possition %v", newNode(loc), lab.getNode(loc), loc)
				}
			}
		}
	}
}

func TestGraphLabyrinth_GetMaxLocation(t *testing.T) {
	// arrange
	maxX := uint(64)
	maxY := uint(64)
	maxZ := uint(64)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	// act
	loc := lab.GetMaxLocation()
	// assert
	if !maxLoc.Compare(loc) {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNodeMaxUsable(t *testing.T) {
	// arrange
	maxX := uint(150)
	maxY := uint(150)
	maxZ := uint(150)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	// act
	lab := NewLabyrinth(maxLoc)
	// assert
	for z := uint(0); z <= maxZ; z++ {
		for y := uint(0); y <= maxY; y++ {
			for x := uint(0); x <= maxX; x++ {
				loc := NewLocation(x, y, z)
				if !newNode(loc).hardCompare(lab.getNode(loc)) {
					t.Errorf("%v should be equal to %v by possition %v", newNode(loc), lab.getNode(loc), loc)
				}
			}
		}
	}
}

func TestGraphLabyrinth_GetNodeNil(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	locOutOfRange := NewLocation(maxX, maxY, 1)
	// act
	lab := NewLabyrinth(maxLoc)
	// assert
	if lab.getNode(locOutOfRange) != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNodeNil1(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	locOutOfRange := NewLocation(maxX, 1, maxZ)
	// act
	lab := NewLabyrinth(maxLoc)
	// assert
	if lab.getNode(locOutOfRange) != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNodeNil2(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	locOutOfRange := NewLocation(1, maxY, maxZ)
	// act
	lab := NewLabyrinth(maxLoc)
	// assert
	if lab.getNode(locOutOfRange) != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNeighborsNil(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetNeighbors(NewLocation(0, 0, 1))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNeighborsNil1(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetNeighbors(NewLocation(0, 1, 0))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNeighborsNil2(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetNeighbors(NewLocation(1, 0, 0))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetNeighborsSingularLab(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := make([]Node, 0)
	// act
	have := lab.GetNeighbors(maxLoc)
	// assert
	if len(have) != 0 || len(have) != len(want) {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsDualLab0(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(1)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{NewLocation(maxX, maxY, maxZ)}
	// act
	have := lab.GetNeighbors(NewLocation(0, 0, 0))
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !want[i].Compare(have[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsDualLab1(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(1)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(maxX, maxY, maxZ),
	}
	// act
	have := lab.GetNeighbors(NewLocation(0, 0, 0))
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsDualLab2(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(maxX, maxY, maxZ),
	}
	// act
	have := lab.GetNeighbors(NewLocation(0, 0, 0))
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsDualLab00(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(1)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(0, 0, 0),
	}
	// act
	have := lab.GetNeighbors(maxLoc)
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsDualLab11(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(1)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(0, 0, 0),
	}
	// act
	have := lab.GetNeighbors(maxLoc)
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsDualLab22(t *testing.T) {
	// arrange
	maxX := uint(1)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(0, 0, 0),
	}
	// act
	have := lab.GetNeighbors(maxLoc)
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsFullLabyrinth(t *testing.T) {
	// arrange
	maxX := uint(2)
	maxY := uint(2)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(1, 0, 0),
		NewLocation(0, 1, 0),
		NewLocation(0, 0, 1),
	}
	// act
	have := lab.GetNeighbors(NewLocation(0, 0, 0))
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsFullLabyrinth1(t *testing.T) {
	// arrange
	maxX := uint(2)
	maxY := uint(2)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(1, 2, 2),
		NewLocation(2, 1, 2),
		NewLocation(2, 2, 1),
	}
	// act
	have := lab.GetNeighbors(NewLocation(2, 2, 2))
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_GetNeighborsFullLabyrinth2(t *testing.T) {
	// arrange
	maxX := uint(2)
	maxY := uint(2)
	maxZ := uint(2)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	want := []Location{
		NewLocation(0, 1, 1),
		NewLocation(2, 1, 1),
		NewLocation(1, 0, 1),
		NewLocation(1, 2, 1),
		NewLocation(1, 1, 0),
		NewLocation(1, 1, 2),
	}
	// act
	have := lab.GetNeighbors(NewLocation(1, 1, 1))
	// assert
	if len(have) == 1 || len(have) == len(want) {
		for i := 0; i < len(have); i++ {
			if !have[i].Compare(want[i]) {
				t.Errorf("%v should be equal to %v", want, have)
			}
		}
	} else {
		t.Errorf("%v should be equal to %v", want, have)
	}
}

func TestGraphLabyrinth_CompareToNil(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.Compare(nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToItself(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.Compare(lab)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToItself2(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	// act
	have := lab.Compare(lab)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherTrue(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	lab2 := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.Compare(lab2)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2True(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	lab2 := NewLabyrinth(NewLocation(64, 64, 64))
	// act
	have := lab.Compare(lab2)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherFalse(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	lab2 := NewLabyrinth(NewLocation(0, 0, 1))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherFalse2(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	lab2 := NewLabyrinth(NewLocation(0, 1, 0))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherFalse3(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	lab2 := NewLabyrinth(NewLocation(1, 0, 0))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2False(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	lab2 := NewLabyrinth(NewLocation(64, 64, 63))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2False1(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	lab2 := NewLabyrinth(NewLocation(64, 63, 64))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2False2(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	lab2 := NewLabyrinth(NewLocation(63, 64, 64))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherFalse0(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 1))
	lab2 := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherFalse21(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 1, 0))
	lab2 := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOtherFalse32(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(1, 0, 0))
	lab2 := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2False3(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 63))
	lab2 := NewLabyrinth(NewLocation(64, 64, 64))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2False14(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 63, 64))
	lab2 := NewLabyrinth(NewLocation(64, 64, 64))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CompareToOther2False25(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(63, 64, 64))
	lab2 := NewLabyrinth(NewLocation(64, 64, 64))
	// act
	have := lab.Compare(lab2)
	// assert
	if have {
		t.Errorf("")
	}
}
