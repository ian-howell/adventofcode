package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ReadProgram(r io.Reader) []int {
	var rawInput string
	fmt.Fscan(r, &rawInput)
	splitInput := strings.Split(rawInput, ",")
	program := make([]int, 0, len(splitInput))
	for _, s := range splitInput {
		program = append(program, atoi(s))
	}
	return program
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
