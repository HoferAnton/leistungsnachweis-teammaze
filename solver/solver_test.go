package solver

import (
	"reflect"
	"testing"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

func TestContains_NotIncluded(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := common.NewLocation(7, 7, 7)
	want := false

	// act
	have := contains(list, element)

	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestContains_first(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := list[0]
	want := true

	// act
	have := contains(list, element)

	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestContains_Last(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := list[len(list)-1]
	want := true

	// act
	have := contains(list, element)

	// assert
	if want != have {
		t.Errorf("")
	}
}

func TestContains_middle(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := list[len(list)/2]
	want := true

	// act
	have := contains(list, element)

	// assert
	if want != have {
		t.Errorf("")
	}
}
func TestRemoveFirstOccurrence_NotIncluded(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := common.NewLocation(7, 7, 7)
	want := list

	// act
	have := removeFirstOccurrence(list, element)

	// assert
	if !reflect.DeepEqual(want, have) {
		t.Errorf("")
	}
}

func TestRemoveFirstOccurrence_first(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := list[0]
	want := []common.Location{
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}

	// act
	have := removeFirstOccurrence(list, element)

	// assert
	if !reflect.DeepEqual(want, have) {
		t.Errorf("")
	}
}

func TestRemoveFirstOccurrence_Last(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := list[len(list)-1]
	want := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
	}

	// act
	have := removeFirstOccurrence(list, element)

	// assert
	if !reflect.DeepEqual(want, have) {
		t.Errorf("")
	}
}

func TestRemoveFirstOccurrence_middle(t *testing.T) {
	// arrange
	list := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(2, 1, 3),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	element := list[2]
	want := []common.Location{
		common.NewLocation(0, 0, 0),
		common.NewLocation(1, 3, 2),
		common.NewLocation(3, 2, 1),
		common.NewLocation(5, 5, 5),
	}
	// act
	have := removeFirstOccurrence(list, element)

	// assert
	if !reflect.DeepEqual(want, have) {
		t.Errorf("")
	}
}
