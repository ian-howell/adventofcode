package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	total := 0
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		if !hasIllegalSequence(input) && hasDoubleConsecutives(input) && hasThreeVowels(input) {
			total++
		}
	}
	fmt.Println(total)
}

func hasIllegalSequence(s string) bool {
	return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
}

func hasThreeVowels(s string) bool {
	return strings.Count(s, "a")+strings.Count(s, "e")+strings.Count(s, "i")+strings.Count(s, "o")+strings.Count(s, "u") >= 3
}

func hasDoubleConsecutives(s string) bool {
	if len(s) < 2 {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}
