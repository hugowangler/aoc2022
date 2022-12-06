package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(rucksacks []string) int {
	prioritySum := 0
	for _, rucksack := range rucksacks {
		chars := make(map[int32]int)
		var duplicate int32
		for i, char := range rucksack {
			if i < len(rucksack)/2 && chars[char] == 0 {
				chars[char]++
			} else if i > len(rucksack)/2-1 && chars[char] == 1 {
				duplicate = char
				break
			}
		}
		if duplicate != 0 {
			if duplicate <= 'Z' {
				prioritySum += int(duplicate) - 38
			} else {
				prioritySum += int(duplicate) - 96
			}
		}
	}
	return prioritySum
}

func part2(rucksacks []string) int {
	prioritySum := 0
	chars := make(map[int32]int)
	for i, rucksack := range rucksacks {
		for _, char := range rucksack {
			if chars[char] == i%3 {
				chars[char]++
			}
			if chars[char] == 3 {
				if char <= 'Z' {
					prioritySum += int(char) - 38
				} else {
					prioritySum += int(char) - 96
				}
				chars = make(map[int32]int)
				break
			}
		}
	}
	return prioritySum
}
