package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeMap struct {
	SrcStart  int
	DestStart int
	Len       int
}

type RangeMapSlice []RangeMap
type RangeMapSliceSlice []RangeMapSlice

func main() {
	seedRanges, almanac := getInputs()
	lowest := -1
	for _, seedRange := range seedRanges {
		// for i, seedRange := range seedRanges {
		// fmt.Printf("Beginning seed range #%d: %v\n", i, seedRange)
		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]; seed++ {
			candidate := almanac.Get(seed)
			if lowest == -1 || candidate < lowest {
				lowest = candidate
			}
		}
	}
	fmt.Println(lowest)
}

func (r RangeMapSlice) Get(index int) int {
	for _, rangeMap := range r {
		delta := index - rangeMap.SrcStart
		if 0 <= delta && delta < rangeMap.Len {
			return rangeMap.DestStart + delta
		}
	}
	return index
}

func (r RangeMapSliceSlice) Get(index int) int {
	for _, rangeMapSlice := range r {
		index = rangeMapSlice.Get(index)
	}
	return index
}

func getInputs() ([][2]int, RangeMapSliceSlice) {
	s := bufio.NewScanner(os.Stdin)

	// Get the first line containing the seeds
	s.Scan()

	seeds := [][2]int{}
	arr := toIntArray(strings.TrimPrefix(s.Text(), "seeds: "))
	for i := 0; i < len(arr); i += 2 {
		seeds = append(seeds, [2]int{arr[i], arr[i+1]})
	}

	// eat the empty line
	s.Scan()

	// eat the header for the first chunk
	s.Scan()

	almanac := RangeMapSliceSlice{}
	current := RangeMapSlice{}
	for s.Scan() {
		if s.Text() == "" {
			almanac = append(almanac, current)
			current = RangeMapSlice{}

			// eat the header for the next chunk
			s.Scan()

			continue
		}
		vals := toIntArray(s.Text())
		current = append(current, RangeMap{
			SrcStart:  vals[1],
			DestStart: vals[0],
			Len:       vals[2],
		})
	}
	almanac = append(almanac, current)

	return seeds, almanac
}

func toIntArray(s string) []int {
	nums := []int{}
	for _, numStr := range strings.Fields(s) {
		i, _ := strconv.Atoi(numStr)
		nums = append(nums, i)
	}
	return nums
}
