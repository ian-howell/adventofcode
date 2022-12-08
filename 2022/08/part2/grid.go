package main

import (
	"bufio"
	"fmt"
	"os"
)

type Row []int
type Grid []Row

func (g Grid) Or(h Grid) Grid {
	result := InitGrid(g)
	for r := range g {
		for c := range g {
			result[r][c] = g[r][c] | h[r][c]
		}
	}
	return result
}

func (g Grid) Count() int {
	total := 0
	for _, row := range g {
		for _, item := range row {
			total += item
		}
	}
	return total
}

func GetGrid() Grid {
	var grid Grid
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		var row Row
		for i := 0; i < len(line); i++ {
			row = append(row, atoi(line[i]))
		}
		grid = append(grid, row)
	}
	return grid
}

func PrintGrid(grid Grid) {
	for _, row := range grid {
		for _, item := range row {
			fmt.Printf("%d ", item)
		}
		fmt.Println()
	}
}

func InitGrid(grid Grid) Grid {
	result := make(Grid, len(grid))
	for i := range result {
		result[i] = make(Row, len(grid[0]))
	}
	return result
}
