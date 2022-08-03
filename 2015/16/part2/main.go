package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	auntSues := []*AuntSue{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		_, input, _ := strings.Cut(s.Text(), ": ")
		auntSues = append(auntSues, NewAuntSue(input))
	}

	for i, auntSue := range auntSues {
		if Filter(auntSue) {
			fmt.Println(i + 1)
			// The input should only have one answer, so there's no need to keep working after we've found it
			break
		}
	}
}
