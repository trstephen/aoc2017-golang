package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// input := "0	2	7	0"
	input := "4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3"

	fmt.Println("input: ", input)
	fmt.Println("cycles:", cyclesBeforeRepeat(input))
}

func cyclesBeforeRepeat(input string) int {
	cycles := 0
	history := make(map[string]int)

	// Convert input to []int
	var arr []int
	for _, s := range strings.Split(input, "\t") {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}

	for {
		arr = balance(arr)
		cycles++

		// Check if we've seen this distribution before
		key := makeKey(arr)
		if _, found := history[key]; found {
			fmt.Println("Duplicate:", key)
			break
		} else {
			history[key] = 1
		}
	}

	return cycles
}

func balance(a []int) []int {
	// Maybe don't need to do this copy here? w/e
	balanced := a

	maxPosn := 0
	maxVal := math.MinInt64

	// Find position and value of largest element.
	for i, val := range a {
		if val > maxVal {
			maxPosn = i
			maxVal = val
		}
	}

	// Do round-robin distribution of largest value starting at next element.
	balanced[maxPosn] = 0
	for maxVal > 0 {
		maxPosn++
		i := maxPosn % len(balanced)
		balanced[i] = balanced[i] + 1
		maxVal--
	}

	return balanced
}

func makeKey(a []int) string {
	keyParts := make([]string, len(a))
	for i, num := range a {
		keyParts[i] = strconv.Itoa(num)
	}

	return strings.Join(keyParts, "-")
}
