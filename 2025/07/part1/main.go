package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	grid := getInput()
	startCol := findStartColumn(grid)
	cols := map[int]bool{startCol: true}

	count := 0
	for _, row := range grid {
		newCols := map[int]bool{}
		for c, val := range row {
			if cols[c] {
				if val == '^' {
					newCols[c-1] = true
					newCols[c+1] = true
					count++
				} else {
					newCols[c] = true
				}
			}
		}
		cols = newCols
	}
	fmt.Println(count)
}

func findStartColumn(grid [][]byte) int {
	return bytes.IndexByte(grid[0], 'S')
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
