package main

// import (
// 	"fmt"
// 	"strconv"
// )

// /*
// SUDOKU RULES
// 1. 9 by 9 square
// 2. each row and column has 1-9
// 4. each 3x3 square has 1-9
// 5. no repeats in rows, columns, and 3x3 squares
// */

// // Outputs the current board to screen
// func displayBoard(puzz [][]int) {
// 	puzz_cols := len(puzz[0])
// 	fmt.Println("-------------------------")
// 	for i := 0; i < len(puzz); i++ {
// 		fmt.Print("| ")
// 		for j := 0; j < puzz_cols; j++ {
// 			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
// 			if (j+1)%3 == 0 {
// 				fmt.Print("| ")
// 			}
// 		}
// 		fmt.Println()
// 		if (i+1)%3 == 0 {
// 			fmt.Println("-------------------------")
// 		}
// 	}
// }

// // Return a valid row and column for a 0 on the grid
// func getEmptySpace(puzz [][]int) (int, int) {
// 	puzz_cols := len(puzz[0])
// 	for i := 0; i < len(puzz); i++ {
// 		for j := 0; j < puzz_cols; j++ {
// 			if puzz[i][j] == 0 {
// 				return i, j
// 			}
// 		}
// 	}
// 	return -1, -1 // No empty spaces on the grid
// }

// // check values on row
// func isNumValid(puzz [][]int, guess int, row int, column int) bool {
// 	for index := range puzz {
// 		if puzz[row][index] == guess && column != index {
// 			return false
// 		}
// 	}

// 	// check values on column
// 	for index := range puzz {
// 		if puzz[index][column] == guess && column != index {
// 			return false
// 		}
// 	}

// 	//Is number valid for box?
// 	//Row 0 & column 3
// 	//Row - (Row % 3) + Value for cycling(0-2)
// 	// 0 - (0 % 3) + 0 = 0 - 0 + 0 = 0 (1st row in box)
// 	// 0 - (0 % 3) + 1 = 0 - 0 + 1 = 1 (1st row in box)
// 	// 0 - (0 % 3) + 2 = 0 - 0 + 2 = 2 (1st row in box)

// 	//Col - (Col % 3) + Value for cycling(0-2)
// 	// 3 - (3 % 3) + 0 = 3 - 0 + 0 = 3 (1st col in box)
// 	// 3 - (3 % 3) + 1 = 3 - 0 + 1 = 4 (1st col in box)
// 	// 3 - (3 % 3) + 2 = 3 - 0 + 2 = 5 (1st col in box)

// 	for k := 0; k < 3; k++ {
// 		for l := 0; l < 3; l++ {
// 			if puzz[row-row%3+k][column-column%3+l] ==
// 				guess && (row-row%3+k != row || column-column%3+l != column) {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// /*
// -------------------------
// | 0 0 0 | 0 3 5 | 0 7 0 |
// | 2 5 0 | 0 4 6 | 8 0 1 |
// | 0 1 3 | 7 0 8 | 0 4 9 |
// -------------------------
// | 1 9 0 | 0 0 7 | 0 0 4 |
// | 0 0 5 | 0 0 2 | 0 9 6 |
// | 8 0 2 | 0 9 4 | 0 0 7 |
// -------------------------
// | 3 7 0 | 0 0 9 | 0 0 0 |
// | 0 6 1 | 0 7 0 | 0 0 0 |
// | 4 0 0 | 5 8 1 | 0 0 0 |
// -------------------------

// ### actual solution
// -------------------------
// | 9 8 4 | 1 3 5 | 6 7 2 |
// | 2 5 7 | 9 4 6 | 8 3 1 |
// | 6 1 3 | 7 2 8 | 5 4 9 |
// -------------------------
// | 1 9 6 | 3 5 7 | 2 8 4 |
// | 7 4 5 | 8 1 2 | 3 9 6 |
// | 8 3 2 | 6 9 4 | 1 5 7 |
// -------------------------
// | 3 7 8 | 2 6 9 | 4 1 5 |
// | 5 6 1 | 4 7 3 | 9 2 8 |
// | 4 2 9 | 5 8 1 | 7 6 3 |
// -------------------------

// // Recursion Problem (Solution)
// // 1. Cycle across the rows column by column (1-9)
// // 2. Check if valid number
// // a. If true update array
// // b. If false change back to 0
// // c. If false find a new value for previous we
// // thought was correct
// // 3. Check next column
// */

// func solvePuzzle(puzz [][]int) bool {
// 	row, column := getEmptySpace(puzz)
// 	if row == -1 {
// 		return true
// 	} else {
// 		row, column = getEmptySpace(puzz)
// 	}
// 	for i := 1; i <= 9; i++ {
// 		if isNumValid(puzz, i, row, column) {
// 			puzz[row][column] = i

// 			if solvePuzzle(puzz) {
// 				return true
// 			}
// 			puzz[row][column] = 0
// 		}
// 	}
// 	return false
// }

// func main() {
// 	puzz := [][]int{
// 		{0, 0, 0, 0, 3, 5, 0, 7, 0},
// 		{2, 5, 0, 0, 4, 6, 8, 0, 1},
// 		{0, 1, 3, 7, 0, 8, 0, 4, 9},
// 		{1, 9, 0, 0, 0, 7, 0, 0, 4},
// 		{0, 0, 5, 0, 0, 2, 0, 9, 6},
// 		{8, 0, 2, 0, 9, 4, 0, 0, 7},
// 		{3, 7, 0, 0, 0, 9, 0, 0, 0},
// 		{0, 6, 1, 0, 7, 0, 0, 0, 0},
// 		{4, 0, 0, 5, 8, 1, 0, 0, 0},
// 	}
// 	displayBoard(puzz)

// 	// // Return first open space
// 	// row, _ := getEmptySpace(puzz)
// 	// if row != -1 {
// 	// 	fmt.Println(getEmptySpace(puzz))
// 	// } else {
// 	// 	fmt.Println("Solved")
// 	// 	return
// 	// }

// 	solvePuzzle(puzz)
// 	fmt.Println("##########################")
// 	displayBoard(puzz)

// }
