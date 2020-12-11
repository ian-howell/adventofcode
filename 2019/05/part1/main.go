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
	icpBase := &IntCodeProcessor{}
	icpBase.Init(os.Stdin)

	icp := icpBase.Clone()
	err := icp.Run()
	if err != nil {
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

	fmt.Println(icp)

	// alias the memory to keep lines short
	a := icp.memory
	ptr := 0
	for {
		fmt.Printf("ptr at %d\n", ptr)
		instruction := a[ptr]
		opcode := instruction % 100
		modes := icp.ParseMode(instruction)
		switch opcode {
		case 1:
			icp.Add(modes, &ptr)
			fmt.Println(icp)
		case 2:
			icp.Multiply(modes, &ptr)
			fmt.Println(icp)
		case 3:
			icp.Read(modes, &ptr)
			fmt.Println(icp)
		case 4:
			icp.Write(modes, &ptr)
		case 99:
			fmt.Println("END")
			return nil
		default:
			return fmt.Errorf("Illegal instruction at location %d: %d", ptr, instruction)
		}
	}
}

func (icp *IntCodeProcessor) Add(modes []ParameterMode, ptr *int) {
	a, b, dest := icp.ThreeVal(modes, *ptr)
	addendA := icp.memory[a]
	addendB := icp.memory[b]
	icp.memory[dest] = addendA + addendB
	fmt.Printf("a[%d](%d) = a[%d](%d) + a[%d](%d)\n",
		dest, icp.memory[dest],
		a, addendA,
		b, addendB)
	*ptr += 4
}

func (icp *IntCodeProcessor) Multiply(modes []ParameterMode, ptr *int) {
	a, b, dest := icp.ThreeVal(modes, *ptr)
	multiplicandA := icp.memory[a]
	multiplicandB := icp.memory[b]
	icp.memory[dest] = multiplicandA * multiplicandB
	fmt.Printf("a[%d](%d) = a[%d](%d) * a[%d](%d)\n",
		dest, icp.memory[dest],
		a, multiplicandA,
		b, multiplicandB)
	*ptr += 4
}

func (icp *IntCodeProcessor) Read(modes []ParameterMode, ptr *int) {
	dest := icp.OneVal(modes, *ptr)
	fmt.Printf("Storing a value at %d\n", dest)
	fmt.Scanf("%d", &icp.memory[dest])
	*ptr += 2
}

func (icp *IntCodeProcessor) Write(modes []ParameterMode, ptr *int) {
	dest := icp.OneVal(modes, *ptr)
	fmt.Printf("Printing the value at %d\n", dest)
	fmt.Printf("%d\n", icp.memory[dest])
	*ptr += 2
}

func (icp *IntCodeProcessor) OneVal(modes []ParameterMode, ptr int) (first int) {
	return icp.Loc(modes[0], ptr+1)
}

func (icp *IntCodeProcessor) TwoVal(modes []ParameterMode, ptr int) (first, second int) {
	return icp.OneVal(modes, ptr), icp.Loc(modes[1], ptr+2)
}

func (icp *IntCodeProcessor) ThreeVal(modes []ParameterMode, ptr int) (first, second, third int) {
	first, second = icp.TwoVal(modes, ptr)
	return first, second, icp.Loc(modes[2], ptr+3)
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
