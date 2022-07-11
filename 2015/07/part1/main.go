package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules := map[string]string{}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := s.Text()
		parts := strings.Split(input, " -> ")
		rules[parts[1]] = string(parts[0])
	}

	fmt.Println(dfs("a", rules))
}

func dfs(target string, rules map[string]string) int {
	// target is a number
	// this happens for rules like [le AND 1 -> x]
	if val, err := strconv.Atoi(target); err == nil {
		return val
	}

	// number -> target
	if val, err := strconv.Atoi(rules[target]); err == nil {
		return val
	}

	parts := strings.Split(rules[target], " ")
	val := 0
	switch {
	case len(strings.Split(rules[target], " ")) == 1:
		// non-terminal -> target
		val = dfs(rules[target], rules)
	case strings.Contains(rules[target], "NOT"):
		val = ^dfs(parts[1], rules)
	case strings.Contains(rules[target], "OR"):
		val = dfs(parts[0], rules) | dfs(parts[2], rules)
	case strings.Contains(rules[target], "AND"):
		val = dfs(parts[0], rules) & dfs(parts[2], rules)
	case strings.Contains(rules[target], "RSHIFT"):
		shift, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		val = dfs(parts[0], rules) >> shift
	case strings.Contains(rules[target], "LSHIFT"):
		shift, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		val = dfs(parts[0], rules) << shift
	default:
		panic("Nothing matched...")
	}

	// finally, simplify this rule to [number -> target]
	rules[target] = strconv.Itoa(val)
	return val
}
