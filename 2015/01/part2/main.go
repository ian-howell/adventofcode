package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	currentFloor := 1
	i := 1
	for _, c := range input {
		currentFloor += map[rune]int{')': -1, '(': 1}[c]
		if currentFloor == 0 {
			break
		}
		i++
	}
	fmt.Println(i)
}
