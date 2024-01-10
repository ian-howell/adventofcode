package main

import (
	"bytes"
	"fmt"
)

func main() {
	grid := getInput()
	fmt.Println(bfs(grid))
}

func bfs(grid []string) int {
	startIndex := findStart(grid)
	startRow, startCol := toRowCol(grid, startIndex)
	distanceTo := map[int]int{startIndex: 0}

	q := []int{}
	realStartSymbol := getRealStartSymbol(grid, startIndex)
	// Initialize the queue with the start symbol's neighbors. Do this to avoid dealing with the S symbol
	for _, neighbor := range getNeighborsForSymbol(realStartSymbol) {
		nr, nc := startRow+neighbor[0], startCol+neighbor[1]
		if isValid(grid, nr, nc) {
			index := toIndex(grid, nr, nc)
			distanceTo[index] = 1
			q = append(q, index)
		}
	}

	for len(q) > 0 {
		var u int
		u, q = q[0], q[1:]

		for _, v := range getNeighbors(grid, u) {
			if _, ok := distanceTo[v]; !ok {
				distanceTo[v] = distanceTo[u] + 1
				q = append(q, v)
			}
		}
	}

	max := -1
	for _, val := range distanceTo {
		if val > max {
			max = val
		}
	}

	return max
}

func getNeighbors(grid []string, index int) []int {
	rowColNeighbors := getNeighborsForSymbol(getSymbolAtIndex(grid, index))
	r, c := toRowCol(grid, index)
	indexNeighbors := []int{}
	for _, rcn := range rowColNeighbors {
		nr, nc := r+rcn[0], c+rcn[1]
		if isValid(grid, nr, nc) {
			indexNeighbors = append(indexNeighbors, toIndex(grid, nr, nc))
		}
	}
	return indexNeighbors
}

func getNeighborsForSymbol(symbol byte) [][]int {
	up := []int{-1, 0}
	down := []int{1, 0}
	left := []int{0, -1}
	right := []int{0, 1}
	switch symbol {
	case '|':
		return [][]int{up, down}
	case '-':
		return [][]int{right, left}
	case 'J':
		return [][]int{left, up}
	case 'F':
		return [][]int{down, right}
	case '7':
		return [][]int{down, left}
	case 'L':
		return [][]int{up, right}
	}
	// Should be unreachable
	return nil
}

func findStart(grid []string) int {
	for i := 0; i < len(grid)*len(grid[0]); i++ {
		if getSymbolAtIndex(grid, i) == 'S' {
			return i
		}
	}
	return -1
}

func getSymbolAtIndex(grid []string, index int) byte {
	r, c := toRowCol(grid, index)
	return grid[r][c]
}

func toRowCol(grid []string, index int) (int, int) {
	r := index / len(grid[0])
	c := index % len(grid[0])
	return r, c
}

func toIndex(grid []string, r int, c int) int {
	return r*len(grid[0]) + c
}

func getRealStartSymbol(grid []string, startIndex int) byte {
	r, c := toRowCol(grid, startIndex)

	connections := 0
	if isValid(grid, r-1, c) && bytes.ContainsAny([]byte{grid[r-1][c]}, "F7|") {
		connections |= 8
	}
	if isValid(grid, r+1, c) && bytes.ContainsAny([]byte{grid[r+1][c]}, "JL|") {
		connections |= 4
	}
	if isValid(grid, r, c-1) && bytes.ContainsAny([]byte{grid[r][c-1]}, "FL-") {
		connections |= 2
	}
	if isValid(grid, r, c+1) && bytes.ContainsAny([]byte{grid[r][c+1]}, "J7-") {
		connections |= 1
	}

	connectionMap := map[int]byte{
		0b1100: '|',
		0b0011: '-',
		0b1010: 'J',
		0b1001: 'L',
		0b0110: '7',
		0b0101: 'F',
	}

	return connectionMap[connections]
}

func isValid(grid []string, r int, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
