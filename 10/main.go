package main

import (
	"container/ring"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	// Using the ring built-in will handle iteration overflow easily
	hash        *ring.Ring
	skip        = 0
	totalOffset = 0
)

func main() {
	// ringSize := 5
	// input := "3,4,1,5"
	ringSize := 256
	input := "18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188"

	// Populate our ring
	hash = ring.New(ringSize)
	for i := 0; i < ringSize; i++ {
		hash.Value = i
		hash = hash.Next()
	}

	for _, move := range strings.Split(input, ",") {
		tieKnot(strconv.Atoi(move))
	}

	// Move the ring back to starting position
	hash = hash.Move(-totalOffset)

	printRing(hash)

}

func printRing(r *ring.Ring) {
	i := 0
	var value string
	r.Do(func(node interface{}) {
		if node == nil {
			value = "nil"
		} else {
			value = fmt.Sprintf("%d", node.(int))
		}
		fmt.Printf("%d: %s\n", i, value)
		i++
	})
}

func tieKnot(knotLength int, err error) {
	if err != nil {
		log.Panic("Oops lol")
	}
	// Let's avoid using the ring built-ins because there's no simple way to
	// reverse a ring. Reversing would involve allocation of a temp ring and
	// then reading the values in backwards from the unlinked ring. It's probably
	// easier to read the values backwards into an int[] and then overwrite
	// the original ring.

	// 0. Assume we're in the correct position to start the next knot

	// 1. create the int[] to hold values, iterating forward through the ring
	knotValues := make([]int, knotLength)
	for i := 0; i < knotLength; i++ {
		knotValues[i] = hash.Value.(int)
		hash = hash.Next()
	}

	// fmt.Println("Before mod")
	// fmt.Println("Knot:", knotValues)
	// printRing(hash)
	// fmt.Println("-----")

	// 2. write the values back into the ring in reverse order
	for i := 0; i < knotLength; i++ {
		hash = hash.Prev()
		hash.Value = knotValues[i]
		// printRing(hash)
	}

	// fmt.Println("-----")

	// 3. do the skip
	moveLength := knotLength + skip
	hash = hash.Move(moveLength)

	totalOffset += moveLength
	skip++
}
