package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
)

const (
	inputFile = "day05/input.txt"
)

func parseInput(filename string) []line {
	f := utils.NewFile(filename)
	defer f.Close()

	var lines []line
	for f.Scan() {
		var (
			x1, y1 int
			x2, y2 int
		)
		_, _ = fmt.Sscanf(f.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		lines = append(lines, newLine(x1, y1, x2, y2))
	}

	return lines
}

func part1(lines []line) floor {
	var nonDiagonalOverlaps = make(floor)

	for _, l := range lines {
		if !l.isDiagonal() {
			for i := l.start.x; i <= l.finish.x; i++ {
				for j := l.start.y; j <= l.finish.y; j++ {
					nonDiagonalOverlaps.addPoint(i, j)
				}
			}
		}
	}

	return nonDiagonalOverlaps
}

func part2(lines []line) floor {
	var allOverlaps = make(floor)
	for _, l := range lines {
		if l.isDiagonal() {
			var di, dj = 1, 1
			if l.start.x > l.finish.x {
				di = -1
			}
			if l.start.y > l.finish.y {
				dj = -1
			}

			i := l.start.x
			j := l.start.y

			for {
				allOverlaps.addPoint(i, j)
				i += di
				j += dj

				if i == l.finish.x || j == l.finish.y {
					allOverlaps.addPoint(i, j)
					break
				}
			}
		} else {
			for i := l.start.x; i <= l.finish.x; i++ {
				for j := l.start.y; j <= l.finish.y; j++ {
					allOverlaps.addPoint(i, j)
				}
			}
		}
	}

	return allOverlaps
}

func main() {
	lines := parseInput(inputFile)

	counter := 0
	for _, val := range part1(lines) {
		if val >= 2 {
			counter += 1
		}
	}
	fmt.Println("Part 1:", counter)

	counter = 0
	for _, val := range part2(lines) {
		if val >= 2 {
			counter += 1
		}
	}
	fmt.Println("Part 2:", counter)
}
