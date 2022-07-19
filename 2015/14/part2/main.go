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

	scores := map[string]int{}
	for i := 0; i < 2503; i++ {
		furthestDist := 0
		for _, r := range reindeer {
			r.Run()
			if r.Distance > furthestDist {
				furthestDist = r.Distance
			}
		}
		for _, r := range reindeer {
			if r.Distance == furthestDist {
				scores[r.Name]++
			}
		}
	}
	bestScore := 0
	for _, v := range scores {
		if v > bestScore {
			bestScore = v
		}
	}
	fmt.Println(bestScore)

}
