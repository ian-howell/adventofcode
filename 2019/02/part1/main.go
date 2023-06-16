package main

import (
	"fmt"
	"os"
)

func main() {
	program := ReadProgram(os.Stdin)
	program[1] = 12
	program[2] = 2
	intCode := NewIntCode(
		// WithDebug(),
		WithProgram(program),
	)

	result, err := intCode.Run()
	if err != nil {
		fmt.Printf("Received the following error during execution: %v\n", err)
	}
	fmt.Println(result)
}
