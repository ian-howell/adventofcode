package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	pc := 0
	instructions := readInstructions()

	for pc < len(instructions) {

		switch instructions[pc][0] {
		case "inc":
			registers[instructions[pc][1]]++
		case "dec":
			registers[instructions[pc][1]]--
		case "cpy":
			if strings.ContainsAny(instructions[pc][1], "abcd") {
				registers[instructions[pc][2]] = registers[instructions[pc][1]]
			} else {
				registers[instructions[pc][2]] = atoi(instructions[pc][1])
			}
		case "jnz":
			if strings.ContainsAny(instructions[pc][1], "abcd") {
				if registers[instructions[pc][1]] != 0 {
					pc += atoi(instructions[pc][2]) - 1
				}
			} else {
				if atoi(instructions[pc][1]) != 0 {
					pc += atoi(instructions[pc][2]) - 1
				}
			}
		}
		pc++
	}

	fmt.Println(registers["a"])
}

func readInstructions() [][]string {
	instructions := [][]string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		instruction := strings.Fields(s.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
