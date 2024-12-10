package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "#") {
			continue
		}
		total += process(s.Text())
	}
	fmt.Println(total)
}

func process(s string) int {
	target, inputs := parse(s)
	if valid(target, inputs) {
		return target
	}
	return 0
}

func valid(target int, inputs []int) bool {
	if len(inputs) == 0 {
		return false
	}

	if len(inputs) == 1 && inputs[0] == target {
		return true
	}

	l := len(inputs) - 1
	if target%inputs[l] == 0 && valid(target/inputs[l], inputs[:l]) {
		return true
	}

	if valid(target-inputs[l], inputs[:l]) {
		return true
	}

	trimmed := trimSuffix(target, inputs[l])
	if trimmed != target && valid(trimmed, inputs[:l]) {
		return true
	}

	return false
}

func trimSuffix(num, suffix int) int {
	numStr := strconv.Itoa(num)
	suffixStr := strconv.Itoa(suffix)
	return atoi(strings.TrimSuffix(numStr, suffixStr))
}

func parse(s string) (int, []int) {
	targetStr, after, _ := strings.Cut(s, ": ")
	inputsStr := strings.Fields(after)
	return atoi(targetStr), toInts(inputsStr)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a) // errors do not occur in my AOC inputs
	return i
}

func toInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		ints = append(ints, atoi(s))
	}
	return ints
}
