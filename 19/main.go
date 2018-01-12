package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type diagram []string

type diagramFollower struct {
	diagram   diagram
	direction string
	row, col  int
	letters   string
}

func main() {
	var d diagram
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// moderately cheating since I know the grid has a width of 200
		padding := strings.Repeat(" ", 200-len(scanner.Text()))
		d = append(d, scanner.Text()+padding)
	}

	f := d.newFollower()

	for f.walk() {
	}

	fmt.Println("Recovered string:", f.letters)
}

func (d diagram) newFollower() *diagramFollower {
	return &diagramFollower{
		diagram:   d,
		direction: DOWN,
		row:       0,
		col:       strings.Index(d[0], "|"),
		letters:   "",
	}
}

func (d diagram) charAt(row, col int) string {
	return string(d[row][col])
}

func (df *diagramFollower) currentChar() string {
	return df.diagram.charAt(df.row, df.col)
}

// gimme enums already
const (
	DOWN  = "DOWN"
	UP    = "UP"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

func (df *diagramFollower) turn() {
	switch df.direction {
	case DOWN, UP:
		switch df.col {
		case 0:
			df.direction = RIGHT
		case len(df.diagram[df.row]) - 1:
			df.direction = LEFT
		default:
			if df.diagram.charAt(df.row, df.col-1) == " " {
				df.direction = RIGHT
			} else {
				df.direction = LEFT
			}
		}
	case LEFT, RIGHT:
		switch df.row {
		case 0:
			df.direction = DOWN
		case len(df.diagram) - 1:
			df.direction = UP
		default:
			if df.diagram.charAt(df.row-1, df.col) == " " {
				df.direction = DOWN
			} else {
				df.direction = UP
			}
		}
	}
}

func (df *diagramFollower) walk() bool {
	// fmt.Printf("(%03d, %03d) %s, %s\n", df.row, df.col, df.currentChar(), df.direction)

	switch df.direction {
	case DOWN:
		df.row++
	case UP:
		df.row--
	case LEFT:
		df.col--
	case RIGHT:
		df.col++
	}

	switch df.currentChar() {
	case "+":
		df.turn()
	case "-", "|":
		break
	case " ":
		// walked off the edge!
		return false
	default: // letters
		df.letters += df.currentChar()
	}

	return true
}
