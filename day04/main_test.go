package main

import "testing"

const (
	inputTestFile = "input_test.txt"
)

func TestPart1(t *testing.T) {
	wantWinnerScore := 4512
	wantLastMatched := 24

	drawnNums, boards := parseInput(inputTestFile)
	winnerBoard, _ := play(boards, drawnNums)
	gotWinnerScore := winnerBoard.CalculateScore()

	if gotWinnerScore != wantWinnerScore {
		t.Errorf("Day 04 Part 1 error: Wanted score %d, got %d", wantWinnerScore, gotWinnerScore)
	}
	if winnerBoard.lastMatched != wantLastMatched {
		t.Errorf("Day 04 Part 1 error: Wanted last matched %d, got %d", wantLastMatched, winnerBoard.lastMatched)
	}
}

func TestPart2(t *testing.T) {
	wantLoserScore := 1924
	wantLastMatched := 13

	drawnNums, boards := parseInput(inputTestFile)
	_, loserBoard := play(boards, drawnNums)
	gotLoserScore := loserBoard.CalculateScore()

	if gotLoserScore != wantLoserScore {
		t.Errorf("Day 04 Part 2 error: Wanted score %d, got %d", wantLoserScore, gotLoserScore)
	}

	if loserBoard.lastMatched != wantLastMatched {
		t.Errorf("Day 04 Part 2 error: Wanted last matched %d, got %d", wantLastMatched, loserBoard.lastMatched)

	}
}
