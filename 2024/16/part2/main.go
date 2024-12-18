package main

import (
	"fmt"
)

func main() {
	grid := getInput()
	fmt.Println(dijkstras(grid))
}

type Cell struct {
	Row, Col int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return map[Direction]string{
		North: "North",
		East:  "East",
		South: "South",
		West:  "West",
	}[d]
}

type Vector struct{ X, Y int }

var Directions = map[Direction]Vector{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},
}

func (c Cell) Neighbor(dir Direction) Cell {
	return Cell{
		Row: c.Row + Directions[dir].Y,
		Col: c.Col + Directions[dir].X,
	}
}
