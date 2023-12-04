package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	total := 0
	gameNo := 1
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if valid(s.Text()) {
			total += gameNo
		}
		gameNo++
	}
	fmt.Println(total)
}

func valid(s string) bool {
	_, games, _ := strings.Cut(s, ": ")
	for _, game := range strings.Split(games, "; ") {
		for _, handful := range strings.Split(game, ", ") {
			numStr, color, _ := strings.Cut(handful, " ")
			num := atoi(numStr)
			if num > numCubes[color] {
				return false
			}
		}
	}
	return true
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
