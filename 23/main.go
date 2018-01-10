package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type instruction interface {
	execute(p *program)
}

type memory map[string]int

type program struct {
	instructions    []instruction
	memory          memory
	nextInstruction int
	mulCount        int
	state           string
}

func main() {

	p := newProgram()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.addInstruction(scanner.Text())
	}

	for p.state == "running" {
		p.executeNextInstruction()
	}

	fmt.Println("Program stopped with state:", p.state)
	fmt.Println("mul count:", p.mulCount)

	// part 2
	b := 84*100 + 100000
	c := b + 17000
	h := 0

	for i := b; i < c+1; i += 17 {
		if !isPrime(i) {
			h++
		}
	}

	fmt.Println("h ->", h)
}

func newProgram() *program {
	return &program{
		// slice is initialized by struct declaration
		memory:          make(memory),
		nextInstruction: 0,
		mulCount:        0,
		state:           "running",
	}
}

func (p *program) addInstruction(s string) {
	instructionType := s[:3]
	args := strings.Split(s[4:], " ")

	switch instructionType {
	case "set":
		p.instructions = append(p.instructions, set{args[0], args[1]})
	case "add":
		p.instructions = append(p.instructions, add{args[0], args[1]})
	case "sub":
		p.instructions = append(p.instructions, sub{args[0], args[1]})
	case "mul":
		p.instructions = append(p.instructions, mul{args[0], args[1]})
	case "jnz":
		p.instructions = append(p.instructions, jnz{args[0], args[1]})
	default:
		fmt.Println("Unrecognized instruction:", s)
		p.state = "error"
	}
}

func (p *program) executeNextInstruction() {
	instr := p.instructions[p.nextInstruction]
	// fmt.Printf("Executing: %s, %+v\n", reflect.TypeOf(instr), instr)
	instr.execute(p)
	// fmt.Println("Memory:", p.memory)
	// fmt.Printf("Next instr: %d/%d\n", p.nextInstruction, len(p.instructions)-1)
	// fmt.Println()
}

func (p *program) incrementInstructionCounter(n int) {
	p.nextInstruction += n

	if p.nextInstruction < 0 || p.nextInstruction >= len(p.instructions) {
		fmt.Println("Jumped to out of bound instruction:", p.nextInstruction)
		p.state = "error"
	}
}

func (m memory) getValue(reg string) int {
	// Maybe we were passed a literal value? Try to convert to verify.
	if val, err := strconv.Atoi(reg); err == nil {
		return val
	}

	// Not a literal, check memory for existing val.
	if val, found := m[reg]; found {
		return val
	}

	// No val in memory, return default.
	return 0
}

type set struct{ reg, val string }

func (s set) execute(p *program) {
	p.memory[s.reg] = p.memory.getValue(s.val)

	p.incrementInstructionCounter(1)
}

type add struct{ reg, val string }

func (a add) execute(p *program) {
	p.memory[a.reg] = p.memory.getValue(a.reg) + p.memory.getValue(a.val)

	p.incrementInstructionCounter(1)
}

type sub struct{ reg, val string }

func (s sub) execute(p *program) {
	p.memory[s.reg] = p.memory.getValue(s.reg) - p.memory.getValue(s.val)

	p.incrementInstructionCounter(1)
}

type mul struct{ reg, val string }

func (m mul) execute(p *program) {
	p.memory[m.reg] = p.memory.getValue(m.reg) * p.memory.getValue(m.val)

	p.mulCount++

	p.incrementInstructionCounter(1)
}

type mod struct{ reg, val string }

func (m mod) execute(p *program) {
	p.memory[m.reg] = p.memory.getValue(m.reg) % p.memory.getValue(m.val)

	p.incrementInstructionCounter(1)
}

type jnz struct{ reg, val string }

func (j jnz) execute(p *program) {
	if p.memory.getValue(j.reg) == 0 {
		p.incrementInstructionCounter(1)
		return
	}

	p.incrementInstructionCounter(p.memory.getValue(j.val))
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return n > 1
}
