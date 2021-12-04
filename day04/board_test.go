package main

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	b, err := NewBoard([]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19})
	if err != nil {
		t.Error(err)
	}

	testField := Field{13, false, 1, 0}

	if b.fields[1] != testField {
		t.Errorf("Field 1 not set correctly, got: %+v, want: %+v", b.fields[1], testField)
	}

	if b.byNumber[13] != &b.fields[1] {
		t.Errorf("ByNumber mapping not correct, got: %+v, want: %+v", b.byNumber[13], &b.fields[1])
	}
}

func TestFilter(t *testing.T) {
	b, err := NewBoard([]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19})
	if err != nil {
		t.Error(err)
	}

	if len(b.Filter(false)) != 25 {
		t.Errorf("Filter(true) should return 25 fields, got: %d", len(b.Filter(true)))
	}

	b.byNumber[22].checked = true
	b.byNumber[13].checked = true

	if len(b.Filter(false)) != 23 {
		t.Errorf("Filter(true) should return 23 fields, got: %d", len(b.Filter(true)))
	}

	if b.Filter(true)[0].number != 22 && b.Filter(true)[1].number != 13 {
		t.Errorf("Filter(true) should return 22 and 13, got: %d & %d", b.Filter(true)[0].number, b.Filter(true)[1].number)
	}
}

func TestHasWon(t *testing.T) {
	b, err := NewBoard([]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19})
	if err != nil {
		t.Error(err)
	}

	b.byNumber[22].checked = true
	b.byNumber[2].checked = true
	b.byNumber[8].checked = true

	if b.HasWon() {
		t.Error("Board 1 should not be won")
	}

	b.byNumber[13].checked = true
	b.byNumber[17].checked = true
	b.byNumber[11].checked = true
	b.byNumber[0].checked = true

	if !b.HasWon() {
		t.Error("Board 1 should be won")
	}

	b, err = NewBoard([]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19})
	if err != nil {
		t.Error(err)
	}

	b.byNumber[6].checked = true
	b.byNumber[10].checked = true
	b.byNumber[3].checked = true
	b.byNumber[18].checked = true
	b.byNumber[5].checked = true

	if !b.HasWon() {
		t.Error("Board 2 should be won")
	}
}

func TestGetScore(t *testing.T) {
	b, err := NewBoard([]int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19})
	if err != nil {
		t.Error(err)
	}

	b.byNumber[6].checked = true
	b.byNumber[10].checked = true
	b.byNumber[3].checked = true
	b.byNumber[18].checked = true
	b.byNumber[5].checked = true

	if b.GetScore(5) != 1290 {
		t.Errorf("Board 1 score should be 1290, got: %d", b.GetScore(5))
	}
}
