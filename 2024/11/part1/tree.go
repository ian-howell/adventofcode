package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Leaves int
}

func NewTree(val int) *Tree {
	return &Tree{
		Value:  val,
		Left:   nil,
		Right:  nil,
		Leaves: 1,
	}
}

func (t *Tree) IsInternal() bool {
	return t.Leaves > 1
}

func (t *Tree) String() string {
	if !t.IsInternal() {
		return strconv.Itoa(t.Value)
	}

	return fmt.Sprintf("%s %s", t.Left, t.Right)
}
