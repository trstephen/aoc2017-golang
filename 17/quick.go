package main

import "log"

func main() {
	shiftRange := 303
	target := int(50e6)

	i := 1
	currentPosn := 0
	ringLength := 1
	lastInsertAfterZero := -1

	for i < target {
		currentPosn = (currentPosn + shiftRange + 1) % ringLength

		if currentPosn == 0 {
			lastInsertAfterZero = i
			log.Println("Inserting after 0:", lastInsertAfterZero)
		}

		ringLength++
		i++
	}

	log.Println("Last insert after 0 was:", lastInsertAfterZero)
}
