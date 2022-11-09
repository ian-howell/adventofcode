package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if supportsTLS(s.Text()) {
			count++
		}
	}
	fmt.Println(count)
}

func supportsTLS(s string) bool {
	inside, outside := separateStrings(s)

	for _, in := range inside {
		if hasABBA(in) {
			return false
		}
	}

	for _, out := range outside {
		if hasABBA(out) {
			return true
		}
	}

	return false
}

func separateStrings(s string) (inside []string, outside []string) {
	for len(s) > 0 {
		before, after, found := strings.Cut(s, "[")
		outside = append(outside, before)

		if !found {
			// We're done
			break
		}

		s = strings.TrimLeft(after, "[")

		// There will always be a matching ], so no need to check if it was found
		before, s, _ = strings.Cut(s, "]")
		inside = append(inside, before)
	}
	return inside, outside
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if isABBA(s[i : i+4]) {
			return true
		}
	}
	return false
}

func isABBA(s string) bool {
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}
