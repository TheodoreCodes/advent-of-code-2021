package main

import (
	"strconv"
	"strings"
)

const boardSize = 5

type boardState struct {
	board       [boardSize][boardSize]int
	matched     [boardSize][boardSize]bool
	lastMatched int
	steps       int
}

func (recv *boardState) Match(val int) bool {
	if recv.isBingo() {
		return true
	}

	recv.steps += 1

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if recv.board[i][j] == val {
				recv.matched[i][j] = true
				recv.lastMatched = recv.board[i][j]
				return true
			}
		}
	}

	return false
}

func (recv *boardState) isBingo() bool {
	var isFullRow bool
	// Check rows
	for i := 0; i < boardSize; i++ {
		isFullRow = true
		for j := 0; j < boardSize; j++ {
			if !recv.matched[i][j] {
				isFullRow = false
				break
			}
		}

		if isFullRow {
			return true
		}
	}
	// Check columns
	if !isFullRow {
		var isFullCol bool
		for j := 0; j < boardSize; j++ {
			isFullCol = true
			for i := 0; i < boardSize; i++ {
				if !recv.matched[i][j] {
					isFullCol = false
					break
				}
			}
			if isFullCol {
				return true
			}
		}
	}

	return false
}

func (recv *boardState) CalculateScore() int {
	var unmatchedSum int
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !recv.matched[i][j] {
				unmatchedSum += recv.board[i][j]
			}
		}
	}
	return recv.lastMatched * unmatchedSum
}

func NewBoardState(input string) *boardState {
	var nums [boardSize][boardSize]int

	rows := strings.Split(input, "\n")
	for i, row := range rows {
		row = strings.Replace(row, "  ", " ", -1)
		row = strings.TrimSpace(row)

		cells := strings.Split(row, " ")
		for j, cell := range cells {
			nums[i][j], _ = strconv.Atoi(cell)
		}
	}

	return &boardState{
		board: nums,
	}
}
