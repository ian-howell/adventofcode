package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

const Debug = false

type Slice []interface{}

func (s Slice) Len() int { return len(s) }
func (s Slice) Less(i, j int) bool {
	result := Compare(0, s[i], s[j])
	return result == Lesser
}
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	total := 1
	div1 := []interface{}{[]interface{}{2.}}
	div2 := []interface{}{[]interface{}{6.}}
	sdiv1 := fmt.Sprintf("%v", div1)
	sdiv2 := fmt.Sprintf("%v", div2)
	signals := append(getInput(), div1, div2)
	sort.Sort(Slice(signals))
	for i, signal := range signals {
		fmt.Println(i, signal)
		s := fmt.Sprintf("%v", signal)
		if s == sdiv1 || s == sdiv2 {
			total *= (i + 1)
		}
	}
	fmt.Println(total)
}

func getInput() []interface{} {
	results := []interface{}{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var data interface{}

		err := json.Unmarshal(s.Bytes(), &data)
		if err != nil {
			panic(err)
		}
		results = append(results, []interface{}{data})

		s.Scan()
		err = json.Unmarshal(s.Bytes(), &data)
		if err != nil {
			panic(err)
		}

		results = append(results, []interface{}{data})

		// eat the empty line
		s.Scan()
	}
	return results
}

const (
	Lesser  = -1
	Greater = 1
	Equal   = 0
)

func Compare(depth int, left, right interface{}) int {
	RecursivePrintf(depth, "Comparing %v to %v\n", left, right)
	switch l := left.(type) {
	case float64:
		// RecursivePrintf(depth, "left is number (%v)... ", l)
		switch r := right.(type) {
		case float64:
			// RecursivePrintf(depth, "right is number (%v)...\n", r)
			return CompareInts(int(l), int(r))
		case []interface{}:
			RecursivePrintf(depth, "converting %v to list\n", l)
			return Compare(depth, []interface{}{l}, r)
		default:
			panic(fmt.Sprintf("Right was %T\n", r))
		}
	case []interface{}:
		switch r := right.(type) {
		case float64:
			// RecursivePrintf(depth, "right is number (%v)...\n", r)
			return Compare(depth, l, []interface{}{r})
		case []interface{}:
			return CompareLists(depth, l, r)
		default:
			panic(fmt.Sprintf("Right was %T\n", r))
		}
	default:
		panic(fmt.Sprintf("Left was %T\n", l))
	}
}

func CompareInts(left, right int) int {
	switch {
	case left > right:
		return Greater
	case left < right:
		return Lesser
	}
	return Equal
}

func CompareLists(depth int, left, right []interface{}) int {
	for i := 0; i < len(left) || i < len(right); i++ {
		if i >= len(left) {
			return Lesser
		}
		if i >= len(right) {
			return Greater
		}
		result := Compare(depth+1, left[i], right[i])
		if result != Equal {
			return result
		}
	}

	return Equal
}

func RecursivePrintf(depth int, format string, args ...any) {
	if Debug {
		fmt.Print(strings.Repeat("  ", depth))
		fmt.Printf(format, args...)
	}
}
