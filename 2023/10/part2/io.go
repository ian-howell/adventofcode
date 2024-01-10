package main

import (
	"bufio"
	"os"
)

func getInput() []string {
	s := bufio.NewScanner(os.Stdin)
	grid := []string{}
	for s.Scan() {
		grid = append(grid, s.Text())
	}
	return grid
}
