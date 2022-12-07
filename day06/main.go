package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day06/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	fmt.Println("Part 1:", part1(reader))
	file.Close()
	file, err = os.Open("day06/input.txt")
	if err != nil {
		panic(err)
	}
	reader = bufio.NewReader(file)
	fmt.Println("Part 2:", part2(reader))
}

func part1(reader *bufio.Reader) int {
	var uniqueInRow []rune
	runesRead := 0
	for {
		r, _, _ := reader.ReadRune()
		runesRead++
		uniqueInRow = append(uniqueInRow, r)
		for i := 0; i < len(uniqueInRow)-1; i++ {
			if len(uniqueInRow) > 1 && uniqueInRow[i] == r {
				uniqueInRow = uniqueInRow[i+1:]
			}
		}
		if len(uniqueInRow) == 4 {
			return runesRead
		}
	}
}

func part2(reader *bufio.Reader) int {
	var uniqueInRow []rune
	runesRead := 0
	for {
		r, _, _ := reader.ReadRune()
		runesRead++
		uniqueInRow = append(uniqueInRow, r)
		for i := 0; i < len(uniqueInRow)-1; i++ {
			if len(uniqueInRow) > 1 && uniqueInRow[i] == r {
				uniqueInRow = uniqueInRow[i+1:]
			}
		}
		if len(uniqueInRow) == 14 {
			return runesRead
		}
	}
}
