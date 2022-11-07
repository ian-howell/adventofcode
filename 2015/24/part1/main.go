package main

import (
	"fmt"
	"sort"
)

func main() {
	weights := readInput()
	totalWeight := SumInts(weights)

	targetWeight := totalWeight / 3

	groups := GreedilyGroup(targetWeight, weights)
	smallestGroupLen := len(weights)
	lengthMap := map[int][][]int{}
	for _, group := range groups {
		lengthMap[len(group)] = append(lengthMap[len(group)], group)
		if len(group) < smallestGroupLen {
			smallestGroupLen = len(group)
		}
	}

	smallestQE := -1
	for _, group := range lengthMap[smallestGroupLen] {
		if qe := QE(group); smallestQE == -1 || qe < smallestQE {
			smallestQE = qe
		}
	}
	fmt.Println(smallestQE)
}

func readInput() []int {
	nums := []int{}
	var num int
	for _, err := fmt.Scan(&num); err == nil; _, err = fmt.Scan(&num) {
		nums = append(nums, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums
}

func SumInts(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// GreedilyGroup takes a list of numbers and a target, and returns the smallest subset of those numbers whose
// sum is the target
func GreedilyGroup(target int, nums []int) [][]int {
	if len(nums) == 0 || target <= 0 {
		return nil
	}

	candidates := [][]int{}
	for i, num := range nums {
		if num == target {
			candidates = append(candidates, []int{num})
		} else if num <= target {
			successors := GreedilyGroup(target-num, nums[i+1:])
			for _, successor := range successors {
				current := append([]int{num}, successor...)
				candidates = append(candidates, current)
			}
		}
	}

	return candidates
}

func SetSubtract(larger, smaller []int) []int {
	// Maybe useful in part 2?
	final := []int{}
	for _, val := range larger {
		if !contains(smaller, val) {
			final = append(final, val)
		}
	}
	return final
}

func contains(set []int, val int) bool {
	for _, x := range set {
		if x == val {
			return true
		}
	}
	return false
}

func QE(nums []int) int {
	qe := 1
	for _, num := range nums {
		qe *= num
	}
	return qe
}
