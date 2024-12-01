package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	left, right := readInput()
	fmt.Println(nLgnStrategy(left, right))

	// This does not work!
	// fmt.Println(nStrategy(left, right))
}

// The nStrategy is to take the difference between the sums of both lists
func nStrategy(l, r []int) int {
	lSum, rSum := 0, 0
	for i := range l {
		lSum += l[i]
		rSum += r[i]
	}
	return abs(lSum - rSum)
}

// The nLgnStrategy is to sort both lists, and then iterate over the lists, aggregating the differences
// between each pair
func nLgnStrategy(l, r []int) int {
	slices.Sort(l)
	slices.Sort(r)

	sum := 0
	for i := 0; i < len(l); i++ {
		sum += abs(l[i] - r[i])
	}
	return sum
}

func readInput() ([]int, []int) {
	var l, r []int

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		// This panics for bad inputs (e.g. len(fields) == 0 and non-integer fields)
		l = append(l, atoi(fields[0]))
		r = append(r, atoi(fields[1]))
	}

	return l, r
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
