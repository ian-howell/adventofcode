package main

import (
	"bufio"
	"fmt"
	"log"
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

	indexOfLast := len(inputs) - 1
	if target%inputs[indexOfLast] == 0 && valid(target/inputs[indexOfLast], inputs[:indexOfLast]) {
		return true
	}

	return valid(target-inputs[indexOfLast], inputs[:indexOfLast])
}

func parse(s string) (int, []int) {
	targetStr, after, _ := strings.Cut(s, ": ")
	inputsStr := strings.Fields(after)
	return atoi(targetStr), toInts(inputsStr)
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("failed to convert '%s' to int: %v", a, err)
	}
	return i
}

func toInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		ints = append(ints, atoi(s))
	}
	return ints
}
