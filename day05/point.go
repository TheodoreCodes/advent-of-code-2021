package main

type point struct {
	x int
	y int
}

type floor map[point]int

func (recv floor) addPoint(x, y int) {
	p := point{x, y}

	if _, exists := recv[p]; exists {
		recv[p] += 1
	} else {
		recv[p] = 1
	}

}
