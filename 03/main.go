package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{
		1,
		12,
		23,
		1024,
		265149,
	}

	for _, num := range nums {
		x, y := getSpiralCoord(num)
		fmt.Printf("%5d: (%4d, %4d) = %3f\n", num, x, y, math.Abs(float64(x))+math.Abs(float64(y)))
	}
}

func getSpiralCoord(end int) (int, int) {
	// Start at the origin and then proceed ccw until a new min/max
	// x,y value is reached to achieve a spiral progression.
	i := 1
	direction := "r"

	x := 0
	y := 0

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for i < end {
		// do the movement
		switch direction {
		case "r":
			x++
		case "u":
			y++
		case "l":
			x--
		case "d":
			y--
		}

		// change direction if at boundary
		switch {
		case direction == "r" && x > maxX:
			maxX = x
			direction = "u"
		case direction == "u" && y > maxY:
			maxY = y
			direction = "l"
		case direction == "l" && x < minX:
			minX = x
			direction = "d"
		case direction == "d" && y < minY:
			minY = y
			direction = "r"
		}

		i++
	}

	return x, y
}
