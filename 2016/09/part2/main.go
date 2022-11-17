package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(len(expand(s.Text())))
	}
}

func expand(input string) string {
	// This is soooo slow. There's definitely a better way, but since my program was able to find the
	// answer (even though it took about 5 minutes...), I don't feel like I have to figure it out :)
	for strings.Contains(input, "(") {
		var (
			index  int
			output string
			sb     strings.Builder
		)
		for index >= 0 {
			output, index = next(input)
			if index != -1 {
				input = input[index:]
			}
			sb.WriteString(output)
		}
		input = sb.String()
	}
	return input
}

func next(input string) (string, int) {
	if len(input) == 0 {
		return "", -1
	}

	if input[0] != '(' {
		index := strings.Index(input, "(")
		if index != -1 {
			return input[:index], index
		}
		return input, -1
	}

	// This should never fail
	before, after, _ := strings.Cut(input[1:], ")")

	parts := strings.Split(before, "x")
	howManyToRepeat, howManyTimes := atoi(parts[0]), atoi(parts[1])

	output := strings.Repeat(after[:howManyToRepeat], howManyTimes)

	return output, len(before) + howManyToRepeat + 2 // includes the 2 parentheses characters
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
