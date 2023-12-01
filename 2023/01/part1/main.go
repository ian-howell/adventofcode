package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		total += findCalibrationValue(s.Text())
	}
	fmt.Println(total)
}

func findCalibrationValue(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		if val, err := strconv.Atoi(string(s[i])); err == nil {
			total = 10 * val
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if val, err := strconv.Atoi(string(s[i])); err == nil {
			total += val
			break
		}
	}

	return total
}
