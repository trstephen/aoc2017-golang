package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{
		800,
		265149,
	}

	for _, num := range nums {
		x, y := getSpiralcoord(num)
		fmt.Printf("%5d: (%4d, %4d) = %3f\n", num, x, y, math.Abs(float64(x))+math.Abs(float64(y)))
	}

	for _, num := range nums {
		fmt.Println(num, getFirstSpiralNumLargerThan(num))
	}
}

func getSpiralcoord(end int) (int, int) {
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

type coord struct {
	x, y int
}
type spiral map[coord]int

func getFirstSpiralNumLargerThan(target int) int {
	spiral := make(map[coord]int)
	spiral[coord{0, 0}] = 1

	offsets := []coord{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	direction := "u"

	x := 1
	y := 0

	minX := 0
	maxX := 1
	minY := 0
	maxY := 0

	val := 0

	for val < target {
		// calculate value from neighbors
		val = 0
		for _, o := range offsets {
			neighbor := coord{
				x: x + o.x,
				y: y + o.y,
			}

			if nval, ok := spiral[neighbor]; ok {
				val += nval
			} else {
				val += 0
			}
		}

		spiral[coord{x, y}] = val
		// fmt.Println(val)

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

	}

	return val
}
