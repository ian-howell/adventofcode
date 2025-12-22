package main

import (
	"fmt"
	"io"
	"os"
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

func isInvalid(id int) bool {
	s := strconv.Itoa(id)
	for i := 1; i <= len(s)/2; i++ {
		q := s[:i]
		if isARepetition(q, s) {
			return true
		}
	}
	return false
}

func isARepetition(q, id string) bool {
	for found := true; found; id, found = strings.CutPrefix(id, q) {
	}
	return id == ""
}
