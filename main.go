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

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 4 {
		fmt.Printf("Expect atleast one tetromino in %s", os.Args[1])
		os.Exit(0)
	}
	// group lines read into 9 to display each character individually
	tetramino := make([]string, 5)
	var tetraminos [][]string
	for i := 0; i < len(lines); i += 4 {
		end := i + 4
		if end > len(lines) {
			end = len(lines)
		}
		tetramino = lines[i:end]
		tetraminos = append(tetraminos, tetramino[:4])
	}

	// r := flag.Args()[0]
	// // Format ("/n") in input string
	// s := strings.Replace(r, "\\n", "\n", -1)
	// s = strings.Replace(s, "\\t", "    ", -1)
	// asciiart.HandleLn(s, result, outputfile)
}
