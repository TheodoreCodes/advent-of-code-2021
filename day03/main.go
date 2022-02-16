package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
)

const (
	inputFile = "day03/input.txt"
	width     = 12 // number of bits in each binary input
)

/**
* For part 1 binary conversion is not required. We keep track of how many lines have the characters '0' and '1'
* on each position.
* If counter value is > 0 for , that means more '1's were found at that position
 */
func part1(filename string, inputWidth int) (uint, uint) {
	f := utils.NewFile(filename)
	defer f.Close()

	var counters = make([]int8, inputWidth)

	for f.Scan() {
		line := f.Text()
		for pos, chr := range line {
			switch chr {
			case '1':
				counters[pos]++
				break
			case '0':
				counters[pos]--
				break
			}
		}
	}

	var gamma, epsilon uint
	for pos, val := range counters {
		if val > 0 { // If val > 0, '1' was more common for this pos
			gamma += 1 << (inputWidth - 1 - pos) // Left shift bits by `pos`
		} else { // Else, '0' was more common
			epsilon += 1 << (inputWidth - 1 - pos)
		}
	}

	return gamma, epsilon
}

func part2(filename string, inputWidth int) (uint, uint) {
	f := utils.NewFile(filename)
	var lines []uint
	for f.Scan() {
		lines = append(lines, uint(f.BinaryInt()))
	}

	o2 := filterNumbers(lines, inputWidth-1, true)
	co2 := filterNumbers(lines, inputWidth-1, false)

	return o2, co2
}

// Boolean `isMore` represents whether the search targets the most common bit, or the least common bit
func filterNumbers(numbers []uint, pos int, isMore bool) uint {
	// If list is size 1, we found our number
	if len(numbers) == 1 {
		return numbers[0]
	}

	var zeros, ones []uint
	bitmask := uint(1 << pos) // left shift *pos* number of bits

	for _, n := range numbers {
		// Binary AND
		if n&bitmask >= bitmask {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}

	if (isMore && len(ones) >= len(zeros)) ||
		(!isMore) && len(ones) < len(zeros) {
		return filterNumbers(ones, pos-1, isMore)
	}

	return filterNumbers(zeros, pos-1, isMore)
}

func main() {
	gamma, epsilon := part1(inputFile, width)
	fmt.Println("Part 1:", gamma*epsilon)

	o2, co2 := part2(inputFile, width)
	fmt.Println("Part 2:", o2*co2)
}
