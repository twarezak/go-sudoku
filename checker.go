package main

func Check(sudoku Sudoku) bool {
	flatSudoku1st, flatSudoku2nd, flatSudoku3rd := makeFlats(sudoku)

	for i := 0; i < 9; i++ {
		if CheckSlice(flatSudoku1st[i*9:(i+1)*9]) == false {
			return false
		}
		if CheckSlice(flatSudoku2nd[i*9:(i+1)*9]) == false {
			return false
		}
		if CheckSlice(flatSudoku3rd[i*9:(i+1)*9]) == false {
			return false
		}
	}

	return true
}

func makeFlats(sudoku Sudoku) ([81]int, [81]int, [81]int) {
	var flatSudoku1st, flatSudoku2nd, flatSudoku3rd [81]int
	var i, j, n, m, k int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			flatSudoku1st[n] = sudoku[i][j]
			flatSudoku2nd[n] = sudoku[j][i]
			n++
		}
	}

	n = 0

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			for m = 0; m < 3; m++ {
				for k = 0; k < 3; k++ {
					flatSudoku3rd[n] = sudoku[3*i+m][3*j+k]
					n++
				}
			}
		}
	}
	return flatSudoku1st, flatSudoku2nd, flatSudoku3rd
}

func CheckSlice(sudokuSlice []int) bool {
	keys := make(map[int]bool)

	for _, entry := range sudokuSlice {
		if entry < 0 || entry > 9 {
			return false
		}
		if _, value := keys[entry]; value {
			return false
		}
		keys[entry] = true
	}
	return true
}

func IsThereMistake(sudoku Sudoku) bool {
	flatSudoku1st, flatSudoku2nd, flatSudoku3rd := makeFlats(sudoku)

	for i := 0; i < 9; i++ {
		if isMistakeInSlice(flatSudoku1st[i*9 : (i+1)*9]) {
			return true
		}
		if isMistakeInSlice(flatSudoku2nd[i*9 : (i+1)*9]) {
			return true
		}
		if isMistakeInSlice(flatSudoku3rd[i*9 : (i+1)*9]) {
			return true
		}
	}

	return false
}

func isMistakeInSlice(numbers []int) bool {
	dict := make(map[int]int)

	for _, i := range numbers {
		if i > 9 {
			return false
		}
		dict[i]++
		if i != 0 && dict[i] > 1 {
			return true
		}
	}

	return false
}