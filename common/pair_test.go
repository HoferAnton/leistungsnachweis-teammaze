package common

import "testing"

func TestCtor_NewPair(t *testing.T) {
	sut := NewPair("asdf", "qwer")
	var nilPair Pair

	if sut == nilPair {
		t.Errorf("")
	}
}

func TestCtor3_NewPair(t *testing.T) {
	sut := NewPair(nil, nil)
	var nilPair Pair

	if sut != nilPair {
		t.Errorf("")
	}
}

func TestCtor_GetFirstAndGetSecond(t *testing.T) {
	sut := NewPair("asdf", "qwer")

	if sut.GetFirst() != "asdf" && sut.GetSecond() != "qwer" {
		t.Errorf("")
	}
}

func TestCtor2_GetFirstAndGetSecond(t *testing.T) {
	sut := NewPair(nil, nil)

	if sut.GetFirst() != nil && sut.GetSecond() != nil {
		t.Errorf("")
	}
}

func TestPair0_Compare(t *testing.T) {
	sut := NewPair(nil, nil)

	if !sut.Compare(sut) {
		t.Errorf("")
	}
}

func TestPairCompareToItself_Compare(t *testing.T) {
	sut := NewPair("nil", "nil")

	if !sut.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair1_Compare(t *testing.T) {
	sut := NewPair(nil, nil)
	var nilPair Pair

	if !sut.Compare(nilPair) {
		t.Errorf("")
	}
}

func TestPair2_Compare(t *testing.T) {
	sut := NewPair(nil, nil)
	sut2 := NewPair(nil, nil)

	if !sut.Compare(sut2) && !sut2.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair3_Compare(t *testing.T) {
	sut := NewPair(nil, nil)
	sut2 := NewPair(nil, "nil")

	if sut.Compare(sut2) && sut2.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair4_Compare(t *testing.T) {
	sut := NewPair(nil, nil)
	sut2 := NewPair("nil", nil)

	if sut.Compare(sut2) && sut2.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair5_Compare(t *testing.T) {
	sut := NewPair(nil, nil)
	sut2 := NewPair("nil", "nil")

	if sut.Compare(sut2) && sut2.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair6_Compare(t *testing.T) {
	sut := NewPair("nil", "nil")
	sut2 := NewPair("nil", "nil")

	if !sut.Compare(sut2) && !sut2.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair7_Compare(t *testing.T) {
	sut := NewPair("nil", "nil")
	sut2 := NewPair("nil", "nil1")

	if sut.Compare(sut2) && sut2.Compare(sut) {
		t.Errorf("")
	}
}

func TestPair8_Compare(t *testing.T) {
	sut := NewPair("nil", "nil")
	sut2 := NewPair("nil1", "nil")

	if sut.Compare(sut2) && sut2.Compare(sut) {
		t.Errorf("")
	}
}
