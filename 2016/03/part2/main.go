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
		a1, a2, a3 := getNums(s.Text())
		s.Scan()
		b1, b2, b3 := getNums(s.Text())
		s.Scan()
		c1, c2, c3 := getNums(s.Text())
		if isTriangle(a1, b1, c1) {
			count++
		}
		if isTriangle(a2, b2, c2) {
			count++
		}
		if isTriangle(a3, b3, c3) {
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
