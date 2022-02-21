package main

import (
	"testing"
)

const (
	inputTestFile = "input_test.txt"
)

func TestPart1(t *testing.T) {
	want := 5

	lines := parseInput(inputTestFile)
	got := 0
	for _, val := range part1(lines) {
		if val >= 2 {
			got += 1
		}
	}

	if want != got {
		t.Errorf("Day 05 Part 1 error: Wanted %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 12

	lines := parseInput(inputTestFile)
	got := 0
	for _, val := range part2(lines) {
		if val >= 2 {
			got += 1
		}
	}

	if want != got {
		t.Errorf("Day 05 Part 2 error: Wanted %d, got %d", want, got)
	}
}
