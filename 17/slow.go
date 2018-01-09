package main

import (
	"container/ring"
	"fmt"
	"log"
)

// lol, this ran for >30min without terminating so I've gotta use the closed
// form. Ring creation only took ~13 min, though.
// And the only reason there's a closed form is because we're looking
// for 0 which happens to be in all cases. Coding questions that rely on fragile
// tricks (read: cleverness) instead of robust design can get fucked.
func slow() {
	shiftRange := 303
	insertions := int(50e6)
	seekVal := 0

	sl := newSpinlock(shiftRange)

	log.Printf("Starting %8d insertions\n", insertions)
	for i := 1; i <= insertions; i++ {
		sl.addValue(i)

		if i%(insertions/20) == 0 {
			log.Printf("Finished %8d insertions (%02d %%)\n", i, (i*100)/insertions)
		}
	}

	log.Println("Seeking", seekVal)
	sl.seek(seekVal)

	fmt.Printf("Value after %d is %d\n", seekVal, sl.r.Next().Value)
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
	// Suppose this leaves ring in state: a (b) c
	sl.r = sl.r.Move(sl.shiftRange)

	// Make new ring: (z)
	r2 := ring.New(1)
	r2.Value = val

	// Ring is now: a z (b) c
	sl.r = sl.r.Link(r2)

	// Move so new element is active: a (z) b c
	sl.r = sl.r.Prev()
}

func (sl *spinlock) seek(val int) {
	for i := 0; i < sl.r.Len(); i++ {
		if rval := sl.r.Value.(int); rval == val {
			return
		}
		sl.r = sl.r.Next()
	}

	// Didn't find value D:
	log.Panicf("Couldn't find %d in the ring\n", val)
}

func printRing(r *ring.Ring) {
	r.Do(func(node interface{}) {
		if node != nil {
			fmt.Printf("%d ", node.(int))
		}
	})

	fmt.Println()
}
