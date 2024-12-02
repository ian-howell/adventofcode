package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	numSafeReports := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if isSafe(s.Text()) {
			numSafeReports++
		}
	}
	fmt.Println(numSafeReports)
}

func isSafe(s string) bool {
	nums := toInts(s)
	if len(nums) < 2 {
		// This doesn't happen for my input, but... safety first!
		return true
	}

	// If the list appears to be in descending order, simplify things by reversing it
	if nums[0] > nums[1] {
		slices.Reverse(nums)
	}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff <= 0 || diff > 3 {
			return false
		}
	}

	return true
}

func toInts(s string) []int {
	nums := []int{}
	for _, str := range strings.Fields(s) {
		nums = append(nums, atoi(str))
	}
	return nums
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to int: %v", s, err))
	}
	return i
}
