package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	for i := 0; i < 40; i++ {
		input = step(input)
	}
	fmt.Println(len(input))
}

func step(s string) string {
	l := len(s)
	if l == 1 {
		return fmt.Sprintf("1%v", s[0]-'0')
	}
	result := strings.Builder{}
	i := 0
	for j := 1; j < len(s); j++ {
		if s[i] != s[j] {
			result.WriteString(fmt.Sprintf("%v%v", j-i, s[i]-'0'))
			i = j
		}
	}
	result.WriteString(fmt.Sprintf("%v%v", l-i, s[l-1]-'0'))
	return result.String()
}
