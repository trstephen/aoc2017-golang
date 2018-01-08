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
	dg := newDanceGroup(size)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawMoves := strings.Split(scanner.Text(), ",")
		for _, rm := range rawMoves {
			parseDanceMove(rm).applyTo(dg)
		}
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
