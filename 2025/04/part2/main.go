package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := getInput()
	answer := countRolls(grid)
	fmt.Println(answer)
}

func countRolls(grid [][]rune) int {
	total := 0
	numRemoved := removeRolls(grid)
	for numRemoved > 0 {
		total += numRemoved
		numRemoved = removeRolls(grid)
	}
	return total
}

func removeRolls(grid [][]rune) int {
	rolls := 0
	for r, row := range grid {
		for c, val := range row {
			if val == '.' {
				continue
			}
			if isAccessible(grid, r, c) {
				rolls++
				grid[r][c] = '.'
			}
		}
	}
	return rolls
}

func isAccessible(grid [][]rune, r, c int) bool {
	neighbors := 0
	matrix := [][]int{
		{-1, -1}, {-1, 0}, {-1, +1},
		{0, -1} /*center*/, {0, +1},
		{+1, -1}, {+1, 0}, {+1, +1}}
	for _, dir := range matrix {
		nr, nc := r+dir[0], c+dir[1]
		if !inBounds(grid, nr, nc) {
			continue
		}

		if grid[nr][nc] == '@' {
			neighbors++
		}
	}
	return neighbors < 4
}

func inBounds(grid [][]rune, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func getInput() [][]rune {
	grid := [][]rune{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return grid
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}
