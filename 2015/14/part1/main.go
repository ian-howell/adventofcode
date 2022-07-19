package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reindeer := []*Reindeer{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		reindeer = append(reindeer, New(s.Text()))
	}

	largest := 0
	for _, r := range reindeer {
		current := r.Run(2503)
		if current > largest {
			largest = current
		}
	}
	fmt.Println(largest)

}
