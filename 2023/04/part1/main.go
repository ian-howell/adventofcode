package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total += getScore(s.Text())
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
	if matches > 0 {
		return pow2(matches - 1)
	}
	return 0
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

func pow2(x int) int {
	return int(math.Pow(2, float64(x)))
}
