package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ingredients := []*Ingredient{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ingredients = append(ingredients, New(s.Text()))
	}

	best := 0
	for _, combo := range getCombinations(100, 4) {
		current := 1
		for i := 0; i <= 3; i++ {
			property := Combine(ingredients, combo, i)
			if property <= 0 {
				current = 0
				break
			}
			current *= property
		}
		if current > best {
			best = current
		}
	}
	fmt.Println(best)
}

func getCombinations(toDistribute, slots int) [][]int {
	if toDistribute == 0 {
		return [][]int{make([]int, slots)}
	}
	if slots == 1 {
		return [][]int{{toDistribute}}
	}

	combos := [][]int{}
	for i := toDistribute; i >= 0; i-- {
		sub_combos := getCombinations(toDistribute-i, slots-1)
		for _, sub_combo := range sub_combos {
			this_combo := append([]int{i}, sub_combo...)
			combos = append(combos, this_combo)
		}
	}

	return combos
}

func Combine(ingredients []*Ingredient, distribution []int, target int) int {
	total := 0
	for i, g := range ingredients {
		total += distribution[i] * g.Properties[target]
	}
	return total
}

// (44*-1 + 56*2) * (44*-2 + 56*3) * (44*6 + 56*-2) * (44*3 + 56*-1)
// (A*-1 + B*2) * (A*-2 + B*3) * (A*6 + B*-2) * (A*3 + B*-1)
// (-A + 2B) * (-2A + 3B) * (6A + -2B) * (3A + -B)
//
