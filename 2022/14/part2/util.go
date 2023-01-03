package main

import (
	"strconv"
	"strings"
)

func str2point(s string) Point {
	parts := strings.Split(s, ",")
	return Point{
		Row: atoi(parts[1]),
		Col: atoi(parts[0]),
	}
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func getDiff(a, b int) int {
	if a == b {
		return 0
	}
	d := b - a
	return d / abs(d)
}
