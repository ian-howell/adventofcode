package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// 326553 is too low
	ranges := readRanges(os.Stdin)
	dumpInputStats(ranges)

	fmt.Println("ANSWER")
	fmt.Println(countTotalInvalidIds(ranges))
}

func dumpInputStats(ranges []Range) {
	fmt.Println("RANGES")
	fmt.Println("======")

	total := 0
	for _, range_ := range ranges {
		size := range_.UpperBound - range_.LowerBound + 1
		total += size
		fmt.Printf("LB: %-15d UB: %-15d len: %-15d\n", range_.LowerBound, range_.UpperBound, size)
	}

	fmt.Println()
	fmt.Printf("Total values: %d\n", total)
}

// magnitude is a helper that was used to determine if the lengths of the ids ever vary by more than 1
func magnitude(range_ Range) int {
	lb := strconv.Itoa(range_.LowerBound)
	ub := strconv.Itoa(range_.UpperBound)
	d := len(lb) - len(ub)
	if d < 0 {
		return -d
	}
	return d
}

type Range struct {
	LowerBound int
	UpperBound int
}

func readRanges(r io.Reader) []Range {
	var unparsedInput string
	fmt.Fscan(r, &unparsedInput)
	inputs := strings.Split(unparsedInput, ",")
	ranges := make([]Range, 0, len(inputs))
	for _, input := range inputs {
		lb, ub, _ := strings.Cut(input, "-")
		var range_ Range
		range_.LowerBound, _ = strconv.Atoi(lb)
		range_.UpperBound, _ = strconv.Atoi(ub)
		ranges = append(ranges, range_)
	}
	return ranges
}

func countTotalInvalidIds(ranges []Range) int {
	totalInvalidRanges := 0
	for _, range_ := range ranges {
		totalForRange := countInvalidIds(range_)
		totalInvalidRanges += totalForRange
		fmt.Printf("Total for %v: %d\n", range_, totalForRange)
	}
	return totalInvalidRanges
}

func countInvalidIds(range_ Range) int {
	total := 0
	for val := range_.LowerBound; val <= range_.UpperBound; val++ {
		if isInvalid(val) {
			total += val
			fmt.Println(total, val)
		}
	}
	return total
}

func nextPowerOfTen(x int) int {
	numDigits := int(math.Log10(float64(x))) + 1
	return int(math.Pow10(numDigits))
}

func findFirstInvalidId(range_ Range) (int, bool) {
	for id := range_.LowerBound; id <= range_.UpperBound; id++ {
		if isInvalid(id) {
			return id, true
		}
	}
	return 0, false
}

func getRepeatingSequence(id int) int {
	s := strconv.Itoa(id)
	firstHalf, _ := strconv.Atoi(s[len(s)/2:])
	return firstHalf
}

func isInvalid(id int) bool {
	s := strconv.Itoa(id)
	if len(s)%2 == 1 {
		return false
	}
	m := len(s) / 2
	return s[:m] == s[m:]
}

// merge is an interesting function. The result is the maximum value for each power of ten
func merge(a, b int) int {
	result := make([]string, 0)
	for a > 0 || b > 0 {
		result = append(result, strconv.Itoa(max(a%10, b%10)))
		a /= 10
		b /= 10
	}
	slices.Reverse(result)
	converted, _ := strconv.Atoi(strings.Join(result, ""))
	return converted
}
