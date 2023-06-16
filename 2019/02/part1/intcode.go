package main

import "fmt"

const (
	Acc = iota + 2
	NumRegisters
)

const Halt = -2

type IntCode struct {
	Debug     bool
	Sp        int
	Memory    []int
	Registers []int
}

func NewIntCode(opts ...IntCodeOption) *IntCode {
	ic := &IntCode{
		Registers: make([]int, NumRegisters),
	}

	for _, opt := range opts {
		opt(ic)
	}

	return ic
}

func (ic *IntCode) Load(instructions []int) {
	ic.Memory = instructions
}

func (ic *IntCode) Run() (exitCode int, err error) {
	for ic.Running() {
		ic.DebugMsg()
		if err = ic.Step(); err != nil {
			return -1, err
		}
	}
	return ic.Memory[0], nil
}

func (ic *IntCode) Step() (err error) {
	switch ic.Memory[ic.Sp] {
	case RefAddOpCode:
		err = ic.RefAdd()
	case RefMulOpCode:
		err = ic.RefMul()
	case HaltOpCode:
		ic.Halt()
	}
	ic.Sp++
	return
}

func (ic *IntCode) Direct(i int) *int {
	return &ic.Memory[i]
}

func (ic *IntCode) Indirect(i int) *int {
	return &ic.Memory[ic.Memory[i]]
}

func (ic *IntCode) RefAdd() (err error) {
	if ic.Sp+3 >= len(ic.Memory) {
		return fmt.Errorf("OutOfBounds Error: accessing element %d of array with length %d", ic.Sp+3, len(ic.Memory))
	}
	ic.Sp++
	ic.Registers[Acc] = *ic.Indirect(ic.Sp)
	ic.Sp++
	ic.Registers[Acc] += *ic.Indirect(ic.Sp)
	ic.Sp++
	*ic.Indirect(ic.Sp) = ic.Registers[Acc]
	return
}

func (ic *IntCode) RefMul() (err error) {
	if ic.Sp+3 >= len(ic.Memory) {
		return fmt.Errorf("OutOfBounds Error: accessing element %d of array with length %d", ic.Sp+3, len(ic.Memory))
	}
	ic.Sp++
	ic.Registers[Acc] = *ic.Indirect(ic.Sp)
	ic.Sp++
	ic.Registers[Acc] *= *ic.Indirect(ic.Sp)
	ic.Sp++
	*ic.Indirect(ic.Sp) = ic.Registers[Acc]
	return
}

func (ic *IntCode) Halt() {
	ic.Sp = Halt
}

func (ic *IntCode) Running() bool {
	return ic.Sp >= 0
}

func (ic *IntCode) DebugMsg() {
	if !ic.Debug {
		return
	}

	fmt.Printf("SP:           %d\n", ic.Sp)
	fmt.Printf("ACC:          %d\n", ic.Registers[Acc])
	fmt.Printf("Instructions: %v\n", ic.Memory)
	// fmt.Print("   ")
	// for i := ic.Sp; i < len(ic.Instructions) && i < (ic.Sp+10); i++ {
	// 	fmt.Printf(" %d", ic.Instructions[i])
	// }
	fmt.Println()
}
