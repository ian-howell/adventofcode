package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {
	results := make(chan int)
	wg := sync.WaitGroup{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "#") {
			continue
		}
		wg.Add(1)
		go func(str string) {
			process(str, results)
			wg.Done()
		}(s.Text())
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for result := range results {
		total += result
	}
	fmt.Println(total)
}

func process(s string, results chan<- int) {
	target, inputs := parse(s)
	if valid(target, inputs) {
		results <- target
	}
}

func valid(target int, inputs []int) bool {
	solutions := generateSolutions(inputs)
	_, ok := solutions[target]
	return ok
}

func generateSolutions(inputs []int) map[int]struct{} {
	if len(inputs) == 0 {
		return nil
	}

	if len(inputs) == 1 {
		return map[int]struct{}{inputs[0]: {}}
	}

	solutions := map[int]struct{}{}
	inputsCopy := slices.Clone(inputs[1:])

	inputsCopy[0] = inputs[0] * inputs[1]
	solutions = Union(solutions, generateSolutions(inputsCopy))

	inputsCopy[0] = inputs[0] + inputs[1]
	solutions = Union(solutions, generateSolutions(inputsCopy))

	inputsCopy[0] = Concatenate(inputs[0], inputs[1])
	solutions = Union(solutions, generateSolutions(inputsCopy))

	return solutions
}

func Union(a, b map[int]struct{}) map[int]struct{} {
	result := map[int]struct{}{}
	for val := range a {
		result[val] = struct{}{}
	}
	for val := range b {
		result[val] = struct{}{}
	}
	return result
}

func Concatenate(a, b int) int {
	numDigitsInB := log10(b)
	return a*pow10(numDigitsInB+1) + b
}

func parse(s string) (int, []int) {
	targetStr, after, _ := strings.Cut(s, ": ")
	inputsStr := strings.Fields(after)
	return atoi(targetStr), toInts(inputsStr)
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("failed to convert '%s' to int: %v", a, err)
	}
	return i
}

func toInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		ints = append(ints, atoi(s))
	}
	return ints
}

func log10(x int) int {
	return int(math.Log10(float64(x)))
}

func pow10(x int) int {
	return int(math.Pow10(x))
}
