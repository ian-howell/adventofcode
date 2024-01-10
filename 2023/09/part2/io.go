package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getInput() [][]int {
	values := [][]int{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parts := strings.Fields(s.Text())
		values = append(values, StrSliceToIntSlice(parts))
	}
	return values
}

func StrSliceToIntSlice(strSlice []string) []int {
	atoi := func(a string) int {
		i, _ := strconv.Atoi(a)
		return i
	}

	intSlice := make([]int, 0, len(strSlice))
	for _, str := range strSlice {
		intSlice = append(intSlice, atoi(str))
	}
	return intSlice
}
