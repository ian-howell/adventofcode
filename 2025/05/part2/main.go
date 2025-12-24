package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	ranges := getInput()
	answer := countFreshIds(ranges)
	fmt.Println(answer)
}

func countFreshIds(ranges []Range) int {
	slices.SortFunc(ranges, func(a Range, b Range) int {
		if a.lb < b.lb {
			return -1
		}
		return 1
	})

	format := "%-40v%-20v%-20v%-20v\n"
	count := ranges[0].Size()
	ptr := ranges[0].ub
	fmt.Printf(format, "RANGE", "SIZE", "NEWLY ADDED", "PTR")
	fmt.Printf(format, ranges[0], ranges[0].Size(), ranges[0].Size(), ptr)
	for i := 1; i < len(ranges); i++ {
		old := count
		if ranges[i].lb > ptr {
			count += ranges[i].Size()
			ptr = ranges[i].ub
		} else if ranges[i].ub > ptr {
			count += ranges[i].ub - ptr
			ptr = ranges[i].ub
		}
		added := count - old
		fmt.Printf(format, ranges[i], ranges[i].Size(), added, ptr)
	}
	return count
}

func getInput() []Range {
	ranges := make([]Range, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		ranges = append(ranges, NewRange(s.Text()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}

	return ranges
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
