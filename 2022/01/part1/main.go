package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf []int

func main() {
	elves := getElves()
	most := 0
	for _, elf := range elves {
		total := elf.Total()
		if total > most {
			most = total
		}
	}
	fmt.Println(most)
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
