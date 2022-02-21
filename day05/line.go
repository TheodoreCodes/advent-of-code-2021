package main

type line struct {
	start  point
	finish point
}

func (recv *line) isDiagonal() bool {
	return recv.start.x != recv.finish.x &&
		recv.start.y != recv.finish.y
}

func newLine(startX, startY, endX, endY int) line {
	if startX > endX || startY > endY {
		startX, endX = endX, startX
		startY, endY = endY, startY
	}

	return line{
		start: point{
			x: startX,
			y: startY,
		},
		finish: point{
			x: endX,
			y: endY,
		},
	}
}
