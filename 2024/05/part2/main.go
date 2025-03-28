package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Set map[int]struct{}
type Graph map[int]Set

func main() {
	rules, updates := getInput()

	total := 0
	for _, update := range updates {
		if !isValid(update, rules) {
			update = sort(rules, update)
			total += middleNum(update)
		}
	}
	fmt.Println(total)
}

func sort(rules Graph, update []int) []int {
	for i := 1; i < len(update); i++ {
		indexOfVal := i
		for j := i - 1; j >= 0; j-- {
			if _, found := rules[update[j]][update[indexOfVal]]; found {
				val := update[indexOfVal]
				update = slices.Delete(update, indexOfVal, indexOfVal+1)
				update = slices.Insert(update, j, val)
				indexOfVal = j
			}
		}
	}
	return update
}

func isValid(update []int, rules Graph) bool {
	for i, u := range update {
		for _, follower := range update[i+1:] {
			if _, found := rules[follower][u]; found {
				return false
			}
		}
	}
	return true
}

func middleNum(nums []int) int {
	return nums[len(nums)/2]
}

func getInput() (rules Graph, updates [][]int) {
	rules = Graph{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan(); s.Text() != ""; s.Scan() {
		from, to, _ := strings.Cut(s.Text(), "|")
		fromInt := atoi(from)
		toInt := atoi(to)
		if rules[fromInt] == nil {
			rules[fromInt] = Set{}
		}
		rules[fromInt][toInt] = struct{}{}
	}

	for s.Scan() {
		update := []int{}
		text := s.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		for text != "" {
			var num string
			num, text, _ = strings.Cut(text, ",")
			update = append(update, atoi(num))
		}
		updates = append(updates, update)
	}
	return rules, updates
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to parse '%s': %v", s, err)
	}
	return i
}
