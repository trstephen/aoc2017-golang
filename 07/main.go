package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	unsortedNodes := make(map[string]towerNode)

	for scanner.Scan() {
		node := parseToTowerNode(scanner.Text())
		// fmt.Println(node)
		unsortedNodes[node.name] = node
	}

	fmt.Println("Total nodes:", len(unsortedNodes))

	// We can be lazy and ignore the whole jazz with weights.
	// Identify the leaves as those with no children.
	// Then identify those with leaves as children, progressing level
	// by level until we get to the root.

	leaves := make(map[string]struct{})

	for key, node := range unsortedNodes {
		if node.children == nil {
			leaves[node.name] = struct{}{}
			delete(unsortedNodes, key)
		}
	}

	// fmt.Println("leaves:", leaves)
	fmt.Println("Num leaves:", len(leaves))

	nextLevelNodes := make(map[string]struct{})
	lastLevelNodes := leaves
	level := 1

	for len(unsortedNodes) > 1 {
		fmt.Println("Processing nodes at level", level)
		fmt.Println("Remaining:", len(unsortedNodes))

		for nodeName := range lastLevelNodes {
			for key, node := range unsortedNodes {
				if node.children.contains(nodeName) {
					// fmt.Println("Assigning", node)
					nextLevelNodes[key] = struct{}{}
					delete(unsortedNodes, key)
				}
			}
		}

		level++
		lastLevelNodes = nextLevelNodes
		nextLevelNodes = make(map[string]struct{})
	}

	fmt.Println("Root is", unsortedNodes)
}

func (cn childNodes) contains(val string) bool {
	for _, n := range cn {
		if n == val {
			return true
		}
	}

	return false
}

type childNodes []string

type towerNode struct {
	name     string
	weight   int
	children childNodes
}

var nodeExpr = regexp.MustCompile(`^(\w+) \(([\d]+)\)( -> )*([\w, ]*)*`)

func parseToTowerNode(s string) towerNode {
	matches := nodeExpr.FindAllStringSubmatch(s, -1)
	weight, _ := strconv.Atoi(matches[0][2])

	// fmt.Println(s)

	children := strings.Split(matches[0][4], ", ")
	if children[0] == "" {
		children = nil
	}

	return towerNode{
		name:     matches[0][1],
		weight:   weight,
		children: children,
	}
}
