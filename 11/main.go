package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type position struct {
	x, y int
}

func main() {
	var moves []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		moves = strings.Split(scanner.Text(), ",")
	}

	var currentPosition position
	maxDistance := math.MinInt64
	origin := position{0, 0}

	// Let moves along the y-axis be length 2 and others 1.
	// Since N = NW + NE, or, (0,2) = (-1,1) + (1,1).
	// Similarly E = NE + SE, or, (2, 0) = (1,1) + (-1,1).

	for _, move := range moves {
		switch move {
		case "nw":
			currentPosition.x--
			currentPosition.y++
		case "n":
			currentPosition.y += 2
		case "ne":
			currentPosition.x++
			currentPosition.y++
		case "se":
			currentPosition.x++
			currentPosition.y--
		case "s":
			currentPosition.y -= 2
		case "sw":
			currentPosition.x--
			currentPosition.y--
		}

		currentDistance := currentPosition.distanceTo(origin)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}

	fmt.Println("Distance:", currentPosition.distanceTo(origin))
	fmt.Println("Max:", maxDistance)
}

func abs(val int) int {
	if val >= 0 {
		return val
	}

	return -val
}

func (p1 position) distanceTo(p2 position) int {
	dx := abs(p1.x - p2.x)
	dy := abs(p1.y - p2.y)

	if dx < dy {
		return (dx + dy) / 2
	}

	return dx
}
