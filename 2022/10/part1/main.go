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

	total := 0
	for i, val := range x {
		switch i {
		case 19, 59, 99, 139, 179, 219:
			total += val * (i + 1)
			// fmt.Println(i+1, val, val*(i+1))
		}
	}
	fmt.Println(total)
}

func Addx(x *[]int, val int) {
	last := (*x)[len(*x)-1]
	*x = append(*x, last, last+val)
}

func Noop(x *[]int) {
	*x = append(*x, (*x)[len(*x)-1])
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
