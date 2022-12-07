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
	for i := 0; i < 4; i++ {
		counts[s[i]]++
	}

	for i := 4; i < len(s); i++ {
		if len(counts) == 4 {
			return i
		}

		counts[s[i-4]]--
		if counts[s[i-4]] == 0 {
			delete(counts, s[i-4])
		}

		counts[s[i]]++
	}

	return -1
}
