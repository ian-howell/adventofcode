package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tiles := getInput()

	largestArea := calcLargestRectangle(tiles)
	fmt.Println(largestArea)
}

type Tile struct {
	row, col int
}

func calcLargestRectangle(tiles []Tile) int {
	largest := 0
	for i := range tiles[:len(tiles)-1] {
		for j := i + 1; j < len(tiles); j++ {
			height := abs(tiles[i].row-tiles[j].row) + 1
			length := abs(tiles[i].col-tiles[j].col) + 1
			area := height * length
			// fmt.Printf("%v and %v = %v\n", tiles[i], tiles[j], area)
			largest = max(largest, area)
		}
	}
	return largest
}

func getInput() []Tile {
	tiles := []Tile{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		r, c, _ := strings.Cut(s.Text(), ",")
		tiles = append(tiles, Tile{atoi(r), atoi(c)})
	}
	return tiles
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
