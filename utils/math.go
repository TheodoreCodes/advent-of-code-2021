package utils

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Max(numbers ...int) int {
	max := math.MinInt

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

func Min(numbers ...int) int {
	min := math.MaxInt

	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}
