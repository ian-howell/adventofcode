package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const debug = true

func main() {
	numSafeReports := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		nums := toInts(s.Text())
		if isSafeTolerant(nums) {
			Debugln("TRUE!")
			numSafeReports++
		}
		Debugln()
	}
	fmt.Println(numSafeReports)
}

// isSafeTolerant is the same as isSafe, but it allows for one "mistake". That is, if the original nums
// is not safe, but could be made safe by removing a single num, isSafeTolerant returns true
func isSafeTolerant(nums []int) bool {
	Debugf("NUMS:   %v\n", nums)
	if isSafe(nums) {
		// The original list was safe
		return true
	}

	// Drop the first num and check if the shorter list is safe. If it is not, drop the second num
	// (instead of the first), and check again. Repeat this for all nums.
	Debugln("SHORTNUMS:")
	for i := 0; i < len(nums); i++ {
		shortNums := slices.Clone(nums)
		// Same as above (it's literally how slices.Clone is implemented)
		// shortNums := append([]int{}, nums...)
		shortNums = slices.Delete(shortNums, i, i+1)
		Debugf("        %v\n", shortNums)
		if isSafe(shortNums) {
			return true
		}
	}
	return false
}

// isSafe returns true if both of the follosing are true:
// * nums is either strictly ascending or descending
// * The difference between any 2 consecutive nums does not exceed 3
func isSafe(nums []int) bool {
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

func Debugf(s string, args ...any) {
	if debug {
		fmt.Printf(s, args...)
	}
}

func Debugln(a ...any) {
	if debug {
		fmt.Println(a...)
	}
}
