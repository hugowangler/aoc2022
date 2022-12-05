package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(data), "\n")
	maxSoFar := 0
	currSum := 0
	for _, r := range rows {
		if r != "" {
			curr, _ := strconv.Atoi(r)
			currSum += curr
			continue
		}
		if currSum > maxSoFar {
			maxSoFar = currSum
		}
		currSum = 0
	}
	fmt.Println("Part 1:", maxSoFar)
	fmt.Println("Part 2:", GetNSum(rows, 3))
}

func GetNSum(rows []string, n int) int {
	maxSum := 0
	maxElems := make([]int, n)
	currSum := 0
	for _, r := range rows {
		if r != "" {
			curr, _ := strconv.Atoi(r)
			currSum += curr
			continue
		}
		if currSum > maxElems[0] {
			maxSum -= maxElems[0]
			maxSum += currSum
			maxElems[0] = currSum
			sort.Ints(maxElems)
		}
		currSum = 0
	}
	return maxSum
}
