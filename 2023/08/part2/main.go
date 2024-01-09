package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dirMap = map[byte]int{
	'L': 0,
	'R': 1,
}

type State struct {
	Position string
	Index    int
}

func main() {
	instructions, graph := getInput()
	getNextState := createNext(instructions, graph)

	cycleLengths := map[string]int{}
	for pos := range graph {
		if strings.HasSuffix(pos, "A") {
			state := State{
				Position: pos,
				Index:    0,
			}
			cycleLengths[pos] = getPathLen(state, getNextState)
		}
	}

	final := 1
	for _, v := range cycleLengths {
		final = leastCommonMultiple(final, v)
	}
	fmt.Println(final)
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

func createNext(instructions string, graph map[string][2]string) func(State) State {
	return func(state State) State {
		newState := State{
			Position: graph[state.Position][dirMap[instructions[state.Index]]],
			Index:    (state.Index + 1) % len(instructions),
		}
		return newState
	}
}

func getPathLen(state State, getNextState func(State) State) int {
	step := 1
	nextState := getNextState(state)
	for {
		if strings.HasSuffix(nextState.Position, "Z") {
			return step
		}
		nextState = getNextState(nextState)
		step++
	}
}

func leastCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return greatestCommonDivisor(b, a%b)
}
