package main

import (
	"reflect"
	"testing"
)

func Test_tetraminos(t *testing.T) {
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
            if got := tetraminos(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("tetraminos() = %v, want %v", got, tt.want)
            }
        })
    }
}
