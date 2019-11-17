package main

import (
	"strings"
	"testing"
)

func TestReadSudoku(t *testing.T) {

	s := `8 9 5 7 4 2 1 3 6
2 7 1 9 6 3 4 8 5
4 6 3 5 8 1 7 9 2
9 3 4 6 1 7 2 5 8
5 1 7 2 3 8 9 6 4
6 8 2 4 5 9 3 7 1
1 5 9 8 7 4 6 2 3
7 4 6 3 2 5 8 1 9
3 2 8 1 9 6 5 4 7`

	want := Sudoku{
		{8, 9, 5, 7, 4, 2, 1, 3, 6},
		{2, 7, 1, 9, 6, 3, 4, 8, 5},
		{4, 6, 3, 5, 8, 1, 7, 9, 2},
		{9, 3, 4, 6, 1, 7, 2, 5, 8},
		{5, 1, 7, 2, 3, 8, 9, 6, 4},
		{6, 8, 2, 4, 5, 9, 3, 7, 1},
		{1, 5, 9, 8, 7, 4, 6, 2, 3},
		{7, 4, 6, 3, 2, 5, 8, 1, 9},
		{3, 2, 8, 1, 9, 6, 5, 4, 7},
	}

	got := readSudoku(strings.NewReader(s))
	if want != got {
		t.Errorf("Incorrect read sudoku.txt. want: %v, got: %v", want, got)
	}
}
