package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := getInput()
	combos, _ := getCombinations(150, input, 0)
	minRequired := -1
	for _, combo := range combos {
		if minRequired == -1 || len(combo) < minRequired {
			minRequired = len(combo)
		}
	}
	count := 0
	for _, combo := range combos {
		if len(combo) == minRequired {
			count++
		}
	}
	fmt.Println(count)
}

func getInput() []int {
	vals := []int{}
	var t int
	for _, err := fmt.Scan(&t); err == nil; _, err = fmt.Scan(&t) {
		vals = append(vals, t)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(vals)))
	return vals
}

func getCombinations(toDistribute int, buckets []int, depth int) ([][]int, bool) {
	// printf(depth, "toDistribute=%d, buckets=%v", toDistribute, buckets)
	if toDistribute == 0 {
		// printf(depth, "No more to distribute!")
		return nil, true
	}
	if toDistribute < 0 || len(buckets) == 0 {
		if toDistribute < 0 {
			// printf(depth, "less than 0 to distribute. Shouldn't be here?")
		} else if len(buckets) == 0 {
			// printf(depth, "Ran out of buckets")
		}
		return nil, false
	}

	combos := [][]int{}
	for i := 0; i < len(buckets); i++ {
		// printf(depth, "comparing toDistribute=%d to buckets[%d]=%d", toDistribute, i, buckets[i])
		if toDistribute >= buckets[i] {
			subCombos, valid := getCombinations(toDistribute-buckets[i], buckets[i+1:], depth+2)
			// printf(depth, "subCombos:%v, valid:%v", subCombos, valid)
			if valid {
				if len(subCombos) == 0 {
					combos = append(combos, []int{buckets[i]})
				}
				for _, subCombo := range subCombos {
					thisCombo := append([]int{buckets[i]}, subCombo...)
					combos = append(combos, thisCombo)
				}
				// printf(depth, "%v", combos)
			} else {
				// printf(depth, "INVALID")
			}
		}
	}

	return combos, len(combos) > 0
}

func printf(depth int, f string, vars ...any) {
	indent := strings.Repeat(" ", depth)
	indented := fmt.Sprintf("%s%s\n", indent, f)
	fmt.Printf(indented, vars...)
}
