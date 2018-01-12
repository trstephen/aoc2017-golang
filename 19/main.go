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
	state     string
}

func main() {
	var d diagram
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		d = append(d, scanner.Text())
	}

	f := d.newFollower()

	for f.walk() {
	}

	fmt.Println("Recovered string:", f.letters)
}

func (d diagram) newFollower() *diagramFollower {
	return &diagramFollower{
		diagram:   d,
		direction: "south",
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

func (df *diagramFollower) walk() bool {
	// some nice logic goes here~

	// for test purposes
	fmt.Printf("(%03d, %03d) %s\n", df.row, df.col, df.currentChar())
	return false
}
