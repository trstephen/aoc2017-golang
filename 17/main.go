package main

import (
	"container/ring"
	"fmt"
)

func main() {
	shiftRange := 303
	insertions := 2017

	sl := newSpinlock(shiftRange)

	for i := 1; i <= insertions; i++ {
		sl.addValue(i)
	}

	fmt.Printf("Element after %d is %d\n", insertions, sl.r.Next().Value)
}

type spinlock struct {
	r          *ring.Ring
	shiftRange int
}

func newSpinlock(shiftRange int) *spinlock {
	initRing := ring.New(1)
	initRing.Value = 0

	return &spinlock{initRing, shiftRange}
}

func (sl *spinlock) addValue(val int) {
	for i := 0; i < sl.shiftRange; i++ {
		sl.r = sl.r.Next()
	}

	// Suppose this leaves ring in state: a (b) c

	// Make new ring: (z)
	r2 := ring.New(1)
	r2.Value = val

	// Ring is now: a z (b) c
	sl.r = sl.r.Link(r2)

	// Move so new element is active: a (z) b c
	sl.r = sl.r.Prev()
}

func printRing(r *ring.Ring) {
	r.Do(func(node interface{}) {
		if node != nil {
			fmt.Printf("%d ", node.(int))
		}
	})

	fmt.Println()
}
