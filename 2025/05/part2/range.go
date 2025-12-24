package main

import (
	"fmt"
	"strings"
)

type Range struct {
	lb int
	ub int
}

func NewRange(raw string) Range {
	r, l, _ := strings.Cut(raw, "-")
	return Range{atoi(r), atoi(l)}
}

func (r Range) In(val int) bool {
	return val >= r.lb && val <= r.ub
}

func (r Range) String() string {
	return fmt.Sprintf("(%d, %d)", r.lb, r.ub)
}

func (r Range) Size() int {
	return r.ub - r.lb + 1
}
