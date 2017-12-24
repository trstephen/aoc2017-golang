package main

import (
	"bufio"
	"fmt"
	"os"
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
		if _, found := wordFreq[word]; !found {
			wordFreq[word] = 1
		} else {
			return 0
		}
	}

	return 1
}
