package main

import (
	"container/ring"
	"fmt"
	"strings"
)

func main() {
	seed := "ffayrhll"
	usedSquareCount := 0

	// To figure out how many ones are in each row we'll use a lookup
	// for the hex digits of the hash.

	onesCount := map[rune]int{
		'0': 0, // 0x0000
		'1': 1, // 0x0001
		'2': 1, // 0x0010
		'3': 2, // 0x0011
		'4': 1, // 0x0100
		'5': 2, // 0x0101
		'6': 2, // 0x0110
		'7': 3, // 0x0111
		'8': 1, // 0x1000
		'9': 2, // 0x1001
		'a': 2, // 0x1010
		'b': 3, // 0x1011
		'c': 2, // 0x1100
		'd': 3, // 0x1101
		'e': 3, // 0x1110
		'f': 4, // 0x1111
	}

	for i := 0; i < 128; i++ {
		hash := computeDenseHash(fmt.Sprintf("%s-%d", seed, i))
		for _, r := range hash {
			usedSquareCount += onesCount[r]
		}
	}

	fmt.Println("Used squares:", usedSquareCount)
}

var (
	sparseHash  *ring.Ring
	skip        int
	totalOffset int
)

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
		denseHashString[i] = fmt.Sprintf("%02x", denseHash[i])
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
