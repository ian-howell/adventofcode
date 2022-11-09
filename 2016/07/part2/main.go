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
		if supportsSSL(s.Text()) {
			count++
		}
	}
	fmt.Println(count)
}

func supportsSSL(s string) bool {
	inside, outside := separateStrings(s)

	for _, out := range outside {
		for i := 0; i < len(out)-2; i++ {
			if isABA(out[i : i+3]) {
				bab := correspondingBAB(out[i : i+3])
				for _, in := range inside {
					if strings.Contains(in, bab) {
						return true
					}
				}
			}
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

func isABA(s string) bool {
	return !strings.ContainsAny(s, "[]") && s[0] == s[2] && s[0] != s[1]
}

func correspondingBAB(aba string) string {
	sb := strings.Builder{}
	sb.WriteByte(aba[1])
	sb.WriteByte(aba[0])
	sb.WriteByte(aba[1])
	return sb.String()
}
