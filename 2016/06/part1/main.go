package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// For sample
// const passwordLen = 6

const passwordLen = 8

func main() {
	freq := make([]map[byte]int, passwordLen)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for i := 0; i < passwordLen; i++ {
			if freq[i] == nil {
				freq[i] = map[byte]int{}
			}
			freq[i][s.Text()[i]]++
		}
	}

	sb := strings.Builder{}
	for i := 0; i < passwordLen; i++ {
		largest := 0
		var c byte

		for letter := range freq[i] {
			if freq[i][letter] > largest {
				largest = freq[i][letter]
				c = letter
			}
		}

		sb.WriteByte(c)
	}
	fmt.Println(sb.String())
}
