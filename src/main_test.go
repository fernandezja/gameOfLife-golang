package main

import "testing"

func TestDummy(t *testing.T) {
	t.Log("test dummy")
}

func TestCreateNeighborhood3x3(t *testing.T) {
	expectedCeld00 := 0
	expectedCeld22 := 0
	got := CreateNeighborhood(3, 3)

	if expectedCeld00 != got[0][0] {

		t.Errorf("Cell [0][0] got '%d' expected '%d'", got[0][0], expectedCeld00)
	}

	if expectedCeld22 != got[2][2] {
		t.Errorf("Cell [2][2] got '%d' expected '%d'", got[2][2], expectedCeld22)
	}
}

func TestCreateNeighborhood9x9(t *testing.T) {
	expectedCeld00 := 0
	expectedCeld88 := 0
	got := CreateNeighborhood(9, 9)

	if expectedCeld00 != got[0][0] {

		t.Errorf("Cell [0][0] got '%d' expected '%d'", got[0][0], expectedCeld00)
	}

	if expectedCeld88 != got[8][8] {
		t.Errorf("Cell [8][8] got '%d' expected '%d'", got[2][2], expectedCeld88)
	}
}
