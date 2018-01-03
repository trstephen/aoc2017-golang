package main

import (
	"bufio"
	"fmt"
	"os"
)

type mergePair struct {
	partner, target string
}

var (
	position = map[string]int{
		"nw": 0,
		"n":  0,
		"ne": 0,
		"se": 0,
		"s":  0,
		"sw": 0,
	}

	mergePartners = map[string][2]mergePair{
		"nw": [2]mergePair{
			mergePair{partner: "s", target: "sw"},
			mergePair{partner: "ne", target: "n"},
		},
		"n": [2]mergePair{
			mergePair{partner: "sw", target: "nw"},
			mergePair{partner: "se", target: "ne"},
		},
		"ne": [2]mergePair{
			mergePair{partner: "nw", target: "n"},
			mergePair{partner: "s", target: "se"},
		},
		"se": [2]mergePair{
			mergePair{partner: "n", target: "ne"},
			mergePair{partner: "sw", target: "s"},
		},
		"s": [2]mergePair{
			mergePair{partner: "ne", target: "se"},
			mergePair{partner: "nw", target: "sw"},
		},
		"sw": [2]mergePair{
			mergePair{partner: "se", target: "s"},
			mergePair{partner: "n", target: "nw"},
		},
	}

	opposites = map[string]string{
		"nw": "se",
		"n":  "s",
		"ne": "sw",
		"se": "nw",
		"s":  "n",
		"sw": "ne",
	}
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		updatePosition(scanner.Text())
	}

	distance := 0
	for _, val := range position {
		distance += val
	}

	fmt.Println(position)
	fmt.Println("distance:", distance)
}

func updatePosition(direction string) {
	lastChangedDir := direction

	for _, merge := range mergePartners[direction] {
		if position[merge.partner] > 0 {
			// Do the merge by consuming the partner. If the new direction
			// doesn't have a cancel partner it will be written.
			position[merge.partner]--
			lastChangedDir = merge.target
			break
		}
	}

	oppositeDir := opposites[lastChangedDir]
	if position[oppositeDir] > 0 {
		position[oppositeDir]--
		return
	}

	position[lastChangedDir]++
}
