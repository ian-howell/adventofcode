package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := getInitialGrid()
	for i := 1; i <= 100; i++ {
		grid = gameOfLife(grid)
	}
	fmt.Println(countLights(grid))
}

func gameOfLife(grid []string) []string {
	newGrid := []string{}
	for r := 0; r < len(grid); r++ {
		sb := strings.Builder{}
		for c := 0; c < len(grid[0]); c++ {
			if isCorner(grid, r, c) {
				sb.WriteByte('#')
			} else {
				n := numNeighbors(grid, r, c)
				if grid[r][c] == '#' {
					if n == 2 || n == 3 {
						sb.WriteByte('#')
					} else {
						sb.WriteByte('.')
					}
				} else {
					if n == 3 {
						sb.WriteByte('#')
					} else {
						sb.WriteByte('.')
					}
				}
			}
		}
		newGrid = append(newGrid, sb.String())
	}
	return newGrid
}

func numNeighbors(grid []string, r, c int) int {
	n := 0
	candidates := [][2]int{{+0, +1}, {+1, +1}, {+1, +0}, {+1, -1}, {+0, -1}, {-1, -1}, {-1, +0}, {-1, +1}}
	for _, candidate := range candidates {
		nr, nc := r+candidate[0], c+candidate[1]
		if isValid(grid, nr, nc) && grid[nr][nc] == '#' {
			n++
		}
	}
	return n
}

func isValid(grid []string, r, c int) bool {
	return r >= 0 && c >= 0 && r < len(grid) && c < len(grid[0])
}

func isCorner(grid []string, r, c int) bool {
	return (r == 0 && c == 0) ||
		(r == 0 && c == len(grid[0])-1) ||
		(r == len(grid)-1 && c == 0) ||
		(r == len(grid)-1 && c == len(grid[0])-1)
}

func countLights(grid []string) int {
	n := 0
	for _, row := range grid {
		for _, val := range row {
			if val == '#' {
				n++
			}
		}
	}
	return n
}

func getInitialGrid() []string {
	grid := []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, s.Text())
	}
	return grid
}

func printGrid(grid []string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
