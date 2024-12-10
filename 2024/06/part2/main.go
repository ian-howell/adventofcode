package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := getInput()
	pathCells := getPathCells(grid)

	total := 0
	for pathCell := range pathCells {
		func() {
			addObstacle(grid, pathCell)
			defer removeObstacle(grid, pathCell)
			if hasCycle(grid) {
				total++
			}
		}()
	}
	fmt.Println(total)
}

func removeObstacle(grid [][]byte, cell Cell) {
	grid[cell.Row][cell.Col] = '#'
}

func addObstacle(grid [][]byte, cell Cell) {
	grid[cell.Row][cell.Col] = '#'
}

func hasCycle(grid [][]byte) bool {
	guard := findGuard(grid)

	visited := map[Guard]struct{}{}

	for inBounds(guard.Pos, grid) {
		if _, ok := visited[guard]; ok {
			return true
		}
		visited[guard] = struct{}{}
		guard.Step(grid)
	}
	return false
}

func getPathCells(grid [][]byte) map[Cell]struct{} {
	guard := findGuard(grid)
	pathCells := map[Cell]struct{}{}

	for inBounds(guard.Pos, grid) {
		pathCells[guard.Pos] = struct{}{}
		guard.Step(grid)
	}
	return pathCells
}

func findGuard(grid [][]byte) Guard {
	for r, row := range grid {
		for c, val := range row {
			if val == '^' {
				return Guard{
					Pos: Cell{Row: r, Col: c},
					Dir: North,
				}
			}
		}
	}
	return Guard{}
}

func inBounds(pos Cell, grid [][]byte) bool {
	return betweenInclusive(0, pos.Row, len(grid)-1) &&
		betweenInclusive(0, pos.Col, len(grid[0])-1)
}

func betweenInclusive(lb, x, ub int) bool {
	return lb <= x && x <= ub
}

func getInput() [][]byte {
	grid := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
