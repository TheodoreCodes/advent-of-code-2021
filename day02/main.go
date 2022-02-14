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

func main() {
	f := utils.NewFile(inputFile)
	defer f.Close()

	x, y, aim := 0, 0, 0

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
	// We multiply x * aim because aim is behaving exactly as the depth in Part 1
	fmt.Println("Part 1:", x*aim)
	fmt.Println("Part 2:", x*y)
}
