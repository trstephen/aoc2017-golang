package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	checksum := 0

	for scanner.Scan() {
		checksum += getRowDifference(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("Checksum is", checksum)
}

func getRowDifference(s string) int {
	fmt.Println(s)

	max := math.MinInt64
	min := math.MaxInt64

	for _, a := range strings.Split(s, "\t") {
		val, _ := strconv.Atoi(a)
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}

	return max - min
}
