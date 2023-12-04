package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total += getPower(s.Text())
	}
	fmt.Println(total)
}

func getPower(s string) int {
	_, games, _ := strings.Cut(s, ": ")
	numCubes := map[string]int{}
	for _, game := range strings.Split(games, "; ") {
		for _, handful := range strings.Split(game, ", ") {
			numStr, color, _ := strings.Cut(handful, " ")
			num := atoi(numStr)
			if num > numCubes[color] {
				numCubes[color] = num
			}
		}
	}
	total := 1
	for _, val := range numCubes {
		total *= val
	}
	return total
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
