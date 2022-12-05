package main

import (
	"fmt"
	"os"
	"strings"
)

var shapeScoreMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var meToOpponent = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var shapeResult = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var shapeScoreMapOpponent = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func main() {
	data, err := os.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}
	score := 0
	for _, round := range strings.Split(string(data), "\n") {
		score += part1(round[0:1], round[2:3])
	}
	fmt.Println("Part 1:", score)
	score = 0
	for _, round := range strings.Split(string(data), "\n") {
		score += part2(round[0:1], round[2:3])
	}
	fmt.Println("Part 2:", score)
}

func part1(opponent string, me string) int {
	shapeScore := shapeScoreMap[me]
	meABC := meToOpponent[me]
	if opponent == meABC {
		return shapeScore + 3
	}
	if shapeResult[meABC] == opponent {
		return shapeScore + 6
	}
	return shapeScore
}

func part2(opponent string, outcome string) int {
	switch outcome {
	case "X":
		return shapeScoreMapOpponent[shapeResult[opponent]]
	case "Y":
		return shapeScoreMapOpponent[opponent] + 3
	case "Z":
		pick := opponent[0] + 1
		if pick == 'C'+1 {
			pick = 'A'
		}
		return shapeScoreMapOpponent[string(pick)] + 6
	}
	return 0
}
