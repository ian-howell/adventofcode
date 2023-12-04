package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scoreCards := []string{}
	scoreCardCounts := []int{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		scoreCards = append(scoreCards, s.Text())
		scoreCardCounts = append(scoreCardCounts, 1)
	}

	total := 0
	for i := 0; i < len(scoreCards); i++ {
		total += scoreCardCounts[i]
		score := getScore(scoreCards[i])
		for j := i + 1; (j < len(scoreCards)) && (j < i+1+score); j++ {
			scoreCardCounts[j] += scoreCardCounts[i]
		}
		// fmt.Println(i, scoreCardCounts)
	}

	fmt.Println(total)
}

func getScore(scoreCard string) int {
	_, scoreCard, _ = strings.Cut(scoreCard, ": ")
	winningNumStr, playerNumsStr, _ := strings.Cut(scoreCard, " | ")
	winningNums := toSet(winningNumStr)
	matches := 0
	for _, playerNum := range strings.Fields(playerNumsStr) {
		if _, ok := winningNums[atoi(playerNum)]; ok {
			matches++
		}
	}
	return matches
}

func toSet(s string) map[int]struct{} {
	set := map[int]struct{}{}
	for _, val := range strings.Fields(s) {
		set[atoi(val)] = struct{}{}
	}
	return set
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
