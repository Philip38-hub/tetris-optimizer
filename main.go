package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename.txt>")
		os.Exit(0)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Unable to open file!")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	// read file line by line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, tetramino := range tetraminos(lines) {
		if err := validator(tetramino); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		// for _, row := range tetramino {
		// 	fmt.Println(row)
		// }
		// fmt.Println()
	}
}

func tetraminos(lines []string) [][]string {
	// read tetrominos and store in a 2D slice
	var tetramino []string
	var tetraminos [][]string
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
					fmt.Printf("Invalid tetramino character: %c ", char)
					os.Exit(1)
				}
			}
			tetramino = append(tetramino, t)
		} else {
			tetraminos = append(tetraminos, tetramino)
			tetramino = []string{} // Reset for the next tetromino
			letter++
		}
	}
	if len(tetramino) > 0 {
        tetraminos = append(tetraminos, tetramino)
    }
	return tetraminos
}

func validator (s []string) error {
	if len(s) != 4 {
		return fmt.Errorf("tetramino %v does not have 4 columns", s)
	}
	countconn := 0
	countchar := 0
	for i, row := range s {
		if len(row) != 4 {
			return fmt.Errorf("tetramino %v has %v rows", row, len(row))
		}

		for j, char := range row{
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
		return fmt.Errorf("tetramino %v is invalid", s)
	}
	return nil
}
