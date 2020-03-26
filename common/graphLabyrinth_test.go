package common

import (
	"testing"
)

func TestGraphLabyrinth_CtorNil(t *testing.T) {
	// arrange
	var want Labyrinth = nil
	// act
	have := NewLabyrinth(nil)
	// assert
	if want != have {
		t.Errorf("")
	}
}

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

func TestGraphLabyrinth_GetNeighborsNil0(t *testing.T) {
	// arrange
	maxX := uint(0)
	maxY := uint(0)
	maxZ := uint(0)
	maxLoc := NewLocation(maxX, maxY, maxZ)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetNeighbors(nil)
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

func TestGraphLabyrinth_CheckLocation(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	loc := NewLocation(0, 0, 0)
	// act
	have := lab.checkLocation(loc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CheckLocation1(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	// act
	have := lab.checkLocation(nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CheckLocation2(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	loc := NewLocation(0, 0, 1)
	// act
	have := lab.checkLocation(loc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CheckLocation3(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	loc := NewLocation(0, 1, 0)
	// act
	have := lab.checkLocation(loc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CheckLocation4(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(0, 0, 0))
	loc := NewLocation(1, 0, 0)
	// act
	have := lab.checkLocation(loc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CheckLocation5(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	loc := NewLocation(0, 0, 0)
	// act
	have := lab.checkLocation(loc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_CheckLocation6(t *testing.T) {
	// arrange
	lab := NewLabyrinth(NewLocation(64, 64, 64))
	loc := NewLocation(6, 56, 13)
	// act
	have := lab.checkLocation(loc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestStatic_GetIndex(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	x := uint(0)
	y := uint(0)
	z := uint(0)
	want := uint(0)
	// act
	have := getIndex(x, y, z, maxLoc)
	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestStatic_GetIndex1(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	x := uint(15)
	y := uint(3)
	z := uint(2)
	want := uint(20)
	// act
	have := getIndex(x, y, z, maxLoc)
	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestStatic_GetIndex2(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 9, 2)
	x := uint(15)
	y := uint(3)
	z := uint(2)
	want := uint(38)
	// act
	have := getIndex(x, y, z, maxLoc)
	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestStatic_GetIndex3(t *testing.T) {
	// arrange
	maxLoc := NewLocation(9, 9, 9)
	x := uint(15)
	y := uint(3)
	z := uint(2)
	want := uint(245)
	// act
	have := getIndex(x, y, z, maxLoc)
	// assert
	if want != have {
		t.Errorf("%v should equal to %v", want, have)
	}
}

func TestStatic_ReplaceNodes(t *testing.T) {
	// arrange
	locOfNode := NewLocation(2, 2, 2)
	nodeToReplace := newNode(locOfNode)
	maxLoc := NewLocation(2, 2, 2)
	have := []Node{
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	want := []Node{
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nodeToReplace,
	}
	// act
	have = replaceNodes(nodeToReplace, have, maxLoc)
	// assert
	for i, node := range want {
		if (node != nil || have[i] != nil) && (node == nil || have[i] == nil) {
			t.Errorf("")
		} else {
			if node != nil {
				if !node.compare(have[i]) {
					t.Errorf("")
				}
			}
		}
	}
}

func TestStatic_ReplaceNodes1(t *testing.T) {
	// arrange
	locOfNode := NewLocation(0, 2, 0)
	nodeToReplace := newNode(locOfNode)
	maxLoc := NewLocation(2, 2, 2)
	have := []Node{
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	want := []Node{
		nil, nil, nil, nil, nil, nil, nodeToReplace, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	// act
	have = replaceNodes(nodeToReplace, have, maxLoc)
	// assert
	for i, node := range want {
		if (node != nil || have[i] != nil) && (node == nil || have[i] == nil) {
			t.Errorf("%v should equal to %v", node, have[i])
		} else {
			if node != nil {
				if !node.compare(have[i]) {
					t.Errorf("")
				}
			}
		}
	}
}

func TestStatic_ReplaceNodes2(t *testing.T) {
	// arrange
	locOfNode := NewLocation(2, 0, 0)
	nodeToReplace := newNode(locOfNode)
	maxLoc := NewLocation(2, 2, 2)
	have := []Node{
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	want := []Node{
		nil, nil, nodeToReplace, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	// act
	have = replaceNodes(nodeToReplace, have, maxLoc)
	// assert
	for i, node := range want {
		if (node != nil || have[i] != nil) && (node == nil || have[i] == nil) {
			t.Errorf("")
		} else {
			if node != nil {
				if !node.compare(have[i]) {
					t.Errorf("")
				}
			}
		}
	}
}

func TestStatic_ReplaceNodes3(t *testing.T) {
	// arrange
	locOfNode := NewLocation(0, 0, 0)
	nodeToReplace := newNode(locOfNode)
	maxLoc := NewLocation(2, 2, 2)
	have := []Node{
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	want := []Node{
		nodeToReplace, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
	}
	// act
	have = replaceNodes(nodeToReplace, have, maxLoc)
	// assert
	for i, node := range want {
		if (node != nil || have[i] != nil) && (node == nil || have[i] == nil) {
			t.Errorf("")
		} else {
			if node != nil {
				if !node.compare(have[i]) {
					t.Errorf("")
				}
			}
		}
	}
}

func TestGraphLabyrinth_Connect(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(nil, maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect1(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect2(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(0, 0, 1))
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect3(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(0, 1, 0))
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect4(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(1, 0, 0))
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect5(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 1)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(0, 0, 0))
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect6(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 1, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(0, 0, 0))
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect7(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(NewLocation(0, 0, 1), maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect8(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(NewLocation(0, 1, 0), maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect9(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(NewLocation(1, 0, 0), maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect10(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 1)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(NewLocation(0, 0, 0), maxLoc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect11(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 1, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(NewLocation(0, 0, 0), maxLoc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect12(t *testing.T) {
	// arrange
	maxLoc := NewLocation(1, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(0, 0, 0))
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Connect13(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 0, 0)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Connect(maxLoc, NewLocation(0, 0, 0))
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Disconnect(nil, maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect1(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.Disconnect(maxLoc, nil)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect2(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(3, 2, 2)
	// act
	have := lab.Disconnect(maxLoc, loc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect3(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(3, 2, 2)
	// act
	have := lab.Disconnect(loc, maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect4(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(0, 2, 2)
	// act
	have := lab.Disconnect(loc, maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect5(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(1, 2, 2)
	// act
	have := lab.Disconnect(loc, maxLoc)
	// assert
	if have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect6(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(1, 2, 2)
	lab.Connect(maxLoc, loc)
	// act
	have := lab.Disconnect(loc, maxLoc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_Disconnect7(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(1, 2, 2)
	lab.Connect(loc, maxLoc)
	// act
	have := lab.Disconnect(loc, maxLoc)
	// assert
	if !have {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetConnected(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetConnected(nil)
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetConnected1(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetConnected(NewLocation(3, 2, 2))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetConnected2(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetConnected(NewLocation(2, 3, 2))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetConnected3(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetConnected(NewLocation(2, 2, 3))
	// assert
	if have != nil {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetConnected4(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	// act
	have := lab.GetConnected(NewLocation(1, 1, 1))
	// assert
	if have == nil || len(have) != 0 {
		t.Errorf("")
	}
}

func TestGraphLabyrinth_GetConnected5(t *testing.T) {
	// arrange
	maxLoc := NewLocation(2, 2, 2)
	lab := NewLabyrinth(maxLoc)
	loc := NewLocation(1, 1, 1)
	con1 := NewLocation(1, 1, 0)
	con2 := NewLocation(1, 2, 1)
	lab.Connect(loc, con1)
	lab.Connect(loc, con2)
	want := []Location{con1, con2}
	// act
	have := lab.GetConnected(loc)
	// assert
	if have == nil || len(have) != 2 {
		t.Errorf("")
	} else {
		for i, l := range want {
			if !l.Compare(have[i]) {
				t.Errorf("")
			}
		}
	}
}
