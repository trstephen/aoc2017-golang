package main

import (
	"container/ring"
	"fmt"
	"log"
	"strings"
)

const (
	gridRows = 128
	gridCols = 128
)

var (
	// Knot has things
	sparseHash  *ring.Ring
	skip        int
	totalOffset int

	// region removal things
	grid [gridRows][gridCols]int
)

func main() {
	seed := "ffayrhll"
	regionCount := 0

	// 1. Populate a grid of []int where each row is a hash value
	for i := 0; i < gridRows; i++ {
		hash := computeDenseHash(fmt.Sprintf("%s-%d", seed, i))

		if len(hash) != gridCols {
			log.Panicf("Expected hash of length %d but got %d\n", gridCols, len(hash))
		}

		// Grid was initialized with all 0s so we just have to write 1s.
		for j, r := range hash {
			if r == '1' {
				grid[i][j] = 1
			}
		}
	}
	fmt.Println("Hash grid created")

	// 2. Start removing regions. Start top left and proceed through every element.
	// When a region member is found (i.e. a 1), break grid iteration and find all
	// elements of the region, overwriting them as 0 to prevent multi-counts.

	for i := 0; i < gridRows; i++ {
		for j := 0; j < gridCols; j++ {
			if grid[i][j] == 1 {
				regionCount++
				removeRegion(coord{i, j})
			}
		}
	}

	fmt.Println("regions detected:", regionCount)
}

func computeDenseHash(input string) string {
	ringSize := 256
	hashCycles := 64
	blockSize := 16
	skip = 0
	totalOffset = 0

	// 1. Convert input string to array of ASCII vals (0->255)
	var inputASCII []int
	for _, s := range input {
		inputASCII = append(inputASCII, int(s))
	}

	// 2. Add the common input suffix
	for _, s := range []int{17, 31, 73, 47, 23} {
		inputASCII = append(inputASCII, s)
	}

	// Populate our ring
	sparseHash = ring.New(ringSize)
	for i := 0; i < ringSize; i++ {
		// Note: ^ (XOR) only works on ints in golang so we'll use that to store
		// ring values for later hash densification.
		sparseHash.Value = i
		sparseHash = sparseHash.Next()
	}

	// 3. Do the hash 64 times, where the output of seq i is the input to i+1
	for i := 0; i < hashCycles; i++ {
		for _, move := range inputASCII {
			tieKnot(move)
		}
	}

	// Move the ring back to starting position.
	sparseHash = sparseHash.Move(-totalOffset)

	// 4. Make the dense hash by doing XOR on blocks of the sparse hash.
	denseHash := make([]int, ringSize/blockSize)
	for i := 0; i < len(denseHash); i++ {
		// Start iteration at 0x00 since 0x00 XOR 0xXX = 0xXX
		denseVal := 0

		for j := 0; j < blockSize; j++ {
			denseVal ^= sparseHash.Value.(int)
			sparseHash = sparseHash.Next()
		}

		denseHash[i] = denseVal
	}

	denseHashString := make([]string, len(denseHash))
	for i := 0; i < len(denseHashString); i++ {
		denseHashString[i] = fmt.Sprintf("%08b", denseHash[i])
	}

	return strings.Join(denseHashString, "")
}

func tieKnot(knotLength int) {
	// See Day 10 for a full explanation.

	// 0. Assume we're in the correct position to start the next knot

	// 1. create the []int to hold values, iterating forward through the ring
	knotValues := make([]int, knotLength)
	for i := 0; i < knotLength; i++ {
		knotValues[i] = sparseHash.Value.(int)
		sparseHash = sparseHash.Next()
	}

	// 2. write the values back into the ring in reverse order
	for i := 0; i < knotLength; i++ {
		sparseHash = sparseHash.Prev()
		sparseHash.Value = knotValues[i]
	}

	// 3. do the skip
	moveLength := knotLength + skip
	sparseHash = sparseHash.Move(moveLength)

	totalOffset += moveLength
	skip++
}

type coord struct {
	row, col int
}

func (c coord) North() coord {
	if c.col == 0 {
		return c
	}

	return coord{c.row, c.col - 1}
}

func (c coord) South() coord {
	if c.col == gridCols-1 {
		return c
	}

	return coord{c.row, c.col + 1}
}

func (c coord) West() coord {
	if c.row == 0 {
		return c
	}

	return coord{c.row - 1, c.col}
}

func (c coord) East() coord {
	if c.row == gridRows-1 {
		return c
	}

	return coord{c.row + 1, c.col}
}

func removeRegion(c coord) {
	// Recursive base case~
	if grid[c.row][c.col] == 0 {
		return
	}

	// Remove this coord from the grid and region
	grid[c.row][c.col] = 0

	// Explor neighbors... Recursively!
	removeRegion(c.North())
	removeRegion(c.South())
	removeRegion(c.West())
	removeRegion(c.East())
}
