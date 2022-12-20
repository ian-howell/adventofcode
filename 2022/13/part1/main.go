package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const Debug = false

func main() {
	total := 0
	for i, input := range getInput() {
		left, right := input[0], input[1]
		switch Compare(0, left, right) {
		case Greater:
			RecursivePrintf(0, "%v is greater that %v\n", left, right)
		case Lesser:
			RecursivePrintf(0, "%v is lesser that %v\n", left, right)
			total += i + 1
		case Equal:
			RecursivePrintf(0, "%v is equal to %v\n", left, right)
		}
	}
	fmt.Println(total)
}

func getInput() [][2]interface{} {
	results := [][2]interface{}{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var left, right interface{}

		// left is already loaded into the scanner
		err := json.Unmarshal(s.Bytes(), &left)
		if err != nil {
			panic(err)
		}

		// load right into scanner
		s.Scan()
		err = json.Unmarshal(s.Bytes(), &right)
		if err != nil {
			panic(err)
		}

		results = append(results, [2]interface{}{left, right})

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
