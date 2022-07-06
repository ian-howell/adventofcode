package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	opens := strings.Count(input, "(")
	closes := strings.Count(input, ")")
	fmt.Println(opens-closes)
}
