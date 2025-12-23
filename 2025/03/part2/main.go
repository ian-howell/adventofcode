package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := getInput()

	answer := calcTotalJoltage(input)
	fmt.Println(answer)
}

func calcTotalJoltage(inputs [][]byte) int {
	total := 0
	for _, input := range inputs {
		joltage := calcJoltage(input)
		// fmt.Printf("input: %-15s joltage: %d\n", input, joltage)
		total += joltage
	}
	return total
}

func calcJoltage(input []byte) int {
	result := make([]byte, 0)
	for i := 12; i > 0; i-- {
		pivot := len(input) - i + 1
		candidates := input[:pivot]
		largest, indexOfLargest := calcLargest(candidates)
		fmt.Printf("%s%s", candidates[:indexOfLargest], red(string(largest)))
		input = input[indexOfLargest+1:]
		result = append(result, largest)
	}
	fmt.Printf("%v -> %v\n", string(input), red(string(result)))
	return byteSliceToInt(result)
}

func calcLargest(input []byte) (byte, int) {
	indexOfLargest := 0
	largest := input[indexOfLargest]
	for i, v := range input[1:] {
		if v > largest {
			indexOfLargest = i + 1
			largest = v
		}
		largest = max(largest, v)
	}
	return largest, indexOfLargest
}

func byteSliceToInt(s []byte) int {
	value := 0
	for _, c := range s {
		value *= 10
		value += int(c - '0')
	}
	return value
}

func getInput() [][]byte {
	var input [][]byte
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input = append(input, bytes.Clone(s.Bytes()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return input
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}
