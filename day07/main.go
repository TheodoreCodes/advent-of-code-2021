package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
	"math"
	"strconv"
	"strings"
)

const (
	inputFile = "day07/input.txt"
)

func parseInput(filename string) []int {
	f := utils.NewFile(filename)
	defer f.Close()

	var initialPositions []int

	positionsStr := strings.Split(f.Read(), ",")
	for _, posStr := range positionsStr {
		pos, _ := strconv.Atoi(posStr)
		initialPositions = append(initialPositions, pos)
	}

	return initialPositions
}

func align(positions []int, desiredPosition int) int {
	var fuelConsumed int

	for _, pos := range positions {
		fuelConsumed += utils.Abs(desiredPosition - pos)
	}

	return fuelConsumed
}

func alignGauss(positions []int, desiredPosition int) int {
	fuelConsumed := 0

	for _, pos := range positions {
		positionsMoved := utils.Abs(desiredPosition - pos)
		fuelConsumed += (positionsMoved * (positionsMoved + 1)) / 2
	}

	return fuelConsumed
}

func part1(positions []int) int {
	minFuel := math.MaxInt
	for _, position := range positions {
		currFuel := align(positions, position)
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	return minFuel
}

func part2(positions []int) int {
	minFuel := math.MaxInt

	minPos := utils.Min(positions...)
	maxPos := utils.Max(positions...)

	for desiredPosition := minPos; desiredPosition <= maxPos; desiredPosition++ {
		currFuel := alignGauss(positions, desiredPosition)

		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	return minFuel
}

func run(positions []int) (int, int) {
	return part1(positions), part2(positions)
}

func main() {
	positions := parseInput(inputFile)
	minFuel1, minFuel2 := run(positions)

	fmt.Println("Part 1:", minFuel1)
	fmt.Println("Part 2:", minFuel2)
}
