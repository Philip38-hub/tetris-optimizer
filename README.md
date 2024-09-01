# Tetris Optimizer
## Description
Tetris solver and optimizer is  a program that receives only one argument, a path to a text file which will contain a list of [tetrominos](https://en.wikipedia.org/wiki/Tetromino) and assembles them in order to create the smallest square possible.
## Implementation details

- **main.go**: The entry point of the program, responsible for creating reading file name containing tetraminos from terminal and reads it line by line.
- **algorithm.go**: A package that handles creation of tetris board, provides an optimized solution.
- **tetroprocessor.go**: A package that checks if tetros are valid and trims the extra spaces.

### Example of a text File
```
#...
#...
#...
#...

....
....
..##
..##
```
### Expected Output for the above Example
```
ABB.
ABB.
A...
A...
```
## Guidelines for tetromino Text files Format

1. The file **must** have a .txt extension e.g text.txt, mug.txt
2. The tetrominos **must** have a length of 4 i.e
```
...# 1
...# 2
...# 4
...# 4
```
3. They **must** be separated by an empty line i.e
```
...#
...#
...#
...#

...#
...#
...#
...#
```
4. The length of each string **must** be 4
5. There **must** only be 4 hashes(#) to represent a tetromino


## Running

- This application requires Go (golang) 1.18 or higher to run. You can get it [here](https://go.dev/doc/install)

- To clone and run this program, you'll need **Git** installed on your computer.

- From the **command line**,

```Bash
# clone this repository
$ git clone https://github.com/Philip38-hub/tetris-optimizer.git

# go into the repository
$ cd tetris-optimizer

# open terminal
$ go run . <textfilename>.txt
```

## Usage

`$ go run . <textfilename>.txt`

## Notion

[Fillit: Solving for the Smallest Square of Tetrominoes](https://medium.com/@bethnenniger/fillit-solving-for-the-smallest-square-of-tetrominos-c6316004f909)

## Enjoy!


## Author
- **[X - @oumaphilip01](https://x.com/oumaphilip01)**
- **[Github - Philip38](https://github.com/Philip38-hub)**

<p align="right">(<a href="#tetris-optimizer">back to top</a>)</p>
