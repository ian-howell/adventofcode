package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	answer := calculate(input)
	fmt.Println(answer)
}

func getInput() string {
	sb := strings.Builder{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		sb.WriteString(fmt.Sprintln(s.Text()))
	}
	return sb.String()
}

func calculate(s string) int {
	total := 0
	product, remainder, found := evaluateNextMul(s)
	for found {
		total += product
		product, remainder, found = evaluateNextMul(remainder)
	}
	return total
}

func evaluateNextMul(s string) (product int, remainder string, found bool) {
	remainder = s
	for {
		var insideParens string
		insideParens, remainder, found = getNextVals(remainder)
		if !found {
			return 0, "", false
		}

		x, y, ok := parseVals(insideParens)
		if !ok {
			continue
		}

		// check that both ints are in the range [0, 1000)
		between := func(lb, x, ub int) bool { return lb <= x && x < ub }
		if !between(0, x, 1000) || !between(0, y, 1000) {
			continue
		}

		return x * y, remainder, true
	}
}

// getNextVals finds the next instance of `mul($vals)` and returns the value of `$vals`. If there are no more
// instances, found will be false. In order to handle nested muls, the remainder will begin with the first
// character following "mul("
func getNextVals(s string) (vals string, remainder string, found bool) {
	_, remainder, found = strings.Cut(s, "mul(")
	if !found {
		return "", "", false
	}

	// NOTE: Don't overwrite the remainder here!
	// Consider this edge case:
	//   mul(mul(172,611))
	vals, _, found = strings.Cut(remainder, ")")
	if !found {
		return "", "", false
	}

	return vals, remainder, true
}

// parseVals takes a string s of the form "x,y", where x and y are integers. It returns the integers. ok will
// be false if there is not exactly one comma, or if either x or y are not integers.
func parseVals(s string) (x int, y int, ok bool) {
	vals := strings.Split(s, ",")
	if len(vals) != 2 {
		return 0, 0, false
	}

	var err error
	x, err = strconv.Atoi(vals[0])
	if err != nil {
		return 0, 0, false
	}

	y, err = strconv.Atoi(vals[1])
	if err != nil {
		return 0, 0, false
	}

	return x, y, true
}
