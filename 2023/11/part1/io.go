package main

import (
	"bufio"
	"fmt"
	"os"
)

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
