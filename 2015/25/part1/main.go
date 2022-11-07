package main

import "fmt"

func main() {
	index := getIndex(2981, 3075)
	code := 20151125
	for i := 2; i <= index; i++ {
		code = nextCode(code)
	}
	fmt.Println(code)
}

func getIndex(r, c int) int {
	switch {
	case r == 1 && c == 1:
		return 1
	case c == 1:
		return getIndex(r-1, c) + r - 1
	default:
		return getIndex(r, c-1) + r + c - 1
	}
}

func nextCode(lastCode int) int {
	return (lastCode * 252533) % 33554393
}
