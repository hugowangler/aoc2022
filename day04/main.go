package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	println("Part 1:", part1(input))
	println("Part 2:", part2(input))
}

func part1(rows []string) int {
	fullyContained := 0
	for _, row := range rows {
		a, b, x, y := getRanges(row)
		if a >= x && b <= y {
			fullyContained++
		} else if x >= a && y <= b {
			fullyContained++
		}
	}
	return fullyContained
}

func part2(rows []string) int {
	overlaps := 0
	for _, row := range rows {
		a, b, x, y := getRanges(row)
		if a >= x && a <= y {
			overlaps++
		} else if b >= x && b <= y {
			overlaps++
		} else if x >= a && x <= b {
			overlaps++
		} else if y >= a && y <= b {
			overlaps++
		}
	}
	return overlaps
}

func getRanges(row string) (int, int, int, int) {
	pair := strings.Split(row, ",")
	right := strings.Split(pair[0], "-")
	left := strings.Split(pair[1], "-")
	a, _ := strconv.Atoi(left[0])
	b, _ := strconv.Atoi(left[1])
	x, _ := strconv.Atoi(right[0])
	y, _ := strconv.Atoi(right[1])
	return a, b, x, y
}
