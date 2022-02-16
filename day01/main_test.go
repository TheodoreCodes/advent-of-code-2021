package main

import "testing"

const (
	inputTestFile = "input_test.txt"
)

func TestPart1(t *testing.T) {
	want := 7
	got := part1(inputTestFile)

	if want != got {
		t.Errorf("Day 01 Part 1 wrong: Wanted %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 5
	got := part2(inputTestFile)

	if want != got {
		t.Errorf("Day 01 Part 2 wrong: Wanted %d, got %d", want, got)
	}
}
