package main

import "fmt"

func main() {
	input := getInput()
	total := 0
	for _, seq := range input {
		total += predict(seq)
	}
	fmt.Println(total)
}

func predict(seq []int) int {
	diff := findDifferenceSequence(seq)
	if allZeros(diff) {
		return seq[0]
	}

	return seq[0] - predict(diff)
}

func findDifferenceSequence(seq []int) []int {
	diff := make([]int, 0, len(seq)-1)
	for i := 1; i < len(seq); i++ {
		diff = append(diff, seq[i]-seq[i-1])
	}
	return diff
}

func allZeros(seq []int) bool {
	for val := range seq {
		if val != 0 {
			return false
		}
	}
	return true
}
