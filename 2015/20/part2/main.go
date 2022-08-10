package main

import (
	"fmt"
	"math"
)

func main() {
	giftsGiven := map[int]int{}
	for j := 1; ; j++ {
		total := 0
		for v := range getDivisors(j) {
			if giftsGiven[v] < 50 {
				giftsGiven[v]++
				total += 11 * v
			}
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
