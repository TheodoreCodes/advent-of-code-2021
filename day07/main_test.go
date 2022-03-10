package main

import "testing"

const (
	inputTestFile = "input_test.txt"
)

func TestRun(t *testing.T) {
	wantPart1 := 37
	wantPart2 := 168

	testPositions := parseInput(inputTestFile)
	gotPart1, gotPart2 := run(testPositions)

	if wantPart1 != gotPart1 {
		t.Errorf("Day 07 Part 1 error: Wanted %d, got %d", wantPart1, gotPart1)
	}

	if wantPart2 != gotPart2 {
		t.Errorf("Day 07 Part 2 error: Wanted %d, got %d", wantPart2, gotPart2)
	}
}
