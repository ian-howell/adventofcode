package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	input = preprocess(input)

	for {
		input = next(input)
		if isValid(input) {
			break
		}
	}

	// Just do it again...
	for {
		input = next(input)
		if isValid(input) {
			fmt.Println(input)
			break
		}
	}
}

func rule1(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		foundIncreasingStraight := true
		for j := i + 1; j <= i+2; j++ {
			if s[j]-s[j-1] != 1 {
				foundIncreasingStraight = false
			}
		}
		if foundIncreasingStraight {
			return true
		}
	}
	return false
}

func rule2(s string) bool {
	return !strings.ContainsAny(s, "iol")
}

func rule3(s string) bool {
	doubles := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			doubles++
			i += 1
			if doubles == 2 {
				return true
			}
		}
	}
	return false
}

func isValid(s string) bool {
	return rule1(s) && rule2(s) && rule3(s)
}

func next(s string) string {
	l := len(s)
	if s[l-1] != 'z' {
		if strings.ContainsAny("hkn", string(s[l-1])) {
			return string(append([]byte(s[:l-1]), s[l-1]+2))
		}
		return string(append([]byte(s[:l-1]), s[l-1]+1))
	}
	return string(append([]byte(next(s[:l-1])), 'a'))
}

func preprocess(s string) string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if strings.ContainsAny("iol", string(s[i])) {
			sb.WriteByte(s[i] + 1)
			for i++; i < len(s); i++ {
				sb.WriteByte('a')
			}
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
