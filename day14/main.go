package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day14/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(string(file)))
	fmt.Println("Part 2:", part2(string(file)))
}

func part1(input string) int {
	grid, maxY := parse(input)
	showGrid(grid)
	sandInPlace := 0
	done := false
	for {
		sand := Point{500, 0}
		for {
			if sand.y == maxY {
				done = true
				break
			}
			if _, exists := grid[Point{sand.x, sand.y + 1}]; !exists {
				sand.y++
				continue
			}
			if _, exists := grid[Point{sand.x - 1, sand.y + 1}]; !exists {
				sand.x--
				sand.y++
				continue
			}
			if _, exists := grid[Point{sand.x + 1, sand.y + 1}]; !exists {
				sand.x++
				sand.y++
				continue
			}
			grid[sand] = "O"
			break
		}
		if done {
			break
		}
		sandInPlace++
	}
	return sandInPlace
}

func part2(input string) int {
	grid, maxY := parse(input)
	maxY += 1
	showGrid(grid)
	sandInPlace := 0
	for {
		sand := Point{500, 0}
		for {
			if sand.y == maxY {
				grid[sand] = "O"
				break
			}
			if _, exists := grid[Point{sand.x, sand.y + 1}]; !exists {
				sand.y++
				continue
			}
			if _, exists := grid[Point{sand.x - 1, sand.y + 1}]; !exists {
				sand.x--
				sand.y++
				continue
			}
			if _, exists := grid[Point{sand.x + 1, sand.y + 1}]; !exists {
				sand.x++
				sand.y++
				continue
			}
			grid[sand] = "O"
			break
		}
		sandInPlace++
		if sand.x == 500 && sand.y == 0 {
			break
		}
	}
	showGrid(grid)
	return sandInPlace
}

type Point struct {
	x, y int
}

func parse(input string) (map[Point]string, int) {
	grid := map[Point]string{}
	maxY := 0
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		paths := strings.Split(row, " -> ")
		for i := 0; i < len(paths)-1; i++ {
			p1 := getPoint(paths[i])
			p2 := getPoint(paths[i+1])
			dx := p1.x - p2.x
			dy := p1.y - p2.y
			if p1.y > maxY {
				maxY = p1.y
			} else if p2.y > maxY {
				maxY = p2.y
			}
			direction := 1
			iter := 0
			directionX := int(math.Abs(float64(dx))) > 0
			if directionX {
				if dx > 0 {
					direction = -1
				}
				iter = int(math.Abs(float64(dx)))
			} else {
				if dy > 0 {
					direction = -1
				}
				iter = int(math.Abs(float64(dy)))
			}
			for j := 0; j <= iter; j++ {
				if directionX {
					grid[Point{x: p1.x + direction*j, y: p1.y}] = "#"
				} else {
					grid[Point{x: p1.x, y: p1.y + direction*j}] = "#"
				}
			}
		}
	}
	return grid, maxY
}

func showGrid(grid map[Point]string) {
	for i := 0; i < 11; i++ {
		for j := 476; j < 525; j++ {
			if i == 0 && j == 500 {
				fmt.Printf("+")
			} else if s, exists := grid[Point{j, i}]; exists {
				fmt.Printf("%s", s)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
		if i == 10 {
			fmt.Printf("#################################################\n\n")
		}
	}
}

func getPoint(str string) Point {
	tmp := strings.Split(str, ",")
	x, _ := strconv.Atoi(tmp[0])
	y, _ := strconv.Atoi(tmp[1])
	return Point{x, y}
}
