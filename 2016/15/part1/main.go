package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disc struct {
	CurrentPosition int
	NumPositions    int
}

func (d *Disc) Rotate() {
	d.CurrentPosition = (d.CurrentPosition + 1) % d.NumPositions
}

func main() {
	discs := getDiscs()
	t := 0
	for !aligned(discs) {
		t++
		rotate(discs)
	}
	fmt.Println(t)
}

func rotate(discs []Disc) {
	for i := range discs {
		discs[i].Rotate()
	}
}

func aligned(discs []Disc) bool {
	for i, disc := range discs {
		if (disc.CurrentPosition+i+1)%disc.NumPositions != 0 {
			return false
		}
	}
	return true
}

func getDiscs() []Disc {
	discs := []Disc{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		discs = append(discs, toDisc(s.Text()))
	}
	return discs
}

func toDisc(s string) Disc {
	segments := strings.Fields(s)
	return Disc{
		CurrentPosition: atoi(strings.Trim(segments[11], ".")),
		NumPositions:    atoi(segments[3]),
	}
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
