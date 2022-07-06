package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Fields(input)
		tl := toPair(parts[len(parts)-3])
		br := toPair(parts[len(parts)-1])

		if parts[0] == "toggle" {
			toggle(grid, tl, br)
		} else if parts[1] == "on" {
			turnOn(grid, tl, br)
		} else if parts[1] == "off" {
			turnOff(grid, tl, br)
		}
	}

	total := 0
	for r := 0; r < 1000; r++ {
		for c := 0; c < 1000; c++ {
			total += grid[r][c]
		}
	}
	fmt.Println(total)
}

func toggle(grid [][]int, tl, br Pair) {
	for r := tl.Row(); r <= br.Row(); r++ {
		for c := tl.Col(); c <= br.Col(); c++ {
			grid[r][c] += 2
		}
	}
}

func set(grid [][]int, tl, br Pair, val int) {
	for r := tl.Row(); r <= br.Row(); r++ {
		for c := tl.Col(); c <= br.Col(); c++ {
			grid[r][c] += val
			if grid[r][c] < 0 {
				grid[r][c] = 0
			}
		}
	}
}

func turnOn(grid [][]int, tl, br Pair)  { set(grid, tl, br, 1) }
func turnOff(grid [][]int, tl, br Pair) { set(grid, tl, br, -1) }

type Pair [2]int

func (p *Pair) Row() int { return p[0] }
func (p *Pair) Col() int { return p[1] }
func toPair(s string) Pair {
	parts := strings.Split(s, ",")
	row, _ := strconv.Atoi(parts[0])
	col, _ := strconv.Atoi(parts[1])
	return Pair{row, col}
}
