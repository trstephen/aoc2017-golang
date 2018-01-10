package main

import (
	"bufio"
	"fmt"
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
	lastPlayedFreq  int
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

}

func newProgram() *program {
	return &program{
		// slice is initialized by struct declaration
		memory:          make(memory),
		nextInstruction: 0,
		lastPlayedFreq:  -1,
		state:           "running",
	}
}

func (p *program) addInstruction(s string) {
	instructionType := s[:3]
	args := strings.Split(s[4:], " ")

	switch instructionType {
	case "snd":
		p.instructions = append(p.instructions, snd{args[0]})
	case "set":
		p.instructions = append(p.instructions, set{args[0], args[1]})
	case "add":
		p.instructions = append(p.instructions, add{args[0], args[1]})
	case "mul":
		p.instructions = append(p.instructions, mul{args[0], args[1]})
	case "mod":
		p.instructions = append(p.instructions, mod{args[0], args[1]})
	case "rcv":
		p.instructions = append(p.instructions, rcv{args[0]})
	case "jgz":
		p.instructions = append(p.instructions, jgz{args[0], args[1]})
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

func (p *program) getValue(reg string) int {
	// Maybe we were passed a literal value? Try to convert to verify.
	if val, err := strconv.Atoi(reg); err == nil {
		return val
	}

	// Not a literal, check memory for existing val.
	if val, found := p.memory[reg]; found {
		return val
	}

	// No val in memory, return default.
	return 0
}

type snd struct{ reg string }

func (s snd) execute(p *program) {
	p.lastPlayedFreq = p.getValue(s.reg)

	p.incrementInstructionCounter(1)
}

type set struct{ reg, val string }

func (s set) execute(p *program) {
	p.memory[s.reg] = p.getValue(s.val)

	p.incrementInstructionCounter(1)
}

type add struct{ reg, val string }

func (a add) execute(p *program) {
	p.memory[a.reg] = p.getValue(a.reg) + p.getValue(a.val)

	p.incrementInstructionCounter(1)
}

type mul struct{ reg, val string }

func (m mul) execute(p *program) {
	p.memory[m.reg] = p.getValue(m.reg) * p.getValue(m.val)

	p.incrementInstructionCounter(1)
}

type mod struct{ reg, val string }

func (m mod) execute(p *program) {
	p.memory[m.reg] = p.getValue(m.reg) % p.getValue(m.val)

	p.incrementInstructionCounter(1)
}

type rcv struct{ reg string }

func (r rcv) execute(p *program) {
	if p.getValue(r.reg) == 0 {
		p.incrementInstructionCounter(1)
		return
	}

	fmt.Println("Last played sound was:", p.lastPlayedFreq)

	p.state = "finished"
}

type jgz struct{ reg, val string }

func (j jgz) execute(p *program) {
	if p.getValue(j.reg) <= 0 {
		p.incrementInstructionCounter(1)
		return
	}

	p.incrementInstructionCounter(p.getValue(j.val))
}
