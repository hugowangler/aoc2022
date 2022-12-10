package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		input string
		want int
	}{
		{
			name:  "part 1",
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			want:  13,
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

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		input string
		want int
	}{
		{
			name:  "part 2",
			input: `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`,
			want:  1,
		},
		{
			name:  "part 2.2",
			input: `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`,
			want:  36,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := part2(tt.input, 10); got != tt.want {
					t.Errorf("part1() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
