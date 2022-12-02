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
		theirs, mine := toRockPaperScissor(parts[0]), toRockPaperScissor(parts[1])
		score += rockPaperSciccors(theirs, mine)
	}
	fmt.Println(score)
}

const (
	Rock = iota + 1
	Paper
	Scissors
)

func toRockPaperScissor(a string) int {
	switch a {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	return 0
}

func rockPaperSciccors(a, b int) int {
	switch b {
	case Rock:
		return rockScore(a)
	case Paper:
		return paperScore(a)
	case Scissors:
		return scissorsScore(a)
	}
	return 0
}

func rockScore(a int) int {
	score := 1
	switch a {
	case Rock:
		return score + 3
	case Paper:
		return score + 0
	case Scissors:
		return score + 6
	}
	return 0
}

func paperScore(a int) int {
	score := 2
	switch a {
	case Rock:
		return score + 6
	case Paper:
		return score + 3
	case Scissors:
		return score + 0
	}
	return 0
}

func scissorsScore(a int) int {
	score := 3
	switch a {
	case Rock:
		return score + 0
	case Paper:
		return score + 6
	case Scissors:
		return score + 3
	}
	return 0
}
