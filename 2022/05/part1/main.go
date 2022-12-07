package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	From    int
	To      int
	HowMany int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	stacks := getStacks(s)
	instructions := getInstructions(s)

	Do(stacks, instructions)

	for _, stack := range stacks {
		fmt.Print(stack.Top())
	}
	fmt.Println()
}

func getStacks(s *bufio.Scanner) []Stack {
	lines := []string{}
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		lines = append(lines, s.Text())
	}

	stacks := []Stack{}
	for i := 0; i < (len(lines[0])+1)/4; i++ {
		newStack := Stack{}
		for r := len(lines) - 2; r >= 0; r-- {
			c := string(lines[r][1+i*4])
			if c == " " {
				break
			}
			newStack.Push(c)
		}
		stacks = append(stacks, newStack)
	}
	return stacks
}

func getInstructions(s *bufio.Scanner) []Instruction {
	var instructions []Instruction
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		instructions = append(instructions, Instruction{
			From:    atoi(parts[3]),
			To:      atoi(parts[5]),
			HowMany: atoi(parts[1]),
		})
	}
	return instructions
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func Do(stacks []Stack, instructions []Instruction) {
	for _, instruction := range instructions {
		for instruction.HowMany > 0 {
			item := stacks[instruction.From-1].Pop()
			stacks[instruction.To-1].Push(item)
			instruction.HowMany--
		}
	}
}
