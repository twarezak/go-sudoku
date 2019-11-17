package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sudoku [9][9]int

func main() {
	handle, err := os.Open("sudoku.txt")

	if err != nil {
		log.Fatalf("cannot read file: %s", err)
	}
	defer handle.Close()
	sudoku := readSudoku(handle)

	sudokus := solve(steps{&sudokuStep{1, sudoku}})
	sudoku = sudokus[len(sudokus)-1].Sudoku
	fmt.Println(sudoku, Check(sudoku))
}

type sudokuStep struct {
	n int
	Sudoku
}
type steps []*sudokuStep

func solve(steps steps) steps {

	step := steps[len(steps)-1]
	sudoku := step.Sudoku
	n := step.n

	for i, line := range sudoku {
		for j, row := range line {
			if row == 0 {
				sudoku[i][j] = n

				if IsThereMistake(sudoku) {
					if n+1 > 9 {
						for k := len(steps) - 1; k > 0; k-- {
							steps = steps[:len(steps)-1]
							lastStep := steps[len(steps)-1]
							if lastStep.n+1 <= 9 {
								steps[len(steps)-1] = &sudokuStep{lastStep.n + 1, lastStep.Sudoku}
								break
							}
						}
					} else {
						sudoku[i][j] = 0
						steps[len(steps)-1] = &sudokuStep{n + 1, sudoku}
					}
				} else {
					steps = append(steps, &sudokuStep{1, sudoku})
				}
				return solve(steps)
			}
		}
	}

	return steps
}

func readSudoku(handle io.Reader) Sudoku {
	var sudoku Sudoku

	scanner := bufio.NewScanner(handle)

	i := 0
	for scanner.Scan() {
		l := scanner.Text()
		sudoku[i] = stringLineToInt(strings.Split(l, " "))
		i++
	}

	return sudoku
}

func stringLineToInt(split []string) [9]int {
	ary := [9]int{}

	for i := range ary {
		ary[i], _ = strconv.Atoi(split[i])
	}

	return ary
}
