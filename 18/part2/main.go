package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type instruction interface {
	execute(p *program)
}

type memory map[string]int

type commChannel chan int

type program struct {
	instructions    []instruction
	memory          memory
	nextInstruction int
	state           string
	id              int
	comms           commChannel
	partnerComms    commChannel
	sendCommsCount  int
}

func main() {

	p0 := newProgram(0)
	p1 := newProgram(1)

	setPartners(p0, p1)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p0.addInstruction(scanner.Text())
		p1.addInstruction(scanner.Text())
	}

	done := new(sync.WaitGroup)

	done.Add(2)

	p0.start(done)
	p1.start(done)

	done.Wait()

	fmt.Println("All programs completed")
}

func newProgram(id int) *program {
	initMem := make(memory)
	initMem["p"] = id

	return &program{
		// slice is initialized by struct declaration
		memory:          initMem,
		nextInstruction: 0,
		state:           "running",
		id:              id,
		comms:           make(commChannel),
		sendCommsCount:  0,
	}
}

func setPartners(p1, p2 *program) {
	p1.partnerComms = p2.comms
	p2.partnerComms = p1.comms
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

func (p *program) start(done *sync.WaitGroup) {
	go func() {
		for p.state == "running" {
			p.executeNextInstruction()
		}

		fmt.Printf("[ %d ] Program finished with state %s\n", p.id, p.state)
		done.Done()
	}()
}

func (p *program) executeNextInstruction() {
	instr := p.instructions[p.nextInstruction]
	fmt.Printf("[ %d ] Executing: %s, %+v\n", p.id, reflect.TypeOf(instr), instr)
	instr.execute(p)
	fmt.Printf("[ %d ] Memory: %+v\n", p.id, p.memory)
	fmt.Printf("[ %d ] Next instr: %d/%d\n", p.id, p.nextInstruction, len(p.instructions)-1)
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

type snd struct{ val string }

func (s snd) execute(p *program) {
	// Prevent concurrent map access errors by reading register val
	// outside of these short-lived goroutines.
	val := p.getValue(s.val)
	go func() { p.partnerComms <- val }()

	p.sendCommsCount++
	fmt.Printf("[ %d ] send count: %d\n", p.id, p.sendCommsCount)

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
	p.memory[r.reg] = <-p.comms

	p.incrementInstructionCounter(1)
}

type jgz struct{ reg, val string }

func (j jgz) execute(p *program) {
	if p.getValue(j.reg) <= 0 {
		p.incrementInstructionCounter(1)
		return
	}

	p.incrementInstructionCounter(p.getValue(j.val))
}
