package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector struct {
	RowDelta int
	ColDelta int
}

func (c Vector) Minus(other Vector) Vector {
	return Vector{
		RowDelta: c.RowDelta - other.RowDelta,
		ColDelta: c.ColDelta - other.ColDelta,
	}
}

func (c Vector) Plus(other Vector) Vector {
	return Vector{
		RowDelta: c.RowDelta + other.RowDelta,
		ColDelta: c.ColDelta + other.ColDelta,
	}
}

func (v Vector) Scale(x int) Vector {
	return Vector{
		RowDelta: x * v.RowDelta,
		ColDelta: x * v.ColDelta,
	}
}

func main() {
	grid := getGrid()
	antennas := getAntennas(grid)
	uniqueAntinodes := map[Vector]struct{}{}
	for _, cells := range antennas {
		for _, antinode := range getAntinodes(cells) {
			if inBounds(antinode, grid) {
				uniqueAntinodes[antinode] = struct{}{}
			}
		}
	}
	fmt.Println(len(uniqueAntinodes))
}

func inBounds(v Vector, grid [][]byte) bool {
	between := func(lb, x, ub int) bool { return lb <= x && x < ub }
	return between(0, v.RowDelta, len(grid)) && between(0, v.ColDelta, len(grid[0]))
}

func getAntinodes(cells []Vector) []Vector {
	antinodes := []Vector{}
	for i, a := range cells {
		for _, b := range cells[i+1:] {
			v := b.Minus(a)
			antinodes = append(
				antinodes,
				a.Minus(v),
				b.Plus(v),
			)
		}
	}
	return antinodes
}

func getAntennas(grid [][]byte) map[byte][]Vector {
	antennas := map[byte][]Vector{}
	for r, row := range grid {
		for c, val := range row {
			if val == '.' {
				continue
			}
			antennas[val] = append(
				antennas[val],
				Vector{
					RowDelta: r,
					ColDelta: c,
				},
			)
		}
	}
	return antennas
}

func getGrid() [][]byte {
	grid := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}
