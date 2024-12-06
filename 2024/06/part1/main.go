package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := getInput()
	pathCells := getPathCells(grid)
	fmt.Println(len(pathCells))
}

func getPathCells(grid []string) map[Cell]struct{} {
	guard := findGuard(grid)
	pathCells := map[Cell]struct{}{}

	for inBounds(guard.Pos, grid) {
		fmt.Println(guard.Pos)
		pathCells[guard.Pos] = struct{}{}
		guard.Step(grid)
	}
	return pathCells
}

func findGuard(grid []string) Guard {
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

func inBounds(pos Cell, grid []string) bool {
	return betweenInclusive(0, pos.Row, len(grid)-1) &&
		betweenInclusive(0, pos.Col, len(grid[0])-1)
}

func betweenInclusive(lb, x, ub int) bool {
	return lb <= x && x <= ub
}

func getInput() []string {
	grid := []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, s.Text())
	}
	return grid
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
