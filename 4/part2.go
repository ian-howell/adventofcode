package main

import (
	"fmt"
)

func main() {
	var lb, ub int
	fmt.Scanf("%d %d", &lb, &ub)
	count := 0
	for i := lb; i <= ub; i++ {
		if valid(i) {
			count++
		}
	}

	fmt.Println(count)
}

func valid(x int) bool {
	return (x >= 100000 && x <= 999999) &&
		!decreases(x) &&
		containsDouble(x)
}

func decreases(x int) bool {
	n := x % 10
	for x > 0 {
		x /= 10
		if n < x%10 {
			return true
		}
		n = x % 10
	}
	return false
}

func containsDouble(x int) bool {
	state := 1
	stateTable := map[int]map[bool]int{
		1: {true: 2, false: 1},
		2: {true: 3, false: 4},
		3: {true: 3, false: 1},
	}
	for x > 9 {
		n := x % 10
		x /= 10
		state = stateTable[state][x%10 == n]
		if state == 4 {
			return true
		}
	}
	return state == 2
}
