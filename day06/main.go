package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
	"strconv"
	"strings"
)

const (
	inputFile = "day06/input.txt"

	lanternInitialTimer = 8
)

func parseInput(filename string) []int {
	f := utils.NewFile(filename)
	defer f.Close()

	var initialState = make([]int, lanternInitialTimer+1)

	lanterns := strings.Split(f.Read(), ",")
	for _, l := range lanterns {
		val, _ := strconv.Atoi(l)
		initialState[val]++
	}

	return initialState
}

func passDay(state *[]int) {
	resetLanterns := (*state)[0]
	*state = (*state)[1:] // shift state to the left

	(*state)[lanternInitialTimer-2] += resetLanterns // reset all lanterns with internal state 0
	*state = append(*state, resetLanterns)           // all reset lanterns spawn new lanterns with internal state 8
}

func run(state []int) (int, int) {
	const (
		daysToPassPart1 = 80
		daysToPassPart2 = 256
	)
	var (
		sumPart1 int
		sumPart2 int
	)
	for d := 1; d <= daysToPassPart2; d++ {
		passDay(&state)
		if d == daysToPassPart1 {
			for _, i := range state {
				sumPart1 += i
			}

		} else if d == daysToPassPart2 {
			for _, i := range state {
				sumPart2 += i
			}
		}
	}

	return sumPart1, sumPart2
}

func main() {
	state := parseInput(inputFile)

	sumPart1, sumPart2 := run(state)

	fmt.Println("Part 1:", sumPart1)
	fmt.Println("Part 2:", sumPart2)
}
