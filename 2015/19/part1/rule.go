package main

import (
	"fmt"
	"strings"
)

type Rule [2]string

func NewRule(s string) Rule {
	from, to, _ := strings.Cut(s, " => ")
	return Rule{from, to}
}

func (r Rule) From() string {
	return r[0]
}

func (r Rule) To() string {
	return r[1]
}

func (r Rule) String() string {
	return fmt.Sprintf("<%s => %s>", r[0], r[1])
}
