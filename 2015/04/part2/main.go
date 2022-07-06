package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	var input []byte
	fmt.Scan(&input)

	i := 1
	for {
		a := strconv.Itoa(i)
		testString := append(input, a...)
		hash := md5.Sum(testString)
		if has5leading0s(hash) {
			break
		}
		i++
	}

	fmt.Println(i)
}

func has5leading0s(b [16]byte) bool {
	return b[0] == 0 && b[1] == 0 && b[2] == 0
}

func printHash(h [16]byte) {
	for _, v := range h {
		fmt.Printf("%02x", v)
	}
	fmt.Println()
}
