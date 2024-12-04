package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	processed := preprocess(input)
	answer := calculate(processed)
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

func preprocess(s string) string {
	enabled := true
	sb := strings.Builder{}
	for s != "" {
		if enabled {
			var before string
			before, s, _ = strings.Cut(s, "don't()")
			sb.WriteString(before)
		} else {
			_, s, _ = strings.Cut(s, "do()")
		}
		enabled = !enabled
	}
	return sb.String()
}

func calculate(s string) int {
	total := 0
	vals, remainder, found := nextMulVals(s)
	for ; found; vals, remainder, found = nextMulVals(remainder) {
		product, err := mul(vals)
		if err != nil {
			log.Printf("skipping `%s`: %v", vals, err)
			continue
		}

		total += product
	}
	return total
}

func mul(vals string) (product int, err error) {
	x, y, err := parseVals(vals)
	if err != nil {
		return 0, fmt.Errorf("failed to parse values: %w", err)
	}

	// check that both ints are in the range [0, 1000)
	between := func(lb, x, ub int) bool { return lb <= x && x < ub }
	if !between(0, x, 1000) || !between(0, y, 1000) {
		return 0, fmt.Errorf("values must be in the range [0,1000), got (x = %d, y = %d)", x, y)
	}

	return x * y, nil
}

// nextMulVals finds the next instance of `mul($vals)` and returns the value of `$vals`. If there are no more
// instances, found will be false. When found the remainder is the text that follows the closing parnethesis.
func nextMulVals(s string) (vals string, remainder string, found bool) {
	remainder = s
	for remainder != "" {
		var upToClosingParen string
		upToClosingParen, remainder, found = strings.Cut(remainder, ")")
		if !found {
			// there's no more closing parentheses, so there can't be any more mul(x,y)
			return "", "", false
		}

		// Work backward from the closing paren
		_, stuffInsideParens, found := cutReverse(upToClosingParen, "mul(")
		if found {
			return stuffInsideParens, remainder, true
		}
	}

	// We made it all the way through s and never found a mul(x,y)
	return "", "", false
}

// cutReverse is the same as strings.Cut, but works from the back of the string
func cutReverse(s string, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

// parseVals takes a string s of the form "x,y", where x and y are integers. It returns the integers. err will
// be non-nil if there is not exactly one comma, or if either x or y are not integers.
func parseVals(s string) (x, y int, err error) {
	vals := strings.Split(s, ",")
	if len(vals) != 2 {
		return 0, 0, fmt.Errorf("expected string `%s` to contain exactly one comma", s)
	}

	x, err = strconv.Atoi(vals[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert `%s` to int: %w", vals[0], err)
	}

	y, err = strconv.Atoi(vals[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert `%s` to int: %w", vals[0], err)
	}

	return x, y, nil
}
