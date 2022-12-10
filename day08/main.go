package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("day08/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func part1(input string) int {
	visible := 0
	rows := strings.Split(input, "\n")
	for i := range rows {
		for j := range rows[i] {
			if i == 0 || j == 0 || i == len(rows)-1 || j == len(rows[i])-1 {
				visible++
				continue
			}
			checking := int(rows[i][j] - '0')
			// check south
			isVisible := true
			for p := i + 1; p < len(rows); p++ {
				if int(rows[p][j] - '0') >= checking {
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
			// check east
			isVisible = true
			for p := j + 1; p < len(rows[i]); p++ {
				if int(rows[i][p] - '0') >= checking {
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
			// check west
			isVisible = true
			for p := j - 1; p >= 0; p-- {
				if int(rows[i][p] - '0') >= checking {
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
			// check north
			isVisible = true
			for p := i - 1; p >= 0; p-- {
				if int(rows[p][j] - '0') >= checking {
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
		}
	}
	return visible
}

func part2(input string) int {
	rows := strings.Split(input, "\n")
	maxScenicScore := 0
	for i := range rows {
		for j := range rows[i] {
			scenicScore := 1
			if i == 0 || j == 0 || i == len(rows)-1 || j == len(rows[i])-1 {
				continue
			}
			checking := int(rows[i][j] - '0')
			// check south
			visible := 0
			for p := i + 1; p < len(rows); p++ {
				visible++
				if int(rows[p][j] - '0') >= checking {
					break
				}
			}
			scenicScore *= visible
			// check east
			visible = 0
			for p := j + 1; p < len(rows[i]); p++ {
				height := int(rows[i][p] - '0')
				visible++
				if height >= checking {
					break
				}
			}
			scenicScore *= visible
			// check west
			visible = 0
			for p := j - 1; p >= 0; p-- {
				visible++
				if int(rows[i][p] - '0') >= checking {
					break
				}
			}
			scenicScore *= visible
			// check north
			visible = 0
			for p := i - 1; p >= 0; p-- {
				visible++
				if int(rows[p][j] - '0') >= checking {
					break
				}
			}
			scenicScore *= visible
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}
