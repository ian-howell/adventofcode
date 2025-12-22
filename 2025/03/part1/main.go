package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := getInput()
	fmt.Println(calcTotalJoltage(input))
}

func calcTotalJoltage(inputs []string) int {
	total := 0
	for _, input := range inputs {
		joltage := calcJoltage(input)
		fmt.Printf("input: %-15s joltage: %d\n", input, joltage)
		total += joltage
	}
	return total
}

func calcJoltage(input string) int {
	for first := '9'; first > '0'; first-- {
		for i, c := range input[:len(input)-1] {
			if c == first {
				rest := input[i+1:]
				largest := calcLargest(rest)
				return int(10*(first-'0') + (largest - '0'))
			}
		}
	}
	return 0
}

func calcLargest(s string) rune {
	largest := rune(s[0])
	for _, v := range s[1:] {
		largest = max(largest, v)
	}
	return largest
}

func getInput() []string {
	var input []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input = append(input, s.Text())
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return input
}
