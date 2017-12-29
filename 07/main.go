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
	unallocatedNodes := make(map[string]towerNode)

	for scanner.Scan() {
		node := parseToTowerNode(scanner.Text())
		// fmt.Println(node)
		unallocatedNodes[node.name] = node
	}

	// build tree with `azqje` as root
	var rootName string
	smallRoot := "tknk"
	bigRoot := "azqje"
	if _, found := unallocatedNodes[smallRoot]; found {
		rootName = smallRoot
	} else {
		rootName = bigRoot
	}

	// do the recursive magic dealy
	tower := buildTower(unallocatedNodes, rootName)

	printTower(tower)

}

func (cn childNodes) contains(val string) bool {
	for _, n := range cn {
		if n.name == val {
			return true
		}
	}

	return false
}

type childNodes []towerNode

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

	childNames := strings.Split(matches[0][4], ", ")
	var children childNodes
	if childNames[0] != "" {
		for _, name := range childNames {
			children = append(children, towerNode{name, 0, childNodes{}})
		}
	}

	return towerNode{
		name:     matches[0][1],
		weight:   weight,
		children: children,
	}
}

func buildTower(unallocatedNodes map[string]towerNode, root string) towerNode {
	tower := unallocatedNodes[root]
	populatedChildren := make(childNodes, len(tower.children))

	// Add actual child nodes
	for i, child := range tower.children {

		child.weight = unallocatedNodes[child.name].weight

		for _, grandchild := range unallocatedNodes[child.name].children {
			child.children = append(child.children, buildTower(unallocatedNodes, grandchild.name))
		}
		// fmt.Println("Child:", child)
		populatedChildren[i] = child
	}

	tower.children = populatedChildren

	// fmt.Println("tower:", tower)

	return tower
}

func printTower(t towerNode) {
	formattedTower := printFormatTower(t, 0)
	formattedTower = strings.Replace(formattedTower, "\n\n", "\n", -1)

	fmt.Println(formattedTower)
}

func printFormatTower(t towerNode, offset int) string {
	var nodeString string

	nodeString += fmt.Sprintf("%s%d (%d):\n", strings.Repeat(" ", offset), t.fullWeight(), t.weight)
	for _, child := range t.children {
		nodeString += fmt.Sprintf("%s%s\n", strings.Repeat(" ", offset), printFormatTower(child, offset+2))
	}

	return nodeString
}

func (tn towerNode) fullWeight() int {
	weight := tn.weight

	for _, cn := range tn.children {
		weight += cn.fullWeight()
	}

	return weight
}
