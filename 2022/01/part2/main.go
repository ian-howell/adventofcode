package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf []int

func main() {
	elves := getElves()

	// Just do the slow, naive thing: sort the array and grab the last three
	var totals []int
	for _, elf := range elves {
		totals = append(totals, elf.Total())
	}
	sort.Ints(totals)

	total := 0
	for _, calories := range totals[len(totals)-3:] {
		total += calories
	}
	fmt.Println(total)
}

func getElves() []Elf {
	var elves []Elf
	var elf Elf

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			elves = append(elves, elf)
			elf = Elf{}
		} else {
			elf = append(elf, atoi(s.Text()))
		}
	}

	return append(elves, elf)
}

func (e Elf) Total() int {
	total := 0
	for _, v := range e {
		total += v
	}
	return total
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
