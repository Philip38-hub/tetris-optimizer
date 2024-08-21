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
			if got := tetrominosGenerator(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
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
		// {
		// 	name: "Test square tetromino",
		// 	args: args{
		// 		tetrominos: [][]string{
		// 			{"....", "BB..", "BB..", "...."},
		// 		},
		// 	},
		// 	want: [][]string{
		// 		{"BB", "BB"},
		// 	},
		// },
		// {
		// 	name: "Test horizontal line tetromino",
		// 	args: args{
		// 		tetrominos: [][]string{
		// 			{"....", "CCCC", "....", "...."},
		// 		},
		// 	},
		// 	want: [][]string{
		// 		{"CCCC"},
		// 	},
		// },
		// {
		// 	name: "Test vertical line tetromino",
		// 	args: args{
		// 		tetrominos: [][]string{
		// 			{"..D.", "..D.", "..D.", "..D."},
		// 		},
		// 	},
		// 	want: [][]string{
		// 		{"D", "D", "D", "D"},
		// 	},
		// },
		// {
		// 	name: "Test no trimming needed",
		// 	args: args{
		// 		tetrominos: [][]string{
		// 			{"AA", "AA"},
		// 		},
		// 	},
		// 	want: [][]string{
		// 		{"AA", "AA"},
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trimmer(tt.args.tetrominos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trimmer() = %v, want %v", got, tt.want)
			}
		})
	}
}
