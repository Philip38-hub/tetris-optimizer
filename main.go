package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	Tetrominos "tetrominos/packages"
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
	// generate tetrominos
	tetrominos := Tetrominos.TetrominosGenerator(lines)
	// validate that each tetramino has 4 #s, 4 rows, 4 columns and 6 or more connections
	for _, tetromino := range tetrominos {
		if err := Tetrominos.Validator(tetromino); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// trim tetrominos
	tetrominos = Tetrominos.Trimmer(tetrominos)

	// create board and solve for optimization
	sizeOfBoard := int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4))))
	var final [][]string
	for {
		board := Tetrominos.CreateBoard(sizeOfBoard)
		final = Tetrominos.Solve(board, tetrominos)
		if final != nil {
			break
		}
		sizeOfBoard++
	}

	// print board
	Tetrominos.PrintBoard(final)
}
