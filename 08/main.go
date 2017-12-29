package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type registers map[string]int

type instruction struct {
	targetReg           string
	targetOp            string
	targetAmt           int
	conditionReg        string
	conditionComparison string
	conditionAmt        int
}

var (
	memory           = make(registers)
	instrExpr        = regexp.MustCompile(`^([\w]+) (inc|dec) (\-?[\d]+) if ([\w]+) (==|!=|>|>=|<|<=) (\-?[\d]+)$`)
	highestMemoryVal = math.MinInt64
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		applyInstruction(parseToInstruction(scanner.Text()))
	}

	largestReg := memory.findLargest()
	fmt.Printf("Largest: %s, %d\n", largestReg, memory[largestReg])

	fmt.Println("High water mark for memory value:", highestMemoryVal)
}

func parseToInstruction(s string) instruction {
	matches := instrExpr.FindAllStringSubmatch(s, -1)
	match := matches[0]

	targetAmt, _ := strconv.Atoi(match[3])
	conditionAmt, _ := strconv.Atoi(match[6])

	return instruction{
		targetReg:           match[1],
		targetOp:            match[2],
		targetAmt:           targetAmt,
		conditionReg:        match[4],
		conditionComparison: match[5],
		conditionAmt:        conditionAmt,
	}
}

func applyInstruction(in instruction) {
	var targetVal, conditionVal int
	var found bool
	if targetVal, found = memory[in.targetReg]; !found {
		targetVal = 0
	}
	if conditionVal, found = memory[in.conditionReg]; !found {
		conditionVal = 0
	}

	var conditionSatisfied bool
	switch in.conditionComparison {
	case "==":
		conditionSatisfied = (conditionVal == in.conditionAmt)
	case "!=":
		conditionSatisfied = (conditionVal != in.conditionAmt)
	case "<":
		conditionSatisfied = (conditionVal < in.conditionAmt)
	case "<=":
		conditionSatisfied = (conditionVal <= in.conditionAmt)
	case ">":
		conditionSatisfied = (conditionVal > in.conditionAmt)
	case ">=":
		conditionSatisfied = (conditionVal >= in.conditionAmt)
	}

	var newVal int
	if conditionSatisfied {
		switch in.targetOp {
		case "inc":
			newVal = targetVal + in.targetAmt
		case "dec":
			newVal = targetVal - in.targetAmt
		}

		memory[in.targetReg] = newVal
		if newVal > highestMemoryVal {
			highestMemoryVal = newVal
			fmt.Println(in.targetReg, highestMemoryVal)
		}
	}
}

func (r registers) findLargest() string {
	var largestReg string
	maxVal := math.MinInt64

	for reg, val := range r {
		if val > maxVal {
			maxVal = val
			largestReg = reg
		}
	}

	return largestReg
}
