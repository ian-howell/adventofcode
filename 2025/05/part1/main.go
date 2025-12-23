package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, ids := getInput()
	answer := findFreshIds(ranges, ids)
	fmt.Println(answer)
}

type Range struct {
	LowerBound int
	UpperBound int
}

func findFreshIds(ranges []Range, ids []int) int {
	numFreshIds := 0
	for _, id := range ids {
		for _, range_ := range ranges {
			if inRange(id, range_) {
				numFreshIds++
				break
			}
		}
	}
	return numFreshIds
}

func inRange(val int, range_ Range) bool {
	return val >= range_.LowerBound && val <= range_.UpperBound
}

func getInput() (ranges []Range, ids []int) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		ranges = append(ranges, parseRange(s.Text()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}

	for s.Scan() {
		ids = append(ids, atoi(s.Text()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}

	return ranges, ids
}

func parseRange(raw string) Range {
	r, l, _ := strings.Cut(raw, "-")
	return Range{atoi(r), atoi(l)}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
