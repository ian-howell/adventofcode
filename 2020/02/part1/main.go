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
		lb, ub         int
		requiredLetter byte
		password       string
	)

	valsRead, _ := fmt.Scanf("%d-%d %c: %s", &lb, &ub, &requiredLetter, &password)
	if valsRead != 4 {
		return false, true
	}

	count := 0
	for i := 0; i < len(password); i++ {
		if password[i] == requiredLetter {
			count++
			if count > ub {
				return false, false
			}
		}
	}

	return (lb <= count && count <= ub), false
}
