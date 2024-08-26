package main

import (
	"reflect"
	"testing"
)

func Test_tetrominosGenerator(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test single tetromino",
			args: args{
				lines: []string{
					"...#",
					"...#",
					"...#",
					"...#",
					"",
					"....",
					"....",
					"..##",
					"..##",
				},
			},
			want: [][]string{
				{
					"...A",
					"...A",
					"...A",
					"...A",
				},
				{
					"....",
					"....",
					"..BB",
					"..BB",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TetrominosGenerator(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tetrominos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimmer(t *testing.T) {
	type args struct {
		tetrominos [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test L-shaped tetromino",
			args: args{
				tetrominos: [][]string{
					{"....", "A...", "A...", "AA.."},
				},
			},
			want: [][]string{
				{"A.", "A.", "AA"},
			},
		},
		{
			name: "Test square tetromino",
			args: args{
				tetrominos: [][]string{
					{"....", "BB..", "BB..", "...."},
				},
			},
			want: [][]string{
				{"BB", "BB"},
			},
		},
		{
			name: "Test horizontal line tetromino",
			args: args{
				tetrominos: [][]string{
					{"....", "CCCC", "....", "...."},
				},
			},
			want: [][]string{
				{"CCCC"},
			},
		},
		{
			name: "Test vertical line tetromino",
			args: args{
				tetrominos: [][]string{
					{"..D.", "..D.", "..D.", "..D."},
				},
			},
			want: [][]string{
				{"D", "D", "D", "D"},
			},
		},
		{
			name: "Test no trimming needed",
			args: args{
				tetrominos: [][]string{
					{"AA", "AA"},
				},
			},
			want: [][]string{
				{"AA", "AA"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trimmer(tt.args.tetrominos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trimmer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateBoard(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 2x2 board",
			args: args{size: 2},
			want: [][]string{
				{".", "."},
				{".", "."},
			},
		},
		{
			name: "Test 3x3 board",
			args: args{size: 3},
			want: [][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
		},
		{
			name: "Test 4x4 board",
			args: args{size: 4},
			want: [][]string{
				{".", ".", ".", "."},
				{".", ".", ".", "."},
				{".", ".", ".", "."},
				{".", ".", ".", "."},
			},
		},
		{
			name: "Test 1x1 board",
			args: args{size: 1},
			want: [][]string{
				{"."},
			},
		},
		{
			name: "Test 0x0 board",
			args: args{size: 0},
			want: [][]string{}, // Expecting an empty board
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateBoard(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
