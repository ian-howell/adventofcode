package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	initialProgram := getProgram()
	result := -1
	program := make([]int, len(initialProgram))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(program, initialProgram)
			result, _ = doProgram(program, noun, verb)
			if result == 19690720 {
				fmt.Println(noun*100 + verb)
				break
			}
		}
	}
}

func getProgram() []int {
	var s string
	fmt.Scanf("%s", &s)
	strProgram := strings.Split(s, ",")
	intProgram := make([]int, len(strProgram))
	for i, item := range strProgram {
		var err error
		intProgram[i], err = strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
	}
	return intProgram
}

func doProgram(a []int, noun, verb int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			result, err = -1, fmt.Errorf("")
		}
	}()

	a[1], a[2] = noun, verb
	ptr := 0
	for {
		switch a[ptr] {
		case 1:
			a[a[ptr+3]] = a[a[ptr+1]] + a[a[ptr+2]]
		case 2:
			a[a[ptr+3]] = a[a[ptr+1]] * a[a[ptr+2]]
		case 99:
			return a[0], nil
		}
		ptr += 4
	}
}

func printProgram(program []int) {
	fmt.Print(program[0])
	for _, item := range program[1:] {
		fmt.Printf(",%d", item)
	}
}
