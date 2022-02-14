package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
	"math"
)

const (
	inputFile = "./day01/input.txt"
)

func part1() int {
	f := utils.NewFile(inputFile)
	defer f.Close()

	c := 0
	prev := math.MaxInt

	for f.Scan() {
		curr := f.Int()
		if curr > prev {
			c++
		}
		prev = curr
	}

	return c
}

func part2() int {
	f := utils.NewFile(inputFile)
	defer f.Close()

	// We compare:
	// 		x0 + x1 + x2 < x1 + x2 + x3 | -x1-x2 (common terms)
	// which is equivalent to:
	//		x0 < x3
	//
	// In our case x0 is *curr*
	// and x3 is *prev3*
	prev1, prev2, prev3 := math.MaxInt, math.MaxInt, math.MaxInt
	c := 0
	for f.Scan() {
		curr := f.Int()
		if prev3 < curr {
			c++
		}

		prev1, prev2, prev3 = curr, prev1, prev2
	}

	return c
}

func main() {
	c1 := part1()
	fmt.Println("Part 1:", c1)

	c2 := part2()
	fmt.Println("Part 2:", c2)

}
