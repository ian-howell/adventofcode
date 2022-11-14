package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	numRows = 6
	numCols = 50
)

func main() {
	grid := makeGrid()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(s.Text())
		do(grid, s.Text())
		printGrid(grid)
		fmt.Println("------------------------------------------------------------")
	}

	total := 0
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if grid[r][c] == '#' {
				total++
			}
		}
	}
	fmt.Println(total)
}

func makeGrid() [][]byte {
	grid := [][]byte{}
	for r := 0; r < numRows; r++ {
		row := []byte{}
		for c := 0; c < numCols; c++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid [][]byte) {
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			fmt.Print(string(grid[r][c]))
		}
		fmt.Println()
	}
}

func do(grid [][]byte, instruction string) {
	parts := strings.Split(instruction, " ")

	if parts[0] == "rect" {
		r, c := parseRect(parts[1])
		doRect(grid, r, c)
		return
	}

	rc, amount := parseAmount(parts[2], parts[4])
	switch parts[1] {
	case "row":
		doRow(grid, rc, amount)
	case "column":
		doCol(grid, rc, amount)
	}
}

func parseRect(s string) (int, int) {
	parts := strings.Split(s, "x")
	r, _ := strconv.Atoi(parts[0])
	c, _ := strconv.Atoi(parts[1])
	return r, c
}

func parseAmount(s1, s2 string) (int, int) {
	parts := strings.Split(s1, "=")
	rc, _ := strconv.Atoi(parts[1])
	amount, _ := strconv.Atoi(s2)
	return rc, amount
}

func doRect(grid [][]byte, r, c int) {
	for r := r - 1; r >= 0; r-- {
		for c := c - 1; c >= 0; c-- {
			grid[c][r] = '#'
		}
	}
}

func doRow(grid [][]byte, r, amount int) {
	for i := 0; i < (amount % numCols); i++ {
		shiftRowOnce(grid, r)
	}
}

func shiftRowOnce(grid [][]byte, r int) {
	tmp := grid[r][numCols-1]
	for c := numCols - 1; c > 0; c-- {
		grid[r][c] = grid[r][c-1]
	}
	grid[r][0] = tmp
}

func doCol(grid [][]byte, c, amount int) {
	for i := 0; i < (amount % numRows); i++ {
		shiftColOnce(grid, c)
	}
}

func shiftColOnce(grid [][]byte, c int) {
	tmp := grid[numRows-1][c]
	for r := numRows - 1; r > 0; r-- {
		grid[r][c] = grid[r-1][c]
	}
	grid[0][c] = tmp
}
