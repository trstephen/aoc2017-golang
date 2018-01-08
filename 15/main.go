package main

import (
	"fmt"
	"math"
)

type generator struct {
	val, factor, nextMultiple int
}

func newGenerator(initVal, factor, nextMultiple int) *generator {
	return &generator{initVal, factor, nextMultiple}
}

func (g *generator) Next() {
	g.val = (g.val * g.factor) % math.MaxInt32

	if g.val%g.nextMultiple != 0 {
		g.Next()
	}
}

func main() {
	seedA := 618
	factorA := 16807
	multipleA := 4

	seedB := 814
	factorB := 48271
	multipleB := 8

	genA := newGenerator(seedA, factorA, multipleA)
	genB := newGenerator(seedB, factorB, multipleB)

	comparisonRounds := int(5e6)
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
