package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type firewall map[int]int

func main() {

	// Store the firewall as a map to make miss detection easy.
	// Alternative would be padding out an array or something so lol.
	fw := make(firewall)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		depth, _ := strconv.Atoi(parts[0])
		scanRange, _ := strconv.Atoi(parts[1])

		fw[depth] = scanRange
	}

	noDelaySeverity, _ := fw.findTraversalSeverityWithDelay(0)
	fmt.Println("Severity with no delay:", noDelaySeverity)
	fmt.Println("Min delay:", fw.findMinimumDelayForSafeTraversal())
}

// We only care when the scanner is in the "top" position since that's
// the path we traverse. For various ranges, the top position will be
// occupied on iterations:
//   2 -> 0, 2, 4, 6 (2)
//   3 -> 0, 4, 8, 12 (4)
//   4 -> 0, 6, 12, 18 (6)
//   5 -> 0, 8, 16, 24 (8)
// The cycle length is (r-1)*2. So, the top spot is occupied when
// `i % ((r-1)*2) == 0`
func (fw *firewall) findTraversalSeverityWithDelay(delay int) (int, int) {
	maxDepth := math.MinInt64
	for depth := range *fw {
		if depth > maxDepth {
			maxDepth = depth
		}
	}

	totalSeverity := 0
	timesCaught := 0

	for i := 0; i <= maxDepth; i++ {
		if scanRange, found := (*fw)[i]; found {
			scanPosition := (i + delay) % ((scanRange - 1) * 2)
			if scanPosition == 0 {
				// fmt.Printf("Found! (%02d,%02d)\n", i, scanRange)
				timesCaught++
				totalSeverity += i * scanRange
			}
		}
	}

	return totalSeverity, timesCaught
}

// iono, just try until it works. something something p = np
func (fw *firewall) findMinimumDelayForSafeTraversal() int {
	delay := 0

	for {
		if _, caught := fw.findTraversalSeverityWithDelay(delay); caught == 0 {
			break
		}

		delay++
	}

	return delay
}
