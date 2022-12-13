package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day13/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(string(file)))
	fmt.Println("Part 2:", part2(string(file)))
}

func part1(input string) int {
	correctSum := 0
	pairs := strings.Split(input, "\n\n")
	for i, pair := range pairs {
		split := strings.Split(pair, "\n")
		left := parse(split[0])
		right := parse(split[1])
		check := checkPacket(left, right)
		if check == 1 {
			correctSum += i + 1
		}
	}
	return correctSum
}

func part2(input string) int {
	var packets []*Packet
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		packets = append(packets, parse(row))
	}
	d2 := &Packet{
		packageType: "LIST",
		children:    []*Packet{{packageType: "NUM", number: 2, isDivider: true}},
		isDivider:   true,
	}
	d6 := &Packet{
		packageType: "LIST",
		children:    []*Packet{{packageType: "NUM", number: 6, isDivider: true}},
		isDivider:   true,
	}
	packets = append(packets, d2, d6)
	sort.Slice(
		packets, func(i, j int) bool {
			return checkPacket(packets[i], packets[j]) == 1
		},
	)
	decoderKey := 1
	for i, p := range packets {
		if p.isDivider {
			decoderKey *= i + 1
		}
	}
	return decoderKey
}

type Packet struct {
	packageType string
	number      int
	children    []*Packet
	isDivider   bool
}

func checkPacket(left *Packet, right *Packet) int {
	if left.packageType == "NUM" && right.packageType == "NUM" {
		return checkNums(left.number, right.number)
	}
	if left.packageType == "LIST" && right.packageType == "LIST" {
		for i := 0; i < len(left.children); i++ {
			if len(right.children)-1 < i {
				return -1
			}
			result := checkPacket(left.children[i], right.children[i])
			if result != 0 {
				return result
			}
		}
		if len(right.children) == len(left.children) {
			return 0
		}
		// left ran out of items first
		return 1
	}
	if left.packageType == "NUM" {
		left.packageType = "LIST"
		left.children = []*Packet{{packageType: "NUM", number: left.number}}
		return checkPacket(left, right)
	}
	right.packageType = "LIST"
	right.children = []*Packet{{packageType: "NUM", number: right.number}}
	return checkPacket(left, right)
}

func checkNums(left int, right int) int {
	if left == right {
		return 0
	}
	if left < right {
		return 1
	}
	return -1
}

func parse(row string) *Packet {
	p := &Packet{packageType: "LIST", children: []*Packet{}}
	for i := 1; i < len(row); i++ {
		if row[i] == '[' {
			listsInLists := 1
			for j := i + 1; j < len(row); j++ {
				if row[j] == '[' {
					listsInLists++
				}
				if row[j] == ']' {
					listsInLists--
				}
				if listsInLists == 0 {
					p.children = append(p.children, parse(row[i:j+1]))
					i = j
					break
				}
			}
		} else if row[i] == ',' {
			continue
		} else {
			for j := i + 1; j < len(row); j++ {
				if row[j] == ',' || row[j] == ']' {
					p.children = append(p.children, getInt(row[i:j]))
					i = j
					break
				}
			}
		}
	}
	return p
}

func getInt(str string) *Packet {
	n := &Packet{packageType: "NUM", children: nil}
	n.number, _ = strconv.Atoi(str)
	return n
}
