package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	instructions, graph := getInput()

	dirMap := map[byte]int{
		'L': 0,
		'R': 1,
	}

	totalSteps := 0
	currPos := "AAA"
	for step := 0; currPos != "ZZZ"; step = ((step + 1) % len(instructions)) {
		currPos = graph[currPos][dirMap[instructions[step]]]
		totalSteps++
	}
	fmt.Println(totalSteps)
}

func getInput() (string, map[string][2]string) {
	s := bufio.NewScanner(os.Stdin)

	// Get the first line of input
	s.Scan()
	instructions := s.Text()

	// Eat the empty line
	s.Scan()

	graph := map[string][2]string{}
	for s.Scan() {
		from, to, _ := strings.Cut(s.Text(), " = ")

		to = strings.Trim(to, "()")
		to1, to2, _ := strings.Cut(to, ", ")
		graph[from] = [2]string{to1, to2}
	}

	return instructions, graph
}
