package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	c := Computer{}
	c.A = 1
	var lines []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	c.LoadProgram(lines)
	// c.PrintInstructions()

	c.Run(false)

	fmt.Println(c.B)
}