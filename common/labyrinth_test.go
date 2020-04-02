package common

import (
	"math/rand"
	"testing"
	"time"
)

func TestStatic_GetIndex(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	x := uint(0)
	y := uint(0)
	z := uint(0)
	want := uint(0)
	// act
	have := GetIndex(x, y, z, maxLoc)
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
	have := GetIndex(x, y, z, maxLoc)
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
	have := GetIndex(x, y, z, maxLoc)
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
	have := GetIndex(x, y, z, maxLoc)
	// assert
	if want != have {
		t.Errorf("%v should equal to %v", want, have)
	}
}

func TestStatic_GetLocation(t *testing.T) {
	// arrange
	maxLoc := NewLocation(0, 0, 0)
	index := uint(1)

	var wantLoc Location
	// act
	haveLoc := GetLocation(index, maxLoc)
	// assert
	if wantLoc != haveLoc {
		t.Errorf("")
	}
}

func TestStatic_GetLocation0(t *testing.T) {
	// arrange
	maxLoc := NewLocation(1, 2, 3)
	index := uint(25)

	var wantLoc Location
	// act
	haveLoc := GetLocation(index, maxLoc)
	// assert
	if wantLoc != haveLoc {
		t.Errorf("")
	}
}

func TestStatic_GetLocation1(t *testing.T) {
	// arrange
	maxLoc := NewLocation(15, 7, 26)
	x := uint(12)
	y := uint(5)
	z := uint(23)
	index := uint(3036)
	wantLoc := NewLocation(x, y, z)
	// act
	haveLoc := GetLocation(index, maxLoc)
	// assert
	if !wantLoc.Compare(haveLoc) {
		t.Errorf("")
	}
}

func TestStatic_GetLocationIsInverseToGetIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	upto := uint(250)

	for maxZ := uint(0); maxZ <= upto; maxZ++ {
		for maxY := uint(0); maxY <= upto; maxY++ {
			for maxX := uint(0); maxX <= upto; maxX++ {
				x := uint(rand.Intn(int(maxX + 1)))
				y := uint(rand.Intn(int(maxY + 1)))
				z := uint(rand.Intn(int(maxZ + 1)))
				// arrange
				maxLoc := NewLocation(maxX, maxY, maxZ)
				index := GetIndex(x, y, z, maxLoc)
				wantLoc := NewLocation(x, y, z)
				// act
				haveLoc := GetLocation(index, maxLoc)
				// assert
				if !wantLoc.Compare(haveLoc) {
					t.Errorf("")
				}
			}
		}
	}
}
