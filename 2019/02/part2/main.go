package main

import (
	"fmt"
	"os"
)

func main() {
	program := ReadProgram(os.Stdin)
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copyOfProgram := make([]int, len(program))
			copy(copyOfProgram, program)
			copyOfProgram[1] = noun
			copyOfProgram[2] = verb
			intCode := NewIntCode(
				// WithDebug(),
				WithProgram(copyOfProgram),
			)
			result, err := intCode.Run()
			if err != nil {
				fmt.Printf("Received the following error during execution: %v\n", err)
			}
			if result == 19690720 {
				fmt.Println(noun*100 + verb)
				return
			}
		}
	}

}
