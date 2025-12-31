package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Scroller struct {
	grid       [][]byte
	operations []byte

	cursor int
	value  int
}

func NewScroller(r io.Reader) *Scroller {
	scroller := &Scroller{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if strings.ContainsAny(s.Text(), "*+") {
			scroller.operations = []byte(s.Text())
			break
		}
		scroller.grid = append(scroller.grid, []byte(s.Text()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return scroller
}

func (s *Scroller) Scroll() bool {
	more := true
	s.value, more = s.get()
	return more
}

func (s *Scroller) Value() int {
	return s.value
}

func (s *Scroller) get() (int, bool) {
	if s.cursor >= len(s.grid[0]) {
		return -1, false
	}

	operation := s.getOperation()

	total := 1 - operation(0, 1) // clever... 1-(0+1)=1 and 1-(0*1)=0
	symbol := []string{"+", "*"}[total]
	fmt.Printf("%d ", total)
	num := s.getNum()
	for num > 0 {
		fmt.Printf("%v %d ", symbol, num)
		total = operation(total, num)

		s.cursor++
		num = s.getNum()
	}
	fmt.Printf("= %d\n", total)

	s.cursor++ // Move the cursor to the next column
	return total, true
}

func (s *Scroller) getOperation() Operation {
	return toOperation(s.operations[s.cursor])
}

func (s *Scroller) getNum() int {
	if s.cursor >= len(s.grid[0]) {
		return 0
	}
	total := 0
	for r := range s.grid {
		if s.grid[r][s.cursor] != ' ' {
			total *= 10
			total += toInt(s.grid[r][s.cursor])
		}
	}
	return total
}

func toInt(b byte) int {
	return int(b - '0')
}

type Operation func(a, b int) int

func add(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }

func toOperation(b byte) Operation {
	switch b {
	case '+':
		return add
	case '*':
		return multiply
	}
	panic(fmt.Sprintf("Unknown operation: %q", b))
}
