package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	wire1 := map[[2]int]int{}
	var rawInput string
	fmt.Scan(&rawInput)

	dirMap := map[byte][2]int{
		'U': {1, 0},
		'D': {-1, 0},
		'L': {0, -1},
		'R': {0, 1},
	}

	row, col, dist := 0, 0, 0
	for _, step := range strings.Split(rawInput, ",") {
		dir := step[0]
		for i := 0; i < atoi(step[1:]); i++ {
			row += dirMap[dir][0]
			col += dirMap[dir][1]
			dist++
			if _, ok := wire1[[2]int{row, col}]; !ok {
				wire1[[2]int{row, col}] = dist
			}
		}
	}

	fmt.Scan(&rawInput)
	row, col, dist = 0, 0, 0
	totalDist := -1
	for _, step := range strings.Split(rawInput, ",") {
		dir := step[0]
		for i := 0; i < atoi(step[1:]); i++ {
			row += dirMap[dir][0]
			col += dirMap[dir][1]
			dist++
			if wire1Dist, ok := wire1[[2]int{row, col}]; ok {
				if totalDist == -1 || wire1Dist+dist < totalDist {
					totalDist = wire1Dist + dist
				}
			}
		}
	}

	fmt.Println(totalDist)
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
