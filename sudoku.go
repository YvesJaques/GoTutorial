package main

import (
	"fmt"
	"strconv"
)

/*
SUDOKU RULES
1. 9 by 9 square
2. each row and column has 1-9
4. each 3x3 square has 1-9
5. no repeats in rows, columns, and 3x3 squares
*/

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	for i := 0; i < len(puzz); i++ {
		for j := 0; j < puzz_cols; j++ {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Println()
	}
}

// Return a valid row and column for a 0 on the grid
func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	for i := 0; i < len(puzz); i++ {
		for j := 0; j < puzz_cols; j++ {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // No empty spaces on the grid
}

func isNumValid(puzz [][]int, guess int, row int, column int) bool {
	for index := range puzz {
		if puzz[row][index] == guess && column != index {
			return false
		}
	}
	return true
}

func main() {
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}
	displayBoard(puzz)

	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Solved")
		return
	}

	fmt.Println(isNumValid(puzz, 7, 0, 0))
}