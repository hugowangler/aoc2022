package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part 1",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  7,
		},
		{
			name:  "part 1.2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
	}
	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		t.Run(
			tt.name, func(t *testing.T) {
				if got := part1(reader); got != tt.want {
					t.Errorf("part1() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "part 2",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
		{
			name:  "part 2.2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  23,
		},
		{
			name:  "part 2.3",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  23,
		},
		{
			name:  "part 2.4",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  29,
		},
		{
			name:  "part 2.5",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  26,
		},
	}
	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		t.Run(
			tt.name, func(t *testing.T) {
				if got := part2(reader); got != tt.want {
					t.Errorf("part2() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
