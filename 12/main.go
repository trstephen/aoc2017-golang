package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 1. Make a big map of connections, (program -> connections)
	unassignedPrograms := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " <-> ")

		program := parts[0]
		connections := strings.Split(parts[1], ", ")

		unassignedPrograms[program] = connections
	}

	// 2. Starting at 0, remove the element from the map of connections
	// and then repeat this for each of the connections for the removed elements

	connectedPrograms := 0

	programQueue := make(chan string, len(unassignedPrograms))
	programQueue <- "0"

	for p := range programQueue {
		// Check if the program is unprocessed
		if connections, found := unassignedPrograms[p]; found {
			// Add connections for processing and remove from the map.
			connectedPrograms++
			delete(unassignedPrograms, p)
			for _, c := range connections {
				// Don't reprocess connections we've removed
				if _, found := unassignedPrograms[c]; found {
					programQueue <- c
				}
			}
		}

		// Terminate the for loop when the channel is empty
		if len(programQueue) == 0 {
			close(programQueue)
		}
	}

	fmt.Println("Connected programs:", connectedPrograms)
}
