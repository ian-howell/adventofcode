package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	X int
	Y int
}

var (
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
	for r := 1; r < len(crossword)-1; r++ {
		for c := 1; c < len(crossword[0])-1; c++ {
			if isXMAS(crossword, r, c) {
				total++
			}
		}
	}
	return total
}

func isXMAS(crossword [][]byte, r, c int) bool {
	var (
		topRightRow = r + upRight.Y
		topRightCol = c + upRight.X

		topLeftRow = r + upLeft.Y
		topLeftCol = c + upLeft.X
	)

	downLeftDiagonal := getFromCrossword(crossword, topRightRow, topRightCol, downLeft, 3)
	downRightDiagonal := getFromCrossword(crossword, topLeftRow, topLeftCol, downRight, 3)

	return (downRightDiagonal == "SAM" || downRightDiagonal == "MAS") &&
		(downLeftDiagonal == "SAM" || downLeftDiagonal == "MAS")
}

func getFromCrossword(crossword [][]byte, r, c int, dir Direction, length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		row := r + i*dir.Y
		col := c + i*dir.X
		sb.WriteByte(crossword[row][col])
	}
	return sb.String()
}

func getInput() [][]byte {
	crossword := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		crossword = append(crossword, []byte(s.Text()))
	}
	return crossword
}
