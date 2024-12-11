package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const numBlinks = 25

func main() {
	trees := getInput()
	for range numBlinks {
		for _, tree := range trees {
			blink(tree)
		}
	}

	totalLeaves := 0
	for _, tree := range trees {
		totalLeaves += tree.Leaves
	}
	fmt.Println(totalLeaves)
}

func blink(tree *Tree) {
	if tree.IsInternal() {
		blink(tree.Left)
		blink(tree.Right)
		tree.Leaves = tree.Left.Leaves + tree.Right.Leaves
		return
	}

	if tree.Value == 0 {
		tree.Value = 1
		return
	}

	if getNumDigits(tree.Value)%2 == 0 {
		l, r := split(tree.Value)
		tree.Left = NewTree(l)
		tree.Right = NewTree(r)
		tree.Leaves = 2
		return
	}

	tree.Value *= 2024
}

func split(x int) (l, r int) {
	str := strconv.Itoa(x)
	middle := len(str) / 2
	return atoi(str[:middle]), atoi(str[middle:])
}

func getNumDigits(x int) int {
	return int(math.Log10(float64(x))) + 1
}

func getInput() []*Tree {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	nodes := []*Tree{}
	for _, num := range toInts(s.Text()) {
		nodes = append(nodes, NewTree(num))
	}
	return nodes
}

func toInts(s string) []int {
	strs := strings.Fields(s)
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		ints = append(ints, atoi(str))
	}
	return ints
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func printTrees(trees []*Tree) {
	fmt.Print("[ ")
	for _, tree := range trees {
		fmt.Printf("%s ", tree)
	}
	fmt.Println("]")
}
