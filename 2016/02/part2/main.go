package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, c := 2, 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		for _, direction := range line {
			r, c = next(r, c, direction)
		}
		fmt.Print(string(keypad(r, c)))
	}
	fmt.Println()
}

func next(r, c int, d rune) (int, int) {
	delta := map[rune][2]int{
		'U': {-1, 0},
		'L': {0, -1},
		'D': {1, 0},
		'R': {0, 1},
	}[d]
	nr, nc := clamp(r+delta[0]), clamp(c+delta[1])
	if keypad(nr, nc) == ' ' {
		return r, c
	}
	return nr, nc
}

func clamp(x int) int {
	if x < 0 {
		return 0
	}
	if x > 4 {
		return 4
	}
	return x
}

func keypad(r, c int) rune {
	return [][]rune{
		{' ', ' ', '1', ' ', ' '},
		{' ', '2', '3', '4', ' '},
		{'5', '6', '7', '8', '9'},
		{' ', 'A', 'B', 'C', ' '},
		{' ', ' ', 'D', ' ', ' '},
	}[r][c]
}
