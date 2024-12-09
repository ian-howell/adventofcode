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
	antennasByFrequency := getAntennasByFrequency(grid)
	uniqueAntinodes := map[Vector]struct{}{}
	for _, antennas := range antennasByFrequency {
		for _, antinode := range getAntinodes(antennas, grid) {
			uniqueAntinodes[antinode] = struct{}{}
		}
	}
	fmt.Println(len(uniqueAntinodes))
}

func inBounds(v Vector, grid [][]byte) bool {
	between := func(lb, x, ub int) bool { return lb <= x && x < ub }
	return between(0, v.RowDelta, len(grid)) && between(0, v.ColDelta, len(grid[0]))
}

func getAntinodes(antennas []Vector, grid [][]byte) []Vector {
	antinodes := []Vector{}
	for i, a := range antennas {
		for _, b := range antennas[i+1:] {
			v := b.Minus(a)
			// From a toward b
			for antinode := a.Plus(v); inBounds(antinode, grid); antinode = antinode.Plus(v) {
				antinodes = append(antinodes, antinode)
			}

			// From b toward a
			for antinode := b.Minus(v); inBounds(antinode, grid); antinode = antinode.Minus(v) {
				antinodes = append(antinodes, antinode)
			}
		}
	}
	return antinodes
}

func getAntennasByFrequency(grid [][]byte) map[byte][]Vector {
	antennasByFrequency := map[byte][]Vector{}
	for r, row := range grid {
		for c, val := range row {
			if val == '.' {
				continue
			}
			antennasByFrequency[val] = append(
				antennasByFrequency[val],
				Vector{
					RowDelta: r,
					ColDelta: c,
				},
			)
		}
	}
	return antennasByFrequency
}

func getGrid() [][]byte {
	grid := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}
