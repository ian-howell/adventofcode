package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := []int{1}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		if len(parts) > 1 {
			Addx(&x, atoi(parts[1]))
		} else {
			Noop(&x)
		}
	}

	// for cycle, val := range x {
	// 	fmt.Printf("% 3d: %v\n", cycle+1, val)
	// }

	Render(x)
}

func Addx(x *[]int, val int) {
	last := (*x)[len(*x)-1]
	*x = append(*x, last, last+val)
}

func Noop(x *[]int) {
	*x = append(*x, (*x)[len(*x)-1])
}

func Render(x []int) {
	i := 0
	for r := 0; r < 6; r++ {
		for c := 0; c < 40; c++ {
			if abs(c-x[i]) <= 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
			i++
		}
		fmt.Println()
	}
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
