package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, freq := readInput()
	fmt.Println(findSimilarityScore(nums, freq))
}

func findSimilarityScore(nums []int, freq map[int]int) int {
	similarityScore := 0
	for _, num := range nums {
		similarityScore += num * freq[num]
	}
	return similarityScore
}

func readInput() ([]int, map[int]int) {
	nums := []int{}
	freq := map[int]int{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		// This panics for bad inputs (e.g. len(fields) == 0 and non-integer fields)
		nums = append(nums, atoi(fields[0]))
		freq[atoi(fields[1])]++
	}

	return nums, freq
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
