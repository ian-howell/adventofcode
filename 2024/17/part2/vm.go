package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation int

func (o Operation) String() string {
	return map[int]string{
		0: "adv", 1: "bxl", 2: "bst", 3: "jnz",
		4: "bxc", 5: "out", 6: "bdv", 7: "cdv",
	}[int(o)]
}

type VM struct {
	a, b, c int
	pc      int
	steps   []int
	results []int
}

func (vm *VM) Run() {
	debugln("INITIAL STATE: ")
	debugln(vm.diagnostics())
	for step, done := vm.next(); !done; step, done = vm.next() {
		vm.Do(step)
		debugln(vm.diagnostics())
	}
}

func (vm *VM) next() (step int, done bool) {
	vm.pc++
	if vm.pc >= len(vm.steps) {
		return 0, true
	}

	return vm.steps[vm.pc], false
}

func (vm *VM) Do(step int) {
	map[int]func(){
		0: vm.adv, 1: vm.bxl, 2: vm.bst, 3: vm.jnz,
		4: vm.bxc, 5: vm.out, 6: vm.bdv, 7: vm.cdv,
	}[step]()
}

func (vm *VM) adv() {
	operand, _ := vm.next()
	divisor := 1 << vm.combo(operand)
	vm.a /= divisor
}

func (vm *VM) bxl() {
	operand, _ := vm.next()
	vm.b ^= operand
}

func (vm *VM) bst() {
	operand, _ := vm.next()
	vm.b = vm.combo(operand) & 0x07
}

func (vm *VM) jnz() {
	operand, _ := vm.next()
	if vm.a != 0 {
		vm.pc = operand - 1
	}
}

func (vm *VM) bxc() {
	_, _ = vm.next()
	vm.b ^= vm.c
}

func (vm *VM) out() {
	operand, _ := vm.next()
	vm.results = append(vm.results, vm.combo(operand)&0x07)
}

func (vm *VM) bdv() {
	operand, _ := vm.next()
	divisor := 1 << vm.combo(operand)
	vm.b = vm.a / divisor
}

func (vm *VM) cdv() {
	operand, _ := vm.next()
	divisor := 1 << vm.combo(operand)
	vm.c = vm.a / divisor
}

func (vm *VM) combo(x int) int {
	return map[int]int{
		0: 0, 1: 1, 2: 2, 3: 3,
		4: vm.a, 5: vm.b, 6: vm.c, 7: 7,
	}[x]
}

func (vm *VM) Results() string {
	return strings.Join(toStr(vm.results), ",")
}

func (vm *VM) diagnostics() string {
	sb := strings.Builder{}
	sb.WriteString("==== DIAGNOSTICS ====\n")
	sb.WriteString(fmt.Sprintf("A: %-20b B: %-20b C: %-20b\n", vm.a, vm.b, vm.c))
	// sb.WriteString("\n")
	// for i, step := range vm.steps {
	// 	if i == vm.pc+1 {
	// 		sb.WriteString(red(fmt.Sprintf("%5d", step)))
	// 	} else {
	// 		sb.WriteString(fmt.Sprintf("%5d", step))
	// 	}
	// }
	// sb.WriteString("\n")
	if vm.pc < len(vm.steps)-1 {
		sb.WriteString("NEXT STEP: ")
		sb.WriteString(yellow(underline(fmt.Sprintf("%v %d\n",
			Operation(vm.steps[vm.pc+1]), vm.steps[vm.pc+2]))))
	}
	sb.WriteString(fmt.Sprintf("RESULTS: %v\n", vm.results))
	sb.WriteString("---------------------\n")
	return sb.String()
}

func getInput() VM {
	vm := VM{pc: -1}
	var progStr string

	// Look what they (the linters) did to my boy...
	_, _ = fmt.Scanf("Register A: %d", &vm.a)
	_, _ = fmt.Scanf("Register B: %d", &vm.b)
	_, _ = fmt.Scanf("Register C: %d", &vm.c)
	_, _ = fmt.Scanln()
	_, _ = fmt.Scanf("Program: %s", &progStr)

	vm.steps = toInts(progStr)
	return vm
}

func toInts(s string) []int {
	parts := strings.Split(s, ",")
	xs := make([]int, 0, len(parts))
	for _, part := range parts {
		i, _ := strconv.Atoi(part)
		xs = append(xs, i)
	}
	return xs
}

func toStr(xs []int) []string {
	ss := make([]string, 0, len(xs))
	for _, x := range xs {
		ss = append(ss, strconv.Itoa(x))
	}
	return ss
}
