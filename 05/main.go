package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var instructions []int

	for scanner.Scan() {
		instruction, _ := strconv.Atoi(scanner.Text())
		instructions = append(instructions, instruction)
	}

	fmt.Println("Steps to escape:", countStepsToEscape(instructions))
}

func countStepsToEscape(instructions []int) int {
	posn := 0
	stepCount := 0

	for {
		stepCount++
		nextPosn := posn + instructions[posn]

		var instructionChange int
		if instructions[posn] >= 3 {
			instructionChange = -1
		} else {
			instructionChange = 1
		}
		instructions[posn] = instructions[posn] + instructionChange

		if nextPosn >= len(instructions) || nextPosn < 0 {
			break
		}

		posn = nextPosn
	}

	return stepCount
}
