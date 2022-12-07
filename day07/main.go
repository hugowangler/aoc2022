package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("day07/input.txt")
	if err != nil {
		panic(err)
	}
	rootDir := part1(string(input))
	answer := rootDir.getSize(100_000)
	fmt.Println("Part 1:", answer)
	fmt.Println("Part 2:", part2(rootDir))
}

type Directory struct {
	dirname  string
	parent   *Directory
	children *[]*Directory
	files    *[]*File
	size     int
}

func (d *Directory) getSize(n int) int {
	totalSize := 0
	if d.children != nil {
		for _, child := range *d.children {
			size := child.getSize(n)
			totalSize += size
			d.size += child.size
		}
	}
	if d.files != nil {
		for _, file := range *d.files {
			d.size += file.size
		}
	}
	if d.size <= n {
		totalSize += d.size
	}
	return totalSize
}

func (d *Directory) findAtLeastSize(n int) int {
	closest := math.MaxInt
	if d.size >= n {
		closest = d.size
	}
	if d.children != nil {
		for _, child := range *d.children {
			closestInChild := child.findAtLeastSize(n)
			if closestInChild < closest {
				closest = closestInChild
			}
		}
	}
	return closest
}

type File struct {
	name string
	size int
}

func part1(input string) *Directory {
	currDir := &Directory{
		"/",
		nil,
		nil,
		nil,
		0,
	}
	commandsAndOutput := strings.Split(input, "$ ")[1:]
	for _, c := range commandsAndOutput {
		cmd, arg, output := parseCommand(c)
		switch cmd {
		case "cd":
			switch arg {
			case "/":
				for currDir.parent != nil {
					currDir = currDir.parent
				}
			case "..":
				currDir = currDir.parent
			default:
				for _, child := range *currDir.children {
					if child.dirname == arg {
						currDir = child
						break
					}
				}
			}
		case "ls":
			for _, row := range output {
				split := strings.Split(row, " ")
				if split[0] == "dir" {
					dir := &Directory{dirname: split[1], parent: currDir}
					if currDir.children == nil {
						currDir.children = &[]*Directory{dir}
					} else {
						children := append(*currDir.children, dir)
						currDir.children = &children
					}
				} else {
					size, _ := strconv.Atoi(split[0])
					file := &File{name: split[1], size: size}
					if currDir.files == nil {
						currDir.files = &[]*File{file}
					} else {
						files := append(*currDir.files, file)
						currDir.files = &files
					}
				}
			}
		}
	}
	for currDir.parent != nil {
		currDir = currDir.parent
	}
	return currDir
}

func part2(root *Directory) int {
	unused := 70000000 - root.size
	missing := 30000000 - unused
	return root.findAtLeastSize(missing)
}

func parseCommand(output string) (string, string, []string) {
	split := strings.SplitN(output, "\n", 2)
	cmdRow := strings.Split(split[0], " ")
	if len(cmdRow) > 1 {
		return cmdRow[0], cmdRow[1], nil
	}
	out := strings.Split(split[1], "\n")
	return split[0], "", out[:len(out)-1]
}
