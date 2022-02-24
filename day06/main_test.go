package main

import (
	"testing"
)

const (
	inputTestFile = "input_test.txt"
)

func TestRun(t *testing.T) {
	wantPart1 := 5934
	wantPart2 := 26984457539

	testState := parseInput(inputTestFile)
	gotPart1, gotPart2 := run(testState)

	if wantPart1 != gotPart1 {
		t.Errorf("Day 06 Part 1 error: Wanted %d, got %d", wantPart1, gotPart1)
	}

	if wantPart2 != gotPart2 {
		t.Errorf("Day 06 Part 2 error: Wanted %d, got %d", wantPart2, gotPart2)
	}
}
