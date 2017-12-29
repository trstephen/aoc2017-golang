package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Group score:", scoreGroups(scanner.Text()))
		fmt.Println("Garbage chars:", countGarbageChars(scanner.Text()))
	}
}

type parserState struct {
	prev, current, next string
}

// State enum, lol thanks Go for the lack of cool enum suppor
const (
	GARBAGE = "GARBAGE"
	CANCEL  = "CANCEL"
	OKAY    = "OKAY"
)

func scoreGroups(s string) int {
	var state parserState
	state.current = OKAY
	state.next = OKAY

	score := 0
	groupDepth := 0

	for _, ch := range s {
		token := fmt.Sprintf("%c", ch)

		switch state.current {
		case OKAY:
			switch token {
			case "{":
				groupDepth++
			case "}":
				score += groupDepth
				groupDepth--
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

	return score
}

func countGarbageChars(s string) int {
	var state parserState
	state.current = OKAY
	state.next = OKAY

	garbageCharCount := 0

	for _, ch := range s {
		token := fmt.Sprintf("%c", ch)

		switch state.current {
		case OKAY:
			switch token {
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
			default:
				garbageCharCount++
			}
		case CANCEL:
			state.next = state.prev
		}

		state.prev = state.current
		state.current = state.next
	}

	return garbageCharCount
}
