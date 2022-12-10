package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day10/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:")
	part2(string(input))
}

func part1(input string) int {
	ops := strings.Split(input, "\n")
	offset := 0
	wait := 0
	x := 1
	add := 0
	signalStrength := 0
	for i := 0; i <= 220; i++ {
		if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
			signalStrength += x * i
		}
		if wait > 0 {
			wait--
			continue
		}
		x += add
		add = 0
		split := strings.Split(ops[offset], " ")
		switch split[0] {
		case "addx":
			val, _ := strconv.Atoi(split[1])
			add += val
			wait++
		}
		offset++
	}
	return signalStrength
}

func part2(input string) {
	ops := strings.Split(input, "\n")
	offset := 0
	wait := 0
	x := 1
	add := 0
	crt := make([]string, 6)
	for i := 0; i < 240; i++ {
		if wait == 0 {
			x += add
			add = 0
			split := strings.Split(ops[offset], " ")
			switch split[0] {
			case "addx":
				val, _ := strconv.Atoi(split[1])
				add += val
				wait++
			}
			offset++
		} else {
			wait--
		}
		index := int(math.Floor(float64(i / 40)))
		if i%40 >= x-1 && i%40 <= x+1 {
			crt[index] += "#"
		} else {
			crt[index] += "."
		}
	}
	for _, row := range crt {
		fmt.Println(row)
	}
}
