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
		total += len(input) - inMemLen(input)
	}
	fmt.Println(total)
}

func inMemLen(s string) int {
	s = s[1 : len(s)-1]
	i := 0
	count := 0
	for i < len(s) {
		switch {
		case s[i] != '\\':
			i++
			count++
		case s[i+1] == '"' || s[i+1] == '\\':
			i += 2
			count++
		case s[i+1] == 'x':
			i += 4
			count++
		}
	}
	return count
}
