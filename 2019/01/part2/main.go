package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total += fuelForMass(atoi(s.Text()))
	}
	fmt.Println(total)
}

func fuelForMass(mass int) (fuel int) {
	for mass > 0 {
		required := mass/3 - 2
		fuel += required
		mass = required
	}
	return
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
