package main

type Guard struct {
	Pos Cell
	Dir Direction
}

type Cell struct {
	Row int
	Col int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Vector struct{ X, Y int }

var Directions = map[Direction]Vector{
	North: {0, -1},
	East:  {1, 0},
	South: {0, 1},
	West:  {-1, 0},
}

func (g *Guard) Step(grid [][]byte) {
	nextCell := getCellInDirection(g.Pos, g.Dir)
	if inBounds(nextCell, grid) && grid[nextCell.Row][nextCell.Col] == '#' {
		g.TurnRight()
		return
	}
	g.Pos = nextCell
}

func (g *Guard) TurnRight() {
	g.Dir = (g.Dir + 1) % Direction(len(Directions))
}

func getCellInDirection(pos Cell, dir Direction) Cell {
	return Cell{
		Row: pos.Row + Directions[dir].Y,
		Col: pos.Col + Directions[dir].X,
	}
}
