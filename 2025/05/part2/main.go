package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ranges := getInput()
	fmt.Println(ranges)
	// answer := findFreshIds(ranges, ids)
	// fmt.Println(answer)
}

func findFreshIds(ranges []Range, ids []int) int {
	numFreshIds := 0
	for _, id := range ids {
		for _, range_ := range ranges {
			if range_.In(id) {
				numFreshIds++
				break
			}
		}
	}
	return numFreshIds
}

func getInput() (ranges RangeCollection) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		ranges.Add(NewRange(s.Text()))
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
