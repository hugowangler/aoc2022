package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day09/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input), 10))
}

type Position struct {
	x int
	y int
}

func part1(input string) int {
	visited := map[Position]struct{}{}
	head := Position{0, 0}
	tail := Position{0, 0}
	visited[tail] = struct{}{}
	for _, motion := range strings.Split(input, "\n") {
		split := strings.Split(motion, " ")
		steps, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "R":
			for i := 0; i < steps; i++ {
				head.x++
				if touching(tail, head) {
					continue
				}
				if tail.y != head.y {
					tail.y = head.y
				}
				tail.x++
				visited[tail] = struct{}{}
			}
		case "L":
			for i := 0; i < steps; i++ {
				head.x--
				if touching(tail, head) {
					continue
				}
				if tail.y != head.y {
					tail.y = head.y
				}
				tail.x--
				visited[tail] = struct{}{}
			}
		case "U":
			for i := 0; i < steps; i++ {
				head.y++
				if touching(tail, head) {
					continue
				}
				if tail.x != head.x {
					tail.x = head.x
				}
				tail.y++
				visited[tail] = struct{}{}
			}
		case "D":
			for i := 0; i < steps; i++ {
				head.y--
				if touching(tail, head) {
					continue
				}
				if tail.x != head.x {
					tail.x = head.x
				}
				tail.y--
				visited[tail] = struct{}{}
			}
		}
	}
	return len(visited)
}

func part2(input string, numKnots int) int {
	visited := map[Position]struct{}{}
	var knots []*Position
	for i := 0; i < numKnots; i++ {
		knots = append(knots, &Position{0, 0})
	}
	visited[Position{0, 0}] = struct{}{}
	for _, motion := range strings.Split(input, "\n") {
		split := strings.Split(motion, " ")
		steps, _ := strconv.Atoi(split[1])
		for i := 0; i < steps; i++ {
			for n := 0; n < numKnots-1; n++ {
				head := knots[n]
				tail := knots[n+1]
				if n == 0 {
					switch split[0] {
					case "R":
						head.x++
					case "L":
						head.x--
					case "U":
						head.y++
					case "D":
						head.y--
					}
				}
				if touching(*tail, *head) {
					continue
				}
				dx, dy := move(tail, head)
				tail.x += dx
				tail.y += dy
				if n == numKnots-2 {
					visited[*tail] = struct{}{}
				}
			}
		}
	}
	return len(visited)
}

func move(tail *Position, head *Position) (int, int) {
	dx := head.x - tail.x
	dy := head.y - tail.y
	if dx != 0 {
		dx = int(float64(dx) / math.Abs(float64(dx)))
	}
	if dy != 0 {
		dy = int(float64(dy) / math.Abs(float64(dy)))
	}
	return dx, dy
}

func touching(tail Position, head Position) bool {
	if math.Abs(float64(tail.y-head.y)) <= 1 && math.Abs(float64(tail.x-head.x)) <= 1 {
		return true
	}
	return false
}
