package main

import "fmt"

func main() {
	numValid := 0
	done := false
	for !done {
		var isValid bool
		isValid, done = checkNext()
		if !done && isValid {
			numValid++
		}
	}
	fmt.Println(numValid)
}

func checkNext() (isValid bool, done bool) {
	var (
		pos1, pos2     int
		requiredLetter byte
		password       string
	)

	valsRead, _ := fmt.Scanf("%d-%d %c: %s", &pos1, &pos2, &requiredLetter, &password)
	if valsRead != 4 {
		return false, true
	}

	return xor(matches(password[pos1-1], requiredLetter), matches(password[pos2-1], requiredLetter)), false
}

func xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

func matches(a, b byte) bool {
	return a == b
}
