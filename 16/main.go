package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	size := 16
	iterations := int(1e9)
	updateInterval := int(1e6)

	dg := newDanceGroup(size)
	var moves []danceMove

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawMoves := strings.Split(scanner.Text(), ",")
		for _, rm := range rawMoves {
			moves = append(moves, parseDanceMove(rm))
		}
	}

	iterationHistory := make(map[string]int)
	var seqHistory []string

	// Stick the base case in history to make duplicate detection in the
	// loop a little simpler.
	iterationHistory[dg.toString()] = 0
	seqHistory = append(seqHistory, dg.toString())

	for i := 1; i < iterations; i++ {
		for _, move := range moves {
			move.applyTo(dg)
		}

		// Log progress
		if i%(iterations/updateInterval) == 0 {
			log.Printf("%02.1f%%: %d\n", float64(i/updateInterval), i)
		}

		if lastSeen, found := iterationHistory[dg.toString()]; found {
			cycleLength := i - lastSeen
			log.Printf("Found cycle with length %d on iteration %d starting with '%s'\n", cycleLength, i, dg.toString())

			offset := iterations % cycleLength
			log.Printf("Sequence at %d iteration and offset %d will be '%s'\n", iterations, offset, seqHistory[offset])

			return
		}

		iterationHistory[dg.toString()] = i
		seqHistory = append(seqHistory, dg.toString())
	}

	fmt.Println(dg.toString())
}

type danceGroup map[string]int

func newDanceGroup(size int) danceGroup {
	dg := make(danceGroup)
	for i := 0; i < size; i++ {
		dg[string('a'+i)] = i
	}

	return dg
}

func (dg danceGroup) toString() string {
	orderedVals := make([]string, len(dg))

	for a, i := range dg {
		orderedVals[i] = a
	}

	return strings.Join(orderedVals, "")
}

func (dg danceGroup) invert() map[int]string {
	inv := make(map[int]string, len(dg))

	for a, i := range dg {
		inv[i] = a
	}

	return inv
}

func parseDanceMove(s string) danceMove {

	switch s[0] {
	case 's':
		size, _ := strconv.Atoi(s[1:])
		return spin{size}
	case 'x':
		indices := strings.Split(s[1:], "/")
		i1, _ := strconv.Atoi(indices[0])
		i2, _ := strconv.Atoi(indices[1])
		return exchange{i1, i2}
	case 'p':
		partners := strings.Split(s[1:], "/")
		return partner{partners[0], partners[1]}
	default:
		log.Panicln("Unrecognized command:", s)
	}

	return nil
}

type danceMove interface {
	applyTo(dg danceGroup)
}

type spin struct {
	size int
}

func (s spin) applyTo(dg danceGroup) {
	for a, i := range dg {
		dg[a] = (i + s.size) % len(dg)
	}
}

type exchange struct {
	i1, i2 int
}

func (x exchange) applyTo(dg danceGroup) {
	inv := dg.invert()

	a1 := inv[x.i1]
	a2 := inv[x.i2]

	dg[a1] = x.i2
	dg[a2] = x.i1
}

type partner struct {
	a1, a2 string
}

func (p partner) applyTo(dg danceGroup) {
	temp := dg[p.a1]

	dg[p.a1] = dg[p.a2]
	dg[p.a2] = temp
}
