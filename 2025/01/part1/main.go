package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	numZeros := 0
	lock := Lock(50)
	rotations := readRotations(os.Stdin)
	for _, rotation := range rotations {
		lock.Rotate(rotation)
		if lock == 0 {
			numZeros++
		}
	}
	fmt.Println(numZeros)
}

type Lock int

type Rotation int

func (l *Lock) Rotate(r Rotation) {
	v := (int(*l) + int(r)) % 100
	if v < 0 {
		v += 100
	}
	*l = Lock(v)
}

func readRotations(r io.Reader) []Rotation {
	scanner := bufio.NewScanner(r)
	rotations := []Rotation{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			panic(fmt.Sprintf("line too short: '%s'", line))
		}
		rotation := stoi(line[1:])
		if line[0] == 'L' {
			rotation *= -1
		}
		rotations = append(rotations, Rotation(rotation))
	}
	if scanner.Err() != nil {
		panic(fmt.Sprintf("failed to scan: %s", scanner.Err()))
	}
	return rotations
}

func stoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to stoi: %s", err))
	}
	return i
}
