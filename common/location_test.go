package common

import (
	"testing"
)

func TestLocationCtorMin(t *testing.T) {
	// arrange
	var wantX uint = 0
	var wantY uint = 0
	var wantZ uint = 0
	// act
	location := NewLocation(0, 0, 0)
	haveX, haveY, haveZ := location.As3DCoordinates()
	// assert
	if wantX != haveX || wantY != haveY || wantZ != haveZ {
		t.Errorf("")
	}
}

func TestLocationCtorNormal(t *testing.T) {
	// arrange
	var wantX uint = 1000
	var wantY uint = 1010
	var wantZ uint = 1200
	// act
	location := NewLocation(1000, 1010, 1200)
	haveX, haveY, haveZ := location.As3DCoordinates()
	// assert
	if wantX != haveX || wantY != haveY || wantZ != haveZ {
		t.Errorf("")
	}
}

func TestLocationCtorMax(t *testing.T) {
	// arrange
	wantX := ^uint(0)
	wantY := ^uint(0) - 1
	wantZ := ^uint(0) - 2
	// act
	location := NewLocation(^uint(0), ^uint(0)-1, ^uint(0)-2)
	haveX, haveY, haveZ := location.As3DCoordinates()
	// assert
	if wantX != haveX || wantY != haveY || wantZ != haveZ {
		t.Errorf("")
	}
}

func TestLocationCompareMinTrue(t *testing.T) {
	// arrange
	location1 := NewLocation(0, 0, 0)
	location2 := NewLocation(0, 0, 0)
	// act
	isEqual := location1.Compare(location2)
	// assert
	if !isEqual {
		t.Errorf("")
	}
}

func TestLocationCompareMinTrue2(t *testing.T) {
	// arrange
	location1 := NewLocation(^uint(0), ^uint(0)-15, ^uint(0)-156)
	location2 := NewLocation(^uint(0), ^uint(0)-15, ^uint(0)-156)
	// act
	isEqual := location1.Compare(location2)
	// assert
	if !isEqual {
		t.Errorf("")
	}
}

func TestLocationCompareMinFalse(t *testing.T) {
	// arrange
	location1 := NewLocation(0, 0, 0)
	location2 := NewLocation(0, 0, 1)
	// act
	isEqual := location1.Compare(location2)
	// assert
	if isEqual {
		t.Errorf("")
	}
}

func TestLocationCompareMinFalse1(t *testing.T) {
	// arrange
	location1 := NewLocation(0, 0, 0)
	location2 := NewLocation(0, 1, 0)
	// act
	isEqual := location1.Compare(location2)
	// assert
	if isEqual {
		t.Errorf("")
	}
}

func TestLocationCompareMinFalse2(t *testing.T) {
	// arrange
	location1 := NewLocation(0, 0, 0)
	location2 := NewLocation(1, 0, 0)
	// act
	isEqual := location1.Compare(location2)
	// assert
	if isEqual {
		t.Errorf("")
	}
}

func TestLocationCompareMinFalse3(t *testing.T) {
	// arrange
	location1 := NewLocation(0, 0, 0)
	var location2 Location = nil
	// act
	isEqual := location1.Compare(location2)
	// assert
	if isEqual {
		t.Errorf("")
	}
}
