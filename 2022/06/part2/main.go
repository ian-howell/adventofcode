package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		markerIndex := findMarker(s.Text())
		fmt.Println(markerIndex)
	}
}

func findMarker(s string) int {
	counts := map[byte]int{}
	for i := 0; i < 14; i++ {
		counts[s[i]]++
	}

	for i := 14; i < len(s); i++ {
		if len(counts) == 14 {
			return i
		}

		counts[s[i-14]]--
		if counts[s[i-14]] == 0 {
			delete(counts, s[i-14])
		}

		counts[s[i]]++
	}

	return -1
}
