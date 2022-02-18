package main

import (
	"fmt"
	"github.com/theodorecodes/advent-of-code-2021/utils"
	"strconv"
	"strings"
	"sync"
)

const (
	inputFile = "day04/input.txt"
)

func parseInput(filename string) ([]int, []*boardState) {
	f := utils.NewFile(filename)

	input := strings.Split(f.Read(), "\n\n")
	f.Close()

	// Extract drawn numbers
	var nums []int
	for _, s := range strings.Split(input[0], ",") {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	var boards []*boardState
	for i := 1; i < len(input); i++ {
		boards = append(boards, NewBoardState(input[i]))
	}
	return nums, boards
}

func solve(wg *sync.WaitGroup, b *boardState, drawnNums []int) {
	defer wg.Done()
	for _, num := range drawnNums {
		isMatch := b.Match(num)
		if isMatch && b.isBingo() {
			return
		}
	}
}

func play(boards []*boardState, drawnNumbers []int) (winnerBoard *boardState, loserBoard *boardState) {
	var wg sync.WaitGroup
	for _, b := range boards {
		wg.Add(1)
		go solve(&wg, b, drawnNumbers)
	}
	wg.Wait()

	winnerBoard = boards[0]
	loserBoard = boards[0]
	for _, b := range boards {
		if b.steps < winnerBoard.steps {
			winnerBoard = b
		}

		if b.steps > loserBoard.steps {
			loserBoard = b
		}
	}
	return
}

func main() {
	drawnNums, boards := parseInput(inputFile)

	winnerBoard, loserBoard := play(boards, drawnNums)

	fmt.Println("Part 1:", winnerBoard.CalculateScore())
	fmt.Println("Part 2:", loserBoard.CalculateScore())
}
