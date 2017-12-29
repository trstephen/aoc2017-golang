package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scoreGroups(scanner.Text()))
	}
}

type parserState struct {
	prev, current, next string
	score, groupDepth   int
}

func scoreGroups(s string) int {
	const (
		GARBAGE = "GARBAGE"
		CANCEL  = "CANCEL"
		OKAY    = "OKAY"
	)

	var state parserState
	state.current = OKAY
	state.next = OKAY

	for _, ch := range s {
		token := fmt.Sprintf("%c", ch)

		switch state.current {
		case OKAY:
			switch token {
			case "{":
				state.groupDepth++
			case "}":
				state.score += state.groupDepth
				state.groupDepth--
			case "<":
				state.next = GARBAGE
			case "!":
				state.next = CANCEL
			}
		case GARBAGE:
			switch token {
			case "!":
				state.next = CANCEL
			case ">":
				state.next = OKAY
			}
		case CANCEL:
			state.next = state.prev
		}

		state.prev = state.current
		state.current = state.next
	}

	return state.score
}
