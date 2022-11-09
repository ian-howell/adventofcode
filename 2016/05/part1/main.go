package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input []byte
	// fmt.Scan(&input)
	input = []byte("uqwqemis")

	sb := strings.Builder{}
	for i := 0; sb.Len() < 8; i++ {
		a := strconv.Itoa(i)
		testString := append(input, a...)
		hash := md5.Sum(testString)
		if has5leading0s(hash) {
			sb.WriteByte(toString(hash)[5])
		}
	}
	fmt.Println(sb.String())
}

func has5leading0s(b [16]byte) bool {
	// Stolen from 2015 day 4
	// Woah. Found a bug. Guess I got lucky last time
	// BROKEN: return b[0] == 0 && b[1] == 0 && b[2] < 10
	return b[0] == 0 && b[1] == 0 && b[2] < 0x10
}

func printHash(h [16]byte) {
	fmt.Println(toString(h))
}

func toString(h [16]byte) string {
	sb := strings.Builder{}
	for _, v := range h {
		sb.WriteString(fmt.Sprintf("%02x", v))
	}
	return sb.String()
}

func printBytes(h [16]byte) {
	for _, v := range h {
		fmt.Printf("%08b ", v)
	}
	fmt.Println()
}
