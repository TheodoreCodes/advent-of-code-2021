package main

import "testing"

const (
	inputTestFile = "input_test.txt"
	inputWidth    = 5
)

func TestPart1(t *testing.T) {
	var wantGamma, wantEpsilon uint = 22, 9
	gotGamma, gotEpsilon := part1(inputTestFile, inputWidth)

	if wantGamma != gotGamma {
		t.Errorf("Day 03 Part 1 error: Wanted gamma %d, got %d", wantGamma, gotGamma)
	}

	if wantEpsilon != gotEpsilon {
		t.Errorf("Day 03 Part 1 error: Wanted epsilon %d, got %d", wantEpsilon, gotEpsilon)
	}
}

func TestPart2(t *testing.T) {
	var wantO2, wantCO2 uint = 23, 10
	gotO2, gotCO2 := part2(inputTestFile, inputWidth)

	if wantO2 != gotO2 {
		t.Errorf("Day 03 Part 2 error: Wanted o2 %d, got %d", wantO2, gotO2)
	}

	if wantCO2 != gotCO2 {
		t.Errorf("Day 03 Part 2 error: Wanted co2 %d, got %d", wantCO2, gotCO2)
	}
}
