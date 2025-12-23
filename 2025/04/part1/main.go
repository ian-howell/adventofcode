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

func countRolls(grid []string) int {
	rolls := 0
	for r, row := range grid {
		for c, val := range row {
			if val == '.' {
				fmt.Print(".")
				continue
			}
			if isAccessible(grid, r, c) {
				rolls++
				fmt.Print(red("X"))
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
	return rolls
}

func isAccessible(grid []string, r, c int) bool {
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

func inBounds(grid []string, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func getInput() []string {
	grid := []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, s.Text())
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return grid
}

func printGrid(grid []string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}
