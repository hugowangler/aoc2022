package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("day12/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(string(input)))
	fmt.Println("Part 2:", part2(string(input)))
}

func part1(input string) int{
	s, end, queue := parseHeightmap(input, 1)
	return shortestPath(s, end, queue, 1)
}
func part2(input string) int{
	_, end, queue := parseHeightmap(input, 2)
	return shortestPath(end, square{}, queue, 2)
}

type square struct {
	i      int
	j      int
	height int
}

func (s *square) isNeighbour(m square) bool {
	return (math.Abs(float64(s.i-m.i)) == 0 && math.Abs(float64(s.j-m.j)) <= 1) || (math.Abs(float64(s.i-m.i)) <= 1 && math.Abs(float64(s.j-m.j)) == 0)
}

func shortestPath(start square, end square, startQueue []square, part int) int {
	prev := make(map[square]square)
	dist := make(map[square]int)
	dist[start] = 0
	queue := make([]square, len(startQueue))
	copy(queue, startQueue)
	for len(queue) > 0 {
		u := dequeue(&queue, dist)
		if part == 1 && u == end {
			return dist[u]
		} else if part == 2 && u.height == 25 {
			return dist[u]
		}
		inQueue := neighboursInQueue(u, queue)
		for _, v := range inQueue {
			if v.height-u.height <= 1 {
				if _, exists := dist[v]; !exists {
					dist[v] = math.MaxInt
				}
				alt := dist[u] + 1
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}
	return math.MaxInt
}

func neighboursInQueue(u square, queue []square) []square {
	var neighbours []square
	for _, v := range queue {
		if u.isNeighbour(v) {
			neighbours = append(neighbours, v)
		}
	}
	return neighbours
}

func dequeue(queue *[]square, dist map[square]int) square {
	minDist := math.MaxInt
	var closest square
	minIndex := 0
	for i, q := range *queue {
		if _, exists := dist[q]; exists {
			if dist[q] <= minDist {
				minDist = dist[q]
				minIndex = i
				closest = q
			}
		}
	}
	(*queue)[minIndex] = (*queue)[len(*queue)-1]
	*queue = (*queue)[:len(*queue)-1]
	return closest
}

func parseHeightmap(input string, part int) (
	square,
	square,
	[]square,
) {
	var s, end square
	rows := strings.Split(input, "\n")
	var queue []square
	for i := range rows {
		for j := range rows[i] {
			var sq square
			if part == 1 {
				sq = square{i, j, int(rows[i][j] - 'a')}
			} else {
				sq = square{i, j, int('z' - rows[i][j])}
			}
			if rows[i][j] == 'S' {
				if part == 1 {
					sq.height = 0
				} else {
					sq.height = int('z' - 'a')
				}
				s = sq
			} else {
				if rows[i][j] == 'E' {
					if part == 1 {
						sq.height = 'z' - 'a'
					} else {
						sq.height = 0
					}
					end = sq
				}
			}
			queue = append(queue, sq)
		}
	}
	return s, end, queue
}
