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
		checksum += getWholeDivisor(scanner.Text())
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

func getWholeDivisor(s string) int {
	fmt.Println(s)

	// We know there is exactly one pair of numbers what will divide without remainder.
	// Do the n^2 approach until something more fancy comes along~

	// Convert to ints now to make the double loop a bit nicer to deal with.
	cells := strings.Split(s, "\t")
	nums := make([]int, len(cells))
	for i, cell := range cells {
		nums[i], _ = strconv.Atoi(cell)
	}

	for i, x := range nums {
		for _, y := range nums[i+1 : len(nums)] {
			if x < y && y%x == 0 {
				return y / x
			}

			if y < x && x%y == 0 {
				return x / y
			}
		}
	}

	// the bad case :(
	return -1
}
