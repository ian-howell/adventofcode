package main

import (
	"fmt"
)

func main() {
	var input string
	total := 0
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		if rule1(input) && rule2(input) {
			total++
		}
	}
	fmt.Println(total)
}

func rule1(s string) bool {
	if len(s) < 4 {
		return false
	}

	for i := 0; i < len(s)-3; i++ {
		for j := i+2; j < len(s)-1; j++ {
			if s[j] == s[i] && s[j+1] == s[i+1] {
				return true
			}
		}
	}
	return false
}

func rule2(s string) bool {
	if len(s) < 3 {
		return false
	}

	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}
