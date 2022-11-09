package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input []byte
	input = []byte("uqwqemis")
	// input = []byte("abc")

	password := make([]byte, 8)
	found := 0
	for i := 0; found < 8; i++ {
		a := strconv.Itoa(i)
		testString := append(input, a...)
		hash := md5.Sum(testString)
		if has5leading0s(hash) {
			s := toString(hash)
			index, err := strconv.Atoi(s[5:6])
			if err == nil && index < 8 && password[index] == 0 {
				password[index] = s[6]
				found++
				printPassword(password)
			}
		}
	}
	fmt.Println(string(password))
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

func printPassword(s []byte) {
	for _, v := range s {
		if v == 0 {
			fmt.Print("_ ")
		} else {
			fmt.Printf("%v ", string(v))
		}
	}
	fmt.Println()
}
