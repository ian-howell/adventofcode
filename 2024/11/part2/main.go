package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const numBlinks = 75

func main() {
	histogram := getInput()
	for range numBlinks {
		histogram = blink(histogram)
	}

	total := 0
	for _, freq := range histogram {
		total += freq
	}
	fmt.Println(len(histogram), total)
}

func blink(histogram map[int]int) map[int]int {
	newHistogram := map[int]int{}
	for val, freq := range histogram {
		switch {
		case val == 0:
			newHistogram[1] += freq
		case getNumDigits(val)%2 == 0:
			l, r := split(val)
			newHistogram[l] += freq
			newHistogram[r] += freq
		default:
			newHistogram[2024*val] += freq
		}
	}
	return newHistogram
}

func split(x int) (l, r int) {
	str := strconv.Itoa(x)
	middle := len(str) / 2
	return atoi(str[:middle]), atoi(str[middle:])
}

func getNumDigits(x int) int {
	return int(math.Log10(float64(x))) + 1
}

func getInput() map[int]int {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	histogram := map[int]int{}
	for _, val := range toInts(s.Text()) {
		histogram[val]++
	}
	return histogram
}

func toInts(s string) []int {
	strs := strings.Fields(s)
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		ints = append(ints, atoi(str))
	}
	return ints
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
