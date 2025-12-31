package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, operations := getInput()
	totals := nums[0]
	for _, row := range nums[1:] {
		for i, value := range row {
			totals[i] = operations[i](totals[i], value)
		}
	}
	total := totals[0]
	for _, value := range totals[1:] {
		total += value
	}
	fmt.Println(total)
}

type Operation func(a, b int) int

func add(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }

func getInput() ([][]int, []Operation) {
	nums := [][]int{}
	operations := []Operation{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ints, ok := toInts(s.Text())
		if ok {
			nums = append(nums, ints)
			continue
		}
		operations = toOperations(s.Text())
	}
	return nums, operations
}

func toInts(s string) (ints []int, ok bool) {
	if strings.ContainsAny(s, "+*") {
		return nil, false
	}
	for field := range strings.FieldsSeq(s) {
		ints = append(ints, atoi(field))
	}
	return ints, true
}

func atoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func toOperations(s string) []Operation {
	operations := []Operation{}
	for field := range strings.FieldsSeq(s) {
		switch field {
		case "+":
			operations = append(operations, add)
		case "*":
			operations = append(operations, multiply)
		}
	}
	return operations
}
