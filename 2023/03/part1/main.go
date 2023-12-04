package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid [][]byte

func (g Grid) Valid(r, c int) bool { return 0 <= r && r < len(g) && 0 <= c && c < len(g[0]) }

func main() {
	grid := readGrid()
	total := 0
	for _, partNum := range getPartNums(grid) {
		total += partNum
	}
	fmt.Println(total)
}

func readGrid() Grid {
	grid := Grid{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		row := []byte(s.Text())
		grid = append(grid, row)
	}
	return grid
}

func getPartNums(grid Grid) []int {
	partNums := []int{}

	for r, row := range grid {
		for c, val := range row {
			if isSymbol(val) {
				// fmt.Printf("Found a symbol (%v) at (%d, %d)\n", string(val), r, c)
				partNums = append(partNums, getNeighboringPartNums(r, c, grid)...)
			}
		}
	}

	return partNums
}

func getNeighboringPartNums(r, c int, grid Grid) []int {
	partNums := []int{}
	dirMatrix := [][]int{
		{+0, +1}, // down
		{+0, -1}, // up
		{+1, +0}, // right
		{-1, +0}, // left
		{+1, +1}, // down-right
		{+1, -1}, // down-left
		{-1, +1}, // up-right
		{-1, -1}, // up-left
	}

	startPositions := map[[2]int]struct{}{}
	for _, m := range dirMatrix {
		newRow, newCol := r+m[0], c+m[1]
		if grid.Valid(newRow, newCol) && isNum(grid[newRow][newCol]) {
			// fmt.Printf("Found a number (%v) at (%d, %d)\n", string(grid[newRow][newCol]), newRow, newCol)
			num, startPosition := numAt(newRow, newCol, grid)
			if _, ok := startPositions[startPosition]; !ok {
				startPositions[startPosition] = struct{}{}
				partNums = append(partNums, num)
				// fmt.Printf("adding %v\n", num)
			}
		}
	}

	return partNums
}

func isNum(b byte) bool {
	return '0' <= b && b <= '9'
}

func isSymbol(b byte) bool {
	return b != '.' && !isNum(b)
}

func numAt(r, c int, grid Grid) (int, [2]int) {
	for grid.Valid(r, c-1) && isNum(grid[r][c-1]) {
		c--
	}

	startPosition := [2]int{r, c}
	num := 0
	for grid.Valid(r, c) && isNum(grid[r][c]) {
		num *= 10
		num += toInt(grid[r][c])
		c++
	}
	return num, startPosition
}

func toInt(b byte) int {
	return int(b - '0')
}
