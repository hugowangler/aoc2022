package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day11/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part(string(input), 1))
	fmt.Println("Part 2:", part(string(input), 2))
}

type monkey struct {
	items          []int
	onInspection   func(int) int
	test           int
	onTrue         int
	onFalse        int
	inspectedItems int
}

func (m *monkey) popItem() int {
	pop := m.items[0]
	m.items = m.items[1:]
	return pop
}

func part(input string, part int) int {
	var monkeys []*monkey
	var tests []int
	for _, m := range strings.Split(input, "\n\n") {
		mo := parseMonkey(m)
		if part == 2 {
			tests = append(tests, mo.test)
		}
		monkeys = append(monkeys, mo)
	}
	mod := 1
	if part == 2 {
		for _, t := range tests {
			mod *= t
		}
	}
	rounds := 20
	if part == 2 {
		rounds = 10_000
	}
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for range m.items {
				item := m.popItem()
				itemValue := m.onInspection(item)
				if part == 1 {
					itemValue = int(math.Floor(float64(itemValue) / 3.0))
				} else {
					itemValue = itemValue % mod
				}
				if itemValue%m.test == 0 {
					monkeys[m.onTrue].items = append(monkeys[m.onTrue].items, itemValue)
				} else {
					monkeys[m.onFalse].items = append(monkeys[m.onFalse].items, itemValue)
				}
				m.inspectedItems++
			}
		}
	}
	var inspected []int
	for _, m := range monkeys {
		inspected = append(inspected, m.inspectedItems)
	}
	sort.Slice(inspected, func(i, j int) bool { return inspected[i] > inspected[j] })
	return inspected[0] * inspected[1]
}

func parseMonkey(rows string) *monkey {
	split := strings.Split(rows, "\n")[1:]
	items := strings.Split(split[0], ": ")[1]
	var tmpItems []int
	var onInspection func(x int) int
	for _, strItem := range strings.Split(items, ", ") {
		item, _ := strconv.Atoi(strItem)
		tmpItems = append(tmpItems, item)
	}
	operation := strings.Split(split[1], ": ")[1]
	equation := strings.Split(operation, " ")
	onOld := false
	if equation[4] == "old" {
		onOld = true
	}
	switch equation[3] {
	case "*":
		if onOld {
			onInspection = func(x int) int {
				return x * x
			}
		} else {
			value, _ := strconv.Atoi(equation[4])
			onInspection = func(x int) int {
				return x * value
			}
		}
	case "+":
		value, _ := strconv.Atoi(equation[4])
		onInspection = func(x int) int {
			return x + value
		}
	}
	test, _ := strconv.Atoi(strings.Split(strings.TrimSpace(split[2]), " ")[3])
	onTrue, _ := strconv.Atoi(strings.Split(strings.TrimSpace(split[3]), " ")[5])
	onFalse, _ := strconv.Atoi(strings.Split(strings.TrimSpace(split[4]), " ")[5])
	return &monkey{
		items:        tmpItems,
		onInspection: onInspection,
		test:         test,
		onTrue:       onTrue,
		onFalse:      onFalse,
	}
}
