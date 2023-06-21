package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	wire1 := map[[2]int]struct{}{}
	var rawInput string
	fmt.Scan(&rawInput)

	dirMap := map[byte][2]int{
		'U': {1, 0},
		'D': {-1, 0},
		'L': {0, -1},
		'R': {0, 1},
	}

	row, col := 0, 0
	for _, step := range strings.Split(rawInput, ",") {
		dir := step[0]
		for i := 0; i < atoi(step[1:]); i++ {
			row += dirMap[dir][0]
			col += dirMap[dir][1]
			wire1[[2]int{row, col}] = struct{}{}
		}
	}

	fmt.Scan(&rawInput)
	row, col = 0, 0
	closestIntersection := -1
	for _, step := range strings.Split(rawInput, ",") {
		dir := step[0]
		for i := 0; i < atoi(step[1:]); i++ {
			row += dirMap[dir][0]
			col += dirMap[dir][1]
			if _, ok := wire1[[2]int{row, col}]; ok {
				manhattan := abs(row) + abs(col)
				if closestIntersection == -1 || manhattan < closestIntersection {
					closestIntersection = manhattan
				}
			}
		}
	}

	fmt.Println(closestIntersection)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
