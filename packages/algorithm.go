package Tetrominos

import "fmt"

// create a board with dimensions of length == number of tetraminos
func CreateBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// Solve attempts to place tetrominosin the smallest square possible using recursive method.
func Solve(board [][]string, tetrominos [][]string) [][]string {
	if Solvetetro(board, tetrominos, 0) {
		return board
	}
	return nil
}

// finally print the board
func PrintBoard(board [][]string) {
	for _, row := range board {
		for _, char := range row {
			fmt.Printf("%s ", char)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Solvetetro(board [][]string, tetrominos [][]string, index int) bool {
	// checks if all tetrominos have been exhausted
	if index == len(tetrominos) {
		return true
	}

	tetromino := tetrominos[index]
	// Try to place the current tetromino at every possible position
	for y := range board {
		for x := range board[y] {
			if canPlace(board, tetromino, x, y) {
				placeTetromino(board, tetromino, x, y)
				// Recursively try to place the next tetromino
				if Solvetetro(board, tetrominos, index+1) {
					return true
				}
				// If unsuccessful remove the current tetromino and try the next position
				removeTetromino(board, tetromino, x, y)
			}
		}
	}
	return false
}

// canPlace checks if a tetromino can be placed at a given position on the board
func canPlace(board [][]string, tetromino []string, x, y int) bool {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				// Check if the position is within the board and empty
				if y+dy >= len(board) || x+dx >= len(board[0]) || board[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

// placeTetromino places a tetromino on the board at a given position
func placeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

// removeTetromino removes a tetromino from the board at a given position
func removeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}
