package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	validCount := 0

	for scanner.Scan() {
		validCount += validatePassword(scanner.Text())
	}

	fmt.Println("Valid passwords:", validCount)
}

func validatePassword(s string) int {
	words := strings.Split(s, " ")

	wordFreq := make(map[string]int)

	for _, word := range words {
		// For part 2 but still works with part 1
		// e.g. adec -> acde
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sortedWord := strings.Join(letters, "")

		if _, found := wordFreq[sortedWord]; !found {
			wordFreq[sortedWord] = 1
		} else {
			return 0
		}
	}

	return 1
}
