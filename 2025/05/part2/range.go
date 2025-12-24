package main

import (
	"fmt"
	"slices"
	"strings"
)

type Range struct {
	LowerBound int
	UpperBound int
}

func NewRange(raw string) Range {
	r, l, _ := strings.Cut(raw, "-")
	return Range{atoi(r), atoi(l)}
}

func (r Range) In(val int) bool {
	return val >= r.LowerBound && val <= r.UpperBound
}

func (r Range) String() string {
	return fmt.Sprintf("(%d, %d)", r.LowerBound, r.UpperBound)
}

type RangeCollection struct {
	ranges []Range
}

func (rc *RangeCollection) Add(r Range) {
	if len(rc.ranges) == 0 {
		rc.ranges = append(rc.ranges, r)
		return
	}

	i := 0
	for i < len(rc.ranges) && r.LowerBound > rc.ranges[i].LowerBound {
		i++
	}

	if i == len(rc.ranges) {
		rc.ranges = append(rc.ranges, r)
		return
	}

	if r.LowerBound > rc.ranges[i].UpperBound {
		rc.ranges = slices.Insert(rc.ranges, i, r)
		return
	}

	if r.UpperBound < rc.ranges[i].LowerBound {
		rc.ranges = slices.Insert(rc.ranges, i, r)
	}
}

func (rc RangeCollection) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, range_ := range rc.ranges {
		sb.WriteString(fmt.Sprintf("  %s,\n", range_.String()))
	}
	sb.WriteString("]\n")
	return sb.String()
}
