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

func TestSetCell1x1ToLife(t *testing.T) {
	expected := 1
	neighborhood := CreateNeighborhood(3, 3)

	SetCell(neighborhood, 0, 0)

	got := neighborhood[0][0]
	if expected != got {

		t.Errorf("Cell [0][0] got '%d' expected '%d'", got, expected)
	}

	got2 := neighborhood[1][1]
	if 0 != got2 {

		t.Errorf("Cell [1][1] got '%d' expected '%d'", got2, 0)
	}

}

func TestSetCellDiagonal(t *testing.T) {
	expected := 1
	neighborhood := CreateNeighborhood(3, 3)

	SetCell(neighborhood, 0, 0)
	SetCell(neighborhood, 1, 1)
	SetCell(neighborhood, 2, 2)

	got0 := neighborhood[0][0]
	got1 := neighborhood[0][0]
	got2 := neighborhood[0][0]

	if expected != got0 {

		t.Errorf("Cell [0][0] got '%d' expected '%d'", got0, expected)
	}

	if expected != got1 {

		t.Errorf("Cell [0][0] got '%d' expected '%d'", got1, expected)
	}

	if expected != got2 {

		t.Errorf("Cell [0][0] got '%d' expected '%d'", got2, expected)
	}

}

func TestNeighborsCountTo1(t *testing.T) {
	expected := 1
	neighborhood := CreateNeighborhood(3, 3)

	SetCell(neighborhood, 0, 0)

	got := NeighborsCount(neighborhood, 1, 1)

	if expected != got {

		t.Errorf("Neighbors Count [1][1] got '%d' expected '%d'", got, expected)
	}

}

func TestNeighborsCountTo0(t *testing.T) {
	expected := 0
	neighborhood := CreateNeighborhood(3, 3)

	SetCell(neighborhood, 0, 0)

	got := NeighborsCount(neighborhood, 2, 2)

	if expected != got {

		t.Errorf("Neighbors Count [2][2] got '%d' expected '%d'", got, expected)
	}

}

func TestNeighborsCountTo2(t *testing.T) {
	expected := 2
	neighborhood := CreateNeighborhood(3, 3)

	SetCell(neighborhood, 0, 0)
	SetCell(neighborhood, 0, 1)

	got := NeighborsCount(neighborhood, 1, 1)

	if expected != got {

		t.Errorf("Neighbors Count [1][1] got '%d' expected '%d'", got, expected)
	}

}

func TestNeighborsCountTo3(t *testing.T) {
	expected := 3
	neighborhood := CreateNeighborhood(3, 3)

	SetCell(neighborhood, 0, 0)
	SetCell(neighborhood, 0, 1)
	SetCell(neighborhood, 0, 2)

	got := NeighborsCount(neighborhood, 1, 1)

	if expected != got {

		t.Errorf("Neighbors Count [1][1] got '%d' expected '%d'", got, expected)
	}

}
