package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ParameterMode int

const (
	Position ParameterMode = iota
	Immediate
)

// IntCodeProcessor processes IntCode
type IntCodeProcessor struct {
	memory []int
}

func main() {
	icp := &IntCodeProcessor{}
	icp.Init(os.Stdin)

	if err := icp.Run(); err != nil {
		panic(err)
	}
}

func (icp *IntCodeProcessor) Init(in io.Reader) error {
	var s string
	fmt.Scanf("%s", &s)
	strProgram := strings.Split(s, ",")
	icp.memory = make([]int, len(strProgram))
	for i, item := range strProgram {
		var err error
		icp.memory[i], err = strconv.Atoi(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (icp *IntCodeProcessor) Clone() *IntCodeProcessor {
	memory := make([]int, len(icp.memory))
	copy(memory, icp.memory)
	return &IntCodeProcessor{
		memory: memory,
	}
}

func (icp *IntCodeProcessor) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Possibly of Bounds")
		}
	}()

	operations := map[int]func([]ParameterMode, *int){
		1: icp.Add,
		2: icp.Multiply,
		3: icp.Read,
		4: icp.Write,
		5: icp.JumpIfTrue,
		6: icp.JumpIfFalse,
		7: icp.Less,
		8: icp.Equal,
	}

	fmt.Println(icp)

	// alias the memory to keep lines short
	a := icp.memory
	ptr := 0
	for {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Printf("ptr at %d\n", ptr)
		instruction := a[ptr]
		opcode := instruction % 100
		if opcode == 99 {
			fmt.Println("END")
			return nil
		}
		modes := icp.ParseMode(instruction)

		operation, ok := operations[opcode]
		if !ok {
			return fmt.Errorf("Illegal instruction at location %d: %d", ptr, instruction)
		}

		operation(modes, &ptr)
		fmt.Println(icp)
	}
}

func (icp *IntCodeProcessor) Add(modes []ParameterMode, ptr *int) {
	a, b, dest := icp.ThreeVal(modes, ptr)
	addendA := icp.memory[a]
	addendB := icp.memory[b]
	icp.memory[dest] = addendA + addendB
	icp.printResultOf("+", a, addendA, b, addendB, dest)
}

func (icp *IntCodeProcessor) Multiply(modes []ParameterMode, ptr *int) {
	a, b, dest := icp.ThreeVal(modes, ptr)
	multiplicandA := icp.memory[a]
	multiplicandB := icp.memory[b]
	icp.memory[dest] = multiplicandA * multiplicandB
	icp.printResultOf("*", a, multiplicandA, b, multiplicandB, dest)
}

func (icp *IntCodeProcessor) Read(modes []ParameterMode, ptr *int) {
	dest := icp.OneVal(modes, ptr)
	fmt.Printf("Storing a[%d] >>> ", dest)
	fmt.Scanf("%d", &icp.memory[dest])
}

func (icp *IntCodeProcessor) Write(modes []ParameterMode, ptr *int) {
	dest := icp.OneVal(modes, ptr)
	fmt.Printf("Printing the value at a[%d]: ", dest)
	fmt.Printf("%d\n", icp.memory[dest])
}

func (icp *IntCodeProcessor) JumpIfTrue(modes []ParameterMode, ptr *int) {
	boolean, target := icp.TwoVal(modes, ptr)
	if icp.memory[boolean] != 0 {
		*ptr = icp.memory[target]
	}
}

func (icp *IntCodeProcessor) JumpIfFalse(modes []ParameterMode, ptr *int) {
	boolean, target := icp.TwoVal(modes, ptr)
	if icp.memory[boolean] == 0 {
		*ptr = icp.memory[target]
	}
}

func (icp *IntCodeProcessor) Less(modes []ParameterMode, ptr *int) {
	a, b, dest := icp.ThreeVal(modes, ptr)
	valA := icp.memory[a]
	valB := icp.memory[b]

	if icp.memory[a] < icp.memory[b] {
		icp.memory[dest] = 1
	} else {
		icp.memory[dest] = 0
	}

	icp.printResultOf("<", a, valA, b, valB, dest)
}

func (icp *IntCodeProcessor) Equal(modes []ParameterMode, ptr *int) {
	a, b, dest := icp.ThreeVal(modes, ptr)
	valA := icp.memory[a]
	valB := icp.memory[b]

	if icp.memory[a] == icp.memory[b] {
		icp.memory[dest] = 1
	} else {
		icp.memory[dest] = 0
	}

	icp.printResultOf("==", a, valA, b, valB, dest)
}

func (icp *IntCodeProcessor) printResultOf(op string, a, valA, b, valB, dest int) {
	fmtString := "a[%d](%d) = a[%d](%d) %s a[%d](%d)\n"
	fmt.Printf(fmtString, dest, icp.memory[dest], a, valA, op, b, valB)
}

func (icp *IntCodeProcessor) OneVal(modes []ParameterMode, ptr *int) (first int) {
	first = icp.Loc(modes[0], *ptr+1)
	*ptr += 2
	return
}

func (icp *IntCodeProcessor) TwoVal(modes []ParameterMode, ptr *int) (first, second int) {
	first = icp.Loc(modes[0], *ptr+1)
	second = icp.Loc(modes[1], *ptr+2)
	*ptr += 3
	return
}

func (icp *IntCodeProcessor) ThreeVal(modes []ParameterMode, ptr *int) (first, second, third int) {
	first = icp.Loc(modes[0], *ptr+1)
	second = icp.Loc(modes[1], *ptr+2)
	third = icp.Loc(modes[2], *ptr+3)
	*ptr += 4
	return
}

func (icp *IntCodeProcessor) Loc(mode ParameterMode, ptr int) int {
	switch mode {
	case Position:
		return icp.memory[ptr]
	case Immediate:
		return ptr
	default:
		return -1
	}
}

func (cp *IntCodeProcessor) ParseMode(instruction int) []ParameterMode {
	// Sheer off the opcode
	mode := instruction / 100
	modes := make([]ParameterMode, 3)
	for i := 0; i < 3; i++ {
		modes[i] = ParameterMode(mode % 10)
		mode /= 10
	}
	return modes
}

func (icp *IntCodeProcessor) String() string {
	s := strings.Builder{}
	s.WriteString("Memory:")
	lineno := 0
	var i, val int
	for i, val = range icp.memory {
		if i%8 == 0 {
			s.WriteString(fmt.Sprintf("\n[%-4d]", lineno))
			lineno += 8
		}
		s.WriteString(fmt.Sprintf("%8d", val))
	}
	for ; (i+1)%8 != 0; i++ {
		s.WriteString(fmt.Sprintf("%8s", "#"))
	}
	return s.String()
}
