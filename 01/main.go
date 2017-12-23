package main

import (
	"fmt"
	"strconv"
)

func main() {
	captchas := []string{
		"1122",
		"1111",
		"1234",
		"918181819",
	}

	fmt.Println("  Captcha  ->   Solved")
	fmt.Println("==========================")

	for _, c := range captchas {
		fmt.Printf("%10s -> %10s\n", c, solveCaptcha(c))
	}
}

func solveCaptcha(s string) string {
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
