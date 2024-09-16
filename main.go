package main

import (
	"fmt"
)

const size = 9

func isValid(board [size][size]int, row, col, num int) bool {
	for x := 0; x < size; x++ {
		if board[row][x] == num || board[x][col] == num || board[row-(row%3)+x/3][col-(col%3)+x%3] == num {
			return false
		}
	}
	return true
}

func solveSudoku(board [size][size]int) bool {
	row, col, empty := -1, -1, true
	for i := 0; i < size && empty; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				row, col, empty = i, j, false
				break
			}
		}
	}

	if empty {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

func printBoard(board [size][size]int) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func main() {
	board := [size][size]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if solveSudoku(board) {
		fmt.Println("Solved Sudoku:")
		printBoard(board)
	} else {
		fmt.Println("No solution exists.")
	}
}