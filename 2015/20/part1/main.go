package main

import (
	"fmt"
	"math"
)

func main() {
	// There's something wrong about this assumption...
	// total := 0
	// i := 1
	// for ; ; i *= 2 {
	// 	total += 10 * i
	// 	if total >= 33100000 {
	// 		break
	// 	}
	// }

	// for j := i / 2; j < i; j++ {
	for j := 2; ; j++ {
		total := 0
		for v := range getDivisors(j) {
			total += 10 * v
		}
		if total >= 33100000 {
			fmt.Println(j)
			break
		}
	}
}

func getDivisors(x int) map[int]struct{} {
	divisors := map[int]struct{}{1: {}, x: {}}
	for i := 2; float64(i) <= math.Sqrt(float64(x)); i++ {
		if x%i == 0 {
			divisors[i] = struct{}{}
			divisors[x/i] = struct{}{}
		}
	}
	return divisors
}
