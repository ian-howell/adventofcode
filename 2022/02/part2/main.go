package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	score := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		theirs, mine := parts[0], parts[1]
		score += rockPaperSciccors(theirs, mine)
	}
	fmt.Println(score)
}

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"

	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

func rockPaperSciccors(a, b string) int {
	switch a {
	case Rock:
		return rockScore(b)
	case Paper:
		return paperScore(b)
	case Scissors:
		return scissorsScore(b)
	}
	return 0
}

func rockScore(a string) int {
	switch a {
	case Win:
		return 6 + 2
	case Draw:
		return 3 + 1
	case Lose:
		return 0 + 3
	}
	return 0
}

func paperScore(a string) int {
	switch a {
	case Win:
		return 6 + 3
	case Draw:
		return 3 + 2
	case Lose:
		return 0 + 1
	}
	return 0
}

func scissorsScore(a string) int {
	switch a {
	case Win:
		return 6 + 1
	case Draw:
		return 3 + 3
	case Lose:
		return 0 + 2
	}
	return 0
}
