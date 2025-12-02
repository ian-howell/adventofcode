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
	fmt.Printf("%10s%10s%10s\n", "Rotation", "Lock", "NumZeros")
	fmt.Printf("          %10d%10d\n", lock, numZeros)
	for _, rotation := range rotations {
		numZeros += lock.Rotate(rotation)
		fmt.Printf("%10d%10d%10d\n", rotation, lock, numZeros)
	}
}

type Lock int

type Rotation int

func (l *Lock) Rotate(r Rotation) int {
	s := int(*l) + int(r)

	if s >= 100 {
		// Crossed or landed on 0 going right
		*l = Lock(s % 100)
		return s / 100
	}

	if s > 0 {
		// Didn't cross 0
		*l = Lock(s)
		return 0
	}

	// Otherwise, we crossed or landed on 0 going left
	// If we started at 0, then we would've crossed 0 one fewer times
	if *l == 0 {
		*l = Lock(100 + (s % 100))
		return 1 - (s / 100) - 1
	}

	*l = Lock((100 + (s % 100)) % 100)
	return 1 - (s / 100)
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
