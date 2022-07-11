package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := s.Text()
		total += encodedLen(input) - len(input)
	}
	fmt.Println(total)
}

func encodedLen(s string) int {
	// count starts at 2 because the encoded string will at least be ""
	count := 2
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '"' || s[i] == '\\':
			count += 2
		case s[i] != '\\':
			count++
		}
	}
	return count
}
