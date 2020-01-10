package main

import (
	"log"
)

func main() {
	var field = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'9', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	log.Println(isValidSudoku(field))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !checkSquare(i, j, board) {
				return false
			}
		}
	}

	return checkColsAndRows(board)
}

func checkSquare(x, y int, grid [][]byte) bool {
	squareElements := make(map[byte]bool, 9)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if grid[i][j] != '.' {
				if _, ok := squareElements[grid[i][j]]; ok {
					return false
				}

				squareElements[grid[i][j]] = true
			}
		}
	}

	return true
}

func checkColsAndRows(grid [][]byte) bool {
	for j := 0; j < 9; j++ {
		colElements := make(map[byte]bool, 9)
		rowElements := make(map[byte]bool, 9)
		for i := 0; i < 9; i++ {
			if grid[i][j] != '.' {
				if _, ok := colElements[grid[i][j]]; ok {
					return false
				}
				colElements[grid[i][j]] = true
			}

			if grid[j][i] != '.' {
				if _, ok := rowElements[grid[j][i]]; ok {
					return false
				}
				rowElements[grid[j][i]] = true
			}
		}
	}

	return true
}
