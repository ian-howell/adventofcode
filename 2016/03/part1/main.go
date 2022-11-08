package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		a, b, c := getNums(s.Text())
		if isTriangle(a, b, c) {
			count++
		}
	}
	fmt.Println(count)
}

func getNums(s string) (int, int, int) {
	var a, b, c int
	fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
	return a, b, c
}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
