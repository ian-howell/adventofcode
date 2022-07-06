package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	total := 0
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		total += calculateRequiredRibbon(input)
	}

	fmt.Println(total)
}

func calculateRequiredRibbon(dimensions string) int {
	dim := strings.Split(dimensions, "x")
	h, _ := strconv.Atoi(dim[0])
	l, _ := strconv.Atoi(dim[1])
	w, _ := strconv.Atoi(dim[2])

	dimInts := []int{h, l, w}
	sort.Ints(dimInts)

	return h*l*w + 2*(dimInts[0]+dimInts[1])
}
