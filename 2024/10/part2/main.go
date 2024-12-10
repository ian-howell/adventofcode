package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct {
	Row int
	Col int
}

func main() {
	grid := getGrid()
	trailHeadScores := findTrailHeadScores(grid)
	total := 0
	for _, score := range trailHeadScores {
		total += score
	}
	fmt.Println(total)
}

func findTrailHeadScores(grid [][]int) map[Cell]int {
	trailHeads := map[Cell]int{}
	for r, row := range grid {
		for c, val := range row {
			if val == 9 {
				scoreTrailHeads(grid, Cell{Row: r, Col: c}, trailHeads)
			}
		}
	}
	return trailHeads
}

func scoreTrailHeads(grid [][]int, start Cell, trailHeads map[Cell]int) {
	queue := []Cell{start}
	numPathsToCell := map[Cell]int{start: 1}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, neighbor := range neighbors(grid, u) {
			if _, ok := numPathsToCell[neighbor]; !ok {
				queue = append(queue, neighbor)
			}
			numPathsToCell[neighbor] += numPathsToCell[u]
		}
	}

	isTrailHead := func(c Cell) bool { return grid[c.Row][c.Col] == 0 }
	for cell, numPaths := range numPathsToCell {

		if isTrailHead(cell) {
			trailHeads[cell] += numPaths
		}
	}
}

func neighbors(grid [][]int, cell Cell) []Cell {
	neighbors := []Cell{}
	dirMatrix := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range dirMatrix {
		newCell := Cell{
			Row: cell.Row + dir[0],
			Col: cell.Col + dir[1],
		}

		if inBounds(grid, newCell) {
			delta := grid[cell.Row][cell.Col] - grid[newCell.Row][newCell.Col]
			if delta == 1 {
				neighbors = append(neighbors, newCell)
			}
		}
	}
	return neighbors
}

func inBounds(grid [][]int, cell Cell) bool {
	return between(0, cell.Row, len(grid)) && between(0, cell.Col, len(grid[0]))
}

func between(lb, x, ub int) bool {
	return lb <= x && x < ub
}

func getGrid() [][]int {
	grid := [][]int{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		row := make([]int, 0, len(line))
		for i := 0; i < len(line); i++ {
			row = append(row, byteToInt(line[i]))
		}
		grid = append(grid, row)
	}
	return grid
}

func byteToInt(b byte) int {
	return int(b & 0x0f)
}
