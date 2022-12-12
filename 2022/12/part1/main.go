package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		grid  gridGraph
		start Cell
		end   Cell
	)

	s := bufio.NewScanner(os.Stdin)
	row := 0
	for s.Scan() {
		col := strings.IndexByte(s.Text(), 'S')
		if col >= 0 {
			start = Cell{row, col}
		}
		col = strings.IndexByte(s.Text(), 'E')
		if col >= 0 {
			end = Cell{row, col}
		}
		line := strings.ReplaceAll(s.Text(), "S", "a")
		line = strings.ReplaceAll(line, "E", "z")
		grid.grid = append(grid.grid, line)
		row++
	}

	path := BreadthFirstSearch[Cell](grid, start, end)
	fmt.Println(len(path) - 1)
}

type Graph[T any] interface {
	Neighbors(node T) []T
}

func BreadthFirstSearch[T comparable](graph Graph[T], start, end T) []T {
	cameFrom := map[T]*T{start: nil}
	queue := []T{start}

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		if u == end {
			return createPath(cameFrom, end)
		}

		for _, v := range graph.Neighbors(u) {
			if _, ok := cameFrom[v]; !ok {
				cameFrom[v] = &u
				queue = append(queue, v)
			}
		}
	}

	return nil
}

func createPath[T comparable](cameFrom map[T]*T, end T) []T {
	path := []T{end}
	for next := cameFrom[end]; next != nil; next = cameFrom[end] {
		path = append([]T{*next}, path...)
		end = *next
	}
	return path
}

type Cell struct {
	Row int
	Col int
}

type gridGraph struct {
	grid []string
}

func (g gridGraph) Neighbors(node Cell) []Cell {
	neighbors := []Cell{}
	for _, m := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		newCell := Cell{Row: node.Row + m[0], Col: node.Col + m[1]}
		if g.IsValid(newCell) && g.grid[newCell.Row][newCell.Col] <= (g.grid[node.Row][node.Col]+1) {
			neighbors = append(neighbors, newCell)
		}
	}
	return neighbors
}

func (g gridGraph) IsValid(cell Cell) bool {
	return cell.Row >= 0 && cell.Col >= 0 &&
		cell.Row < len(g.grid) && cell.Col < len(g.grid[0])
}
