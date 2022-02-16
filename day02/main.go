package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
	"strconv"
	"strings"
)

const (
	inputFile = "day02/input.txt"
)

func part1(filename string) (x, y int) {
	f := utils.NewFile(filename)
	defer f.Close()

	x, y = 0, 0

	for f.Scan() {
		l := f.Text()
		val, _ := strconv.Atoi(strings.Fields(l)[1])
		switch strings.Fields(l)[0] {
		case "forward":
			x += val
		case "down":
			y += val
		case "up":
			y -= val
		}
	}

	return x, y
}

func part2(filename string) (x, y int) {
	f := utils.NewFile(filename)
	defer f.Close()

	x, y = 0, 0
	aim := 0

	for f.Scan() {
		l := f.Text()
		val, _ := strconv.Atoi(strings.Fields(l)[1])
		switch strings.Fields(l)[0] {
		case "forward":
			x += val
			y += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return x, y
}

func main() {
	x, y := part1(inputFile)
	fmt.Println("Part 1:", x*y)

	x, y = part2(inputFile)
	fmt.Println("Part 2:", x*y)
}
