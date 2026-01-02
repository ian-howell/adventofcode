package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	timelineCounter := NewTimelineCounter(os.Stdin)
	fmt.Println(timelineCounter.Count())
}

type Cell struct {
	r, c int
}

type TimelineCounter struct {
	grid  [][]byte
	start Cell
	memo  map[Cell]int
}

func NewTimelineCounter(r io.Reader) *TimelineCounter {
	scanner := bufio.NewScanner(r)
	tc := &TimelineCounter{memo: map[Cell]int{}}
	// Grab the first line
	scanner.Scan()
	startCol := strings.Index(scanner.Text(), "S")
	tc.start = Cell{r: 0, c: startCol}

	scan := func() bool { scanner.Scan(); return scanner.Scan() }
	for scan() {
		tc.grid = append(tc.grid, []byte(scanner.Text()))
	}
	return tc
}

func (tc *TimelineCounter) Count() int {
	return tc.count(tc.start)
}

func (tc *TimelineCounter) count(cell Cell) int {
	if cell.r >= len(tc.grid) {
		return 1
	}

	if val, ok := tc.memo[cell]; ok {
		return val
	}

	if tc.grid[cell.r][cell.c] == '^' {
		tc.memo[cell] = 0 +
			tc.count(Cell{r: cell.r + 1, c: cell.c - 1}) +
			tc.count(Cell{r: cell.r + 1, c: cell.c + 1})
	} else {
		tc.memo[cell] = tc.count(Cell{r: cell.r + 1, c: cell.c})
	}

	return tc.memo[cell]
}
