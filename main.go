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
	return tetraminos
}
