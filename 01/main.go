package main

import (
	"fmt"
	"strconv"
)

func main() {
	captchas := []string{
		"1212",
		"123123",
		"1221",
		"123425",
		"12131415",
	}

	fmt.Println("  Captcha  ->   Solved")
	fmt.Println("==========================")

	for _, c := range captchas {
		fmt.Printf("%10s -> %7s\n", c, solveCaptchaPart2(c))
	}
}

func solveCaptchaPart1(s string) string {
	var prev, curr int
	acc := 0

	switch len(s) {
	case 0:
		// do nothing!
	case 1:
		i, _ := strconv.Atoi(s)
		acc += i
	default:
		wrapped := fmt.Sprintf("%s%c", s, s[0])

		for _, ch := range wrapped {
			// int() does conversion to ASCII value, then remove offest
			curr = (int(ch) - '0')
			// println("curr: ", curr)
			if curr == prev {
				acc += curr
			}
			prev = curr
		}
	}

	return strconv.Itoa(acc)
}

func solveCaptchaPart2(s string) string {
	// Assume no funny business with extremely long or short captchas.
	// Also, all captures are guaranteed to be even length. lol.

	acc := 0
	lookahead := len(s) / 2

	for i, ch := range s {
		if byte(ch) == s[(i+lookahead)%len(s)] {
			acc += (int(ch) - '0')
		}
	}

	return strconv.Itoa(acc)
}
