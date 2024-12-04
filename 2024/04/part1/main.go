package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	X int
	Y int
}

var (
	down      = Direction{X: 0, Y: 1}
	up        = Direction{X: 0, Y: -1}
	right     = Direction{X: 1, Y: 0}
	left      = Direction{X: -1, Y: 0}
	downRight = Direction{X: 1, Y: 1}
	upRight   = Direction{X: 1, Y: -1}
	downLeft  = Direction{X: -1, Y: 1}
	upLeft    = Direction{X: -1, Y: -1}
)

func main() {
	crosswordPuzzle := getInput()
	fmt.Println(countXMASForCrossword(crosswordPuzzle))
}

func countXMASForCrossword(crossword [][]byte) int {
	total := 0
	for r, row := range crossword {
		for c := range row {
			total += countXMASForCell(crossword, r, c)
		}
	}
	return total
}

func countXMASForCell(crossword [][]byte, r, c int) int {
	directionMatrix := []Direction{
		up, down, left, right,
		downRight, downLeft, upRight, upLeft,
	}

	total := 0
	for _, dir := range directionMatrix {
		if matchesDirection(crossword, r, c, dir) {
			total++
		}
	}
	return total
}

func matchesDirection(crossword [][]byte, r, c int, dir Direction) bool {
	for i, letter := range "XMAS" {
		row := r + i*dir.X
		col := c + i*dir.Y
		if !inBounds(crossword, row, col) || crossword[row][col] != byte(letter) {
			return false
		}
	}
	return true
}

func inBounds(crossword [][]byte, r, c int) bool {
	return 0 <= r && r < len(crossword) &&
		0 <= c && c < len(crossword[0])
}

func getInput() [][]byte {
	crossword := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		crossword = append(crossword, []byte(s.Text()))
	}
	return crossword
}
