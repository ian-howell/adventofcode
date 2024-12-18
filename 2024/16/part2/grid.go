package main

import (
	"bufio"
	"os"
	"strings"
)

func getInput() Grid {
	grid := Grid{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}

type Grid [][]byte

func (g Grid) Find(b byte) Cell {
	for r, row := range g {
		for c, val := range row {
			if b == val {
				return Cell{r, c}
			}
		}

	}
	// This should really be an error, but meh. AOC
	return Cell{-1, -1}
}

func (g Grid) At(cell Cell) byte {
	return g[cell.Row][cell.Col]
}

func (g Grid) String() string {
	sb := strings.Builder{}
	for _, row := range g {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	return strings.TrimRight(sb.String(), "\n")
}
