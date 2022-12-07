package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type move struct {
	crates int
	from   string
	to     string
}

func main() {
	data, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}
	println("Part 1:", part1(string(data)))
	println("Part 2:", part2(string(data)))
}

func parseDrawing(input string) (map[string][]string, []move) {
	split := strings.Split(input, "\n\n")
	strStartingStacks := split[0]
	strMoves := split[1]
	stacks := make(map[string][]string)
	rows := strings.Split(strStartingStacks, "\n")
	for _, row := range rows[:len(rows)-1] {
		for j := 1; j < len(row); j += 4 {
			if string(row[j]) != " " {
				index := strconv.Itoa(int(math.Ceil(float64(j) / 4.0)))
				stacks[index] = append(stacks[index], string(row[j]))
			}
		}
	}
	var moves []move
	rows = strings.Split(strMoves, "\n")
	for _, row := range rows {
		tmp := strings.Split(row, " ")
		crates, _ := strconv.Atoi(tmp[1])
		moves = append(moves, move{crates, tmp[3], tmp[5]})
	}
	return stacks, moves
}

func part1(input string) string {
	stacks, moves := parseDrawing(input)
	for _, move := range moves {
		for _, crate := range stacks[move.from][:move.crates] {
			stacks[move.to] = append([]string{crate}, stacks[move.to]...)
		}
		stacks[move.from] = stacks[move.from][move.crates:]
	}
	top := ""
	for i := 1; i <= len(stacks); i++ {
		top += stacks[strconv.Itoa(i)][0]
	}
	return top
}

func part2(input string) string {
	stacks, moves := parseDrawing(input)
	for _, move := range moves {
		moving := stacks[move.from][:move.crates]
		for i := len(moving) - 1; i >= 0; i-- {
			stacks[move.to] = append([]string{moving[i]}, stacks[move.to]...)
		}
		stacks[move.from] = stacks[move.from][move.crates:]
	}
	top := ""
	for i := 1; i <= len(stacks); i++ {
		top += stacks[strconv.Itoa(i)][0]
	}
	return top
}
