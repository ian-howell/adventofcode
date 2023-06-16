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

func fuelForMass(mass int) int {
	return mass/3 - 2
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
