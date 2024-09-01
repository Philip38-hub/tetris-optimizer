package Tetrominos

import (
	"fmt"
	"os"
	"strings"
)

func TetrominosGenerator(lines []string) [][]string {
	// read tetrominos and store in a 2D slice
	var tetromino []string
	var tetrominos [][]string
	letter := 'A'
	for _, line := range lines {
		t := ""
		if line != "" {
			for _, char := range line {
				if char == '#' {
					t += string(letter)
				} else if char == '.' {
					t += string(char)
				} else {
					fmt.Printf("Invalid tetromino character: %c ", char)
					os.Exit(1)
				}
			}
			tetromino = append(tetromino, t)
		} else {
			tetrominos = append(tetrominos, tetromino)
			tetromino = []string{} // Reset for the next tetromino
			letter++
		}
	}
	if len(tetromino) > 0 {
		tetrominos = append(tetrominos, tetromino)
	}
	return tetrominos
}

func Validator(s []string) error {
	if len(s) != 4 {
		return fmt.Errorf("tetromino %v does not have 4 columns", s)
	}
	countconn := 0
	countchar := 0
	for i, row := range s {
		if len(row) != 4 {
			return fmt.Errorf("tetromino %v has %v rows", row, len(row))
		}

		for j, char := range row {
			if char != '.' {
				countchar++
				if i > 0 && s[i-1][j] == byte(char) {
					countconn++
				}
				if i < len(s)-1 && s[i+1][j] == byte(char) {
					countconn++
				}
				if j > 0 && s[i][j-1] == byte(char) {
					countconn++
				}
				if j < len(row)-1 && s[i][j+1] == byte(char) {
					countconn++
				}
			}
		}
	}
	if countconn < 6 || countchar != 4 {
		return fmt.Errorf("tetromino %v is invalid", s)
	}
	return nil
}

func Trimmer(tetrominos [][]string) [][]string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// remove strings with no letters i.e rows
	var newTetrominos [][]string
	for _, tetromino := range tetrominos {
		newTetromino := []string{}
		for _, str := range tetromino {
			for _, letter := range letters {
				if strings.Contains(str, string(letter)) {
					newTetromino = append(newTetromino, str)
				}
			}
		}
		newTetrominos = append(newTetrominos, newTetromino)
	}

	// remove columns without letters in them
	var newTetrominos2 [][]string
	for _, tetromino := range newTetrominos {
		width := len(tetromino[0])

		// Check each column for the presence of letters ('A')
		columnHasLetters := make([]bool, width)
		for col := 0; col < width; col++ {
			for row := 0; row < len(tetromino); row++ {
				for _, letter := range letters {
					if tetromino[row][col] == byte(letter) {
						columnHasLetters[col] = true
						break
					}
				}
			}
		}

		// Remove columns that do not contain any letters ('A')
		var result []string
		for _, row := range tetromino {
			var newRow strings.Builder
			for col := 0; col < width; col++ {
				if columnHasLetters[col] {
					newRow.WriteByte(row[col])
				}
			}
			result = append(result, newRow.String())
		}
		newTetrominos2 = append(newTetrominos2, result)
	}
	return newTetrominos2
}
