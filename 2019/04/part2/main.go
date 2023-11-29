package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var rawInput string
	fmt.Scan(&rawInput)
	hyphenIndex := strings.IndexByte(rawInput, '-')

	lb := atoi(rawInput[:hyphenIndex])
	ub := atoi(rawInput[hyphenIndex+1:])

	var count int
	for i := lb; i <= ub; i++ {
		// Notably, this could be sped up by combining the 2 functions.
		// However, this is fast enough that I can't even properly profile it, so it's not really
		// worth my time.
		if weirdoRule(i) && isNonDecreasing(i) {
			count++
			// fmt.Println("X  ", i)
		} else {
			// fmt.Println("   ", i)
		}
	}
	fmt.Println(count)
}

func isNonDecreasing(x int) bool {
	for x >= 11 {
		if x%10 < (x/10)%10 {
			return false
		}
		x /= 10
	}
	return true
}

func weirdoRule(x int) bool {
	last := x % 10
	inARow := 1
	for x > 0 {
		x /= 10
		if x%10 == last {
			inARow++
		} else {
			if inARow == 2 {
				return true
			}
			// Otherwise, restart the search
			inARow = 1
			last = x % 10
		}
	}
	return false
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}