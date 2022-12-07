package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		p1, p2 := getPairs(s.Text())
		if p1.Overlaps(p2) {
			total++
		}
	}
	fmt.Println(total)
}

type Pair struct {
	Lower int
	Upper int
}

func NewPair(s string) Pair {
	parts := strings.Split(s, "-")
	return Pair{
		Lower: atoi(parts[0]),
		Upper: atoi(parts[1]),
	}
}

func (p Pair) Overlaps(o Pair) bool {
	return p.Lower <= o.Upper && p.Upper >= o.Lower
}

func getPairs(s string) (Pair, Pair) {
	parts := strings.Split(s, ",")
	return NewPair(parts[0]), NewPair(parts[1])
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
