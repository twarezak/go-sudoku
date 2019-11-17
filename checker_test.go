package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	sudoku := [9][9]int{
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


	sudokuIncorrect1st := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 3, 1, 5, 6, 4, 8, 9, 7},
		{3, 1, 2, 6, 4, 5, 9, 7, 8},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 4, 8, 9, 7, 2, 3, 1},
		{6, 4, 5, 9, 7, 8, 3, 1, 2},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 7, 2, 3, 1, 5, 6, 4},
		{9, 7, 8, 3, 1, 2, 6, 4, 5},
	}

	sudokuIncorrect2nd := [9][9]int{
		{8, 9, 5, 7, 4, 2, 1, 3, 6},
		{2, 7, 1, 9, 6, 3, 4, 8, 5},
		{4, 6, 3, 5, 8, 1, 7, 9, 2},
		{9, 3, 4, 6, 1, 7, 2, 5, 8},
		{5, 1, 7, 2, 3, 8, 9, 6, 4},
		{6, 2, 8, 4, 5, 9, 3, 7, 1},
		{1, 5, 9, 8, 7, 4, 6, 2, 3},
		{7, 4, 6, 3, 2, 5, 8, 1, 9},
		{3, 2, 8, 1, 9, 6, 5, 4, 7},
	}

	sudokuIncorrect3rd := [9][9]int{
		{8, 9, 5, 7, 4, 2, 1, 3, 6},
		{2, 7, 1, 9, 6, 3, 4, 8, 5},
		{4, 6, 3, 5, 8, 1, 7, 9, 2},
		{5, 3, 4, 6, 1, 7, 2, 5, 8},
		{9, 1, 7, 2, 3, 8, 9, 6, 4},
		{6, 8, 2, 4, 5, 9, 3, 7, 1},
		{1, 5, 9, 8, 7, 4, 6, 2, 3},
		{7, 4, 6, 3, 2, 5, 8, 1, 9},
		{3, 2, 8, 1, 9, 6, 5, 4, 7},
	}

	got := Check(sudoku)
	want := true

	if got != want {
		t.Errorf("Correct Sudoku: got %t, want %t ", got, want)
	}

	got = Check(sudokuIncorrect1st)
	want = false

	if got != want {
		t.Errorf("Incorrect Sudoku 1st case: got %t, want %t ", got, want)
	}

	got = Check(sudokuIncorrect2nd)
	want = false

	if got != want {
		t.Errorf("Incorrect Sudoku 2nd case: got %t, want %t ", got, want)
	}

	got = Check(sudokuIncorrect3rd)
	want = false

	if got != want {
		t.Errorf("Incorrect Sudoku 3rd case: got %t, want %t ", got, want)
	}
}

func TestCheckSlice(t *testing.T) {
	sliceCorrect := []int{8, 9, 5, 7, 4, 2, 1, 3, 6}
	sliceIncorrect := []int{8, 9, 5, 7, 3, 2, 1, 3, 6}

	got := CheckSlice(sliceCorrect)
	want := true

	if got != want {
		t.Errorf("Correct Sudoku Slice: got %t, want %t ", got, want)
	}

	got = CheckSlice(sliceIncorrect)
	want = false

	if got != want {
		t.Errorf("Incorrect Sudoku Slice: got %t, want %t ", got, want)
	}
}

func TestIsThereMistake(t *testing.T) {
	t.Run("Test correct sudoku 1", func(t *testing.T) {
		correct := [9][9]int{
			{8, 9, 5, 7, 0, 2, 1, 3, 6},
			{2, 7, 1, 9, 0, 3, 4, 8, 5},
			{4, 0, 3, 5, 0, 1, 7, 9, 2},
			{9, 0, 4, 6, 1, 7, 2, 5, 8},
			{5, 1, 7, 2, 3, 8, 9, 6, 4},
			{6, 8, 2, 4, 5, 9, 3, 7, 1},
			{1, 0, 0, 8, 7, 4, 6, 2, 3},
			{7, 4, 0, 3, 0, 5, 8, 1, 9},
			{3, 2, 8, 1, 0, 6, 5, 4, 7},
		}

		want := false
		got := IsThereMistake(correct)
		if got != want  {
			t.Errorf("check `IsThereMistake` faild. got %t, want %t ", got, want)
		}
	})
	t.Run("Test correct sudoku 2", func(t *testing.T) {
		correct := [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}

		want := false
		got := IsThereMistake(correct)
		if got != want  {
			t.Errorf("check `IsThereMistake` faild. got %t, want %t ", got, want)
		}
	})

	t.Run("Test sudoku with mistake", func(t *testing.T) {
		correct := [9][9]int{
			{1, 1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}

		want := true
		got := IsThereMistake(correct)
		if got != want  {
			t.Errorf("check `IsThereMistake` faild. got %t, want %t ", got, want)
		}
	})
	t.Run("Test sudoku with mistake", func(t *testing.T) {
		correct := [9][9]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{4, 5, 6, 1, 2, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}

		want := true
		got := IsThereMistake(correct)
		if got != want  {
			t.Errorf("check `IsThereMistake` faild. got %t, want %t ", got, want)
		}
	})
}