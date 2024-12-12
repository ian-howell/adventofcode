package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct{ Row, Col int }

func main() {
	grid := getGrid()

	visited := map[Cell]struct{}{}
	total := 0
	for r, row := range grid {
		for c := range row {
			cell := Cell{r, c}
			if _, ok := visited[cell]; !ok {
				total += bfs(grid, cell, visited)
			}
		}
	}
	fmt.Println(total)
}

func bfs(grid [][]byte, start Cell, visited map[Cell]struct{}) int {
	queue := []Cell{start}
	visited[start] = struct{}{}
	symbol := grid[start.Row][start.Col]

	perimeter := 0
	area := 0

	for len(queue) > 0 {
		area++
		u := queue[0]
		queue = queue[1:]
		vs := next(u, grid, symbol)
		perimeter += (4 - len(vs))
		for _, v := range vs {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				queue = append(queue, v)
			}
		}
	}

	return area * perimeter
}

func next(cell Cell, grid [][]byte, symbol byte) []Cell {
	vs := []Cell{}
	dirMatrix := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range dirMatrix {
		v := Cell{
			Row: cell.Row + dir[0],
			Col: cell.Col + dir[1],
		}
		if inBounds(v, grid) && grid[v.Row][v.Col] == symbol {
			vs = append(vs, v)
		}
	}
	return vs
}

func inBounds(u Cell, grid [][]byte) bool {
	between := func(lb, x, ub int) bool { return lb <= x && x < ub }
	return between(0, u.Row, len(grid)) && between(0, u.Col, len(grid[0]))
}

func getGrid() [][]byte {
	grid := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}
