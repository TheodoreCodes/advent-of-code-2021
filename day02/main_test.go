package main

import "testing"

const inputTestFile = "input_test.txt"

func TestPart1(t *testing.T) {
	wantX, wantY := 15, 10
	gotX, gotY := part1(inputTestFile)

	if wantX != gotX {
		t.Errorf("Day 02 Part 1 error: Wanted horizontal position %d, got %d", wantX, gotX)
	}

	if wantY != gotY {
		t.Errorf("Day 02 Part 1 error: Wanted depth %d, got %d", wantY, gotY)
	}
}

func TestPart2(t *testing.T) {
	wantX, wantY := 15, 60
	gotX, gotY := part2(inputTestFile)

	if wantX != gotX {
		t.Errorf("Day 02 Part 2 error: Wanted horizontal position %d, got %d", wantX, gotX)
	}

	if wantY != gotY {
		t.Errorf("Day 02 Part 2 error: Wanted depth %d, got %d", wantY, gotY)
	}
}
