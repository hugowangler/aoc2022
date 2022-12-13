package main

import (
	"testing"
)

var example = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		input string
		part  int
	}{
		{
			name:  "part 1",
			want:  31,
			input: example,
			part:  1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := part1(tt.input); got != tt.want {
					t.Errorf("part1() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
