package main

import (
	"fmt"
	"strings"
)

type Computer struct {
	A int
	B int

	PC           int
	instructions []Instruction
}

func (c *Computer) Reset() {
	c.A = 0
	c.B = 0
	c.PC = 0
	c.instructions = nil
}

func (c *Computer) LoadProgram(instructions []string) {
	for _, instruction := range instructions {
		instruction = normalize(instruction)
		parts := strings.Fields(instruction)
		switch parts[0] {
		case "hlf":
			c.instructions = append(c.instructions, newHlf(parts))
		case "tpl":
			c.instructions = append(c.instructions, newTpl(parts))
		case "inc":
			c.instructions = append(c.instructions, newInc(parts))
		case "jmp":
			c.instructions = append(c.instructions, newJmp(parts))
		case "jie":
			c.instructions = append(c.instructions, newJie(parts))
		case "jio":
			c.instructions = append(c.instructions, newJio(parts))
		}
	}
}

func (c *Computer) Finished() bool {
	return c.PC >= len(c.instructions)
}

func (c *Computer) Run(debug bool) {
	for !c.Finished() {
		if debug {
			c.PrintStatus()
		}
		c.Step()
	}
}

func (c *Computer) PrintStatus() {
	fmt.Println("========================================")
	fmt.Printf("A: %v\n", c.A)
	fmt.Printf("B: %v\n", c.B)
	fmt.Printf("PC: %v\n", c.PC)
	fmt.Printf("next: %+v\n", c.instructions[c.PC])
}

func (c *Computer) PrintInstructions() {
	for i, instruction := range c.instructions {
		fmt.Printf("%v: %+v\n", i, instruction)
	}
}

func (c *Computer) Step() {
	switch instruction := c.instructions[c.PC]; instruction.Op {
	case HlfOp:
		c.Hlf(instruction)
	case TplOp:
		c.Tpl(instruction)
	case IncOp:
		c.Inc(instruction)
	case JmpOp:
		c.Jmp(instruction)
	case JieOp:
		c.Jie(instruction)
	case JioOp:
		c.Jio(instruction)
	}
	c.PC++
}

func (c *Computer) Hlf(instruction Instruction) {
	switch instruction.Reg {
	case "A":
		c.A /= 2
	case "B":
		c.B /= 2
	}
}

func (c *Computer) Tpl(instruction Instruction) {
	switch instruction.Reg {
	case "A":
		c.A *= 3
	case "B":
		c.B *= 3
	}
}

func (c *Computer) Inc(instruction Instruction) {
	switch instruction.Reg {
	case "A":
		c.A++
	case "B":
		c.B++
	}
}

func (c *Computer) Jmp(instruction Instruction) {
	c.PC += instruction.Offset - 1 // -1 because the PC always advances
}

func (c *Computer) Jie(instruction Instruction) {
	if (instruction.Reg == "A" && c.A%2 == 0) || (instruction.Reg == "B" && c.B%2 == 0) {
		c.PC += instruction.Offset - 1 // -1 because the PC always advances
	}
}

func (c *Computer) Jio(instruction Instruction) {
	if (instruction.Reg == "A" && c.A == 1) || (instruction.Reg == "B" && c.B == 1) {
		c.PC += instruction.Offset - 1 // -1 because the PC always advances
	}
}
