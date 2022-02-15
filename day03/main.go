package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
)

const (
	inputFile = "day03/input.txt"
	width     = 12
)

func part1() (uint, uint) {
	f := utils.NewFile(inputFile)
	defer f.Close()

	var counters [width]int8

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
		if val > 0 {
			gamma += 1 << (width - pos - 1)
		} else {
			epsilon += 1 << (width - pos - 1)
		}
	}

	return gamma, epsilon
}

func part2() (uint, uint) {
	f := utils.NewFile(inputFile)
	var lines []uint
	for f.Scan() {
		lines = append(lines, uint(f.Base2Int()))
	}

	o2 := filterNumbers(lines, width-1, true)
	co2 := filterNumbers(lines, width-1, false)

	return o2, co2
}

func filterNumbers(numbers []uint, pos int, isMore bool) uint {
	if len(numbers) == 1 {
		return numbers[0]
	}

	var zeros, ones []uint
	bitmask := uint(1 << pos)

	for _, n := range numbers {
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
	gamma, epsilon := part1()
	fmt.Println("Part 1:", gamma*epsilon)

	o2, co2 := part2()
	fmt.Println("Part 1:", o2*co2)
}
