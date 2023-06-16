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
	n := x % 10
	for x > 0 {
		x /= 10
		if n == x%10 {
			return true
		}
		n = x % 10
	}
	return false
}
