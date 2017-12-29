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

func scoreGroups(s string) int {
	score := 0

	// the magic!

	return score
}
