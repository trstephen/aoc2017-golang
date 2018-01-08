package main

import (
	"fmt"
	"math"
)

type generator struct {
	val, factor int
}

func newGenerator(initVal, factor int) *generator {
	return &generator{initVal, factor}
}

func (g *generator) Next() {
	g.val = (g.val * g.factor) % math.MaxInt32
}

func main() {
	// seedA := 65 // demo
	seedA := 618
	factorA := 16807
	// seedB := 8921 // demo
	seedB := 814
	factorB := 48271

	genA := newGenerator(seedA, factorA)
	genB := newGenerator(seedB, factorB)

	comparisonRounds := int(40e6)
	pairCount := 0

	for i := 0; i < comparisonRounds; i++ {
		genA.Next()
		genB.Next()

		if judgeEqual(genA.val, genB.val) {
			pairCount++
		}
	}

	fmt.Println("Pairs:", pairCount)
}

func judgeEqual(a, b int) bool {
	mask := math.MaxUint16 // 0xFFFF

	return (a & mask) == (b & mask)
}
