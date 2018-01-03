package main

import (
	"container/ring"
	"fmt"
)

var (
	// Using the ring built-in will handle iteration overflow easily
	sparseHash  *ring.Ring
	skip        = 0
	totalOffset = 0
)

func main() {
	ringSize := 256
	hashCycles := 64
	blockSize := 16
	input := "18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188"

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

	fmt.Println("input:", input)
	fmt.Printf("hash: ")
	for _, val := range denseHash {
		// Be sure to print leading 0s.
		// Using 3 instead of 03 will lead to incorrect hashes!
		fmt.Printf("%02x", val)
	}
	fmt.Println()

}

func printRing(r *ring.Ring) {
	i := 0
	var value string
	r.Do(func(node interface{}) {
		if node == nil {
			value = "nil"
		} else {
			value = fmt.Sprintf("%d", node.(byte))
		}
		fmt.Printf("%d: %s\n", i, value)
		i++
	})
}

func tieKnot(knotLength int) {
	// Let's avoid using the ring built-ins because there's no simple way to
	// reverse a ring. Reversing would involve allocation of a temp ring and
	// then reading the values in backwards from the unlinked ring. It's probably
	// easier to read the values backwards into an []int and then overwrite
	// the original ring.

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
