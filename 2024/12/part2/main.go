package main

import (
	"bufio"
	"fmt"
	"os"
)

// 899196 is the answer

type Cell struct{ Row, Col int }

func (c Cell) West() Cell  { return Cell{c.Row, c.Col - 1} }
func (c Cell) East() Cell  { return Cell{c.Row, c.Col + 1} }
func (c Cell) North() Cell { return Cell{c.Row - 1, c.Col} }
func (c Cell) South() Cell { return Cell{c.Row + 1, c.Col} }

type Region map[Cell]struct{}

func (r Region) Has(c Cell) bool { _, ok := r[c]; return ok }

func main() {
	grid := getGrid()
	total := 0
	for r, row := range grid {
		for c, val := range row {
			// Modify the grid in place by deleting every part of this particular region
			if val != '.' {
				total += getPrice(grid, Cell{r, c})
			}
		}
	}
	fmt.Println(total)
}

func getPrice(grid [][]byte, cell Cell) int {
	region := floodFill(grid, cell)
	numWalls := countWalls(region)
	return len(region) * numWalls
}

func floodFill(grid [][]byte, start Cell) Region {
	queue := []Cell{start}
	visited := map[Cell]struct{}{start: {}}
	symbol := grid[start.Row][start.Col]

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for _, v := range next(u, grid, symbol) {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				queue = append(queue, v)
			}
		}
		grid[u.Row][u.Col] = '.'
	}

	return visited
}

func next(cell Cell, grid [][]byte, symbol byte) []Cell {
	vs := []Cell{}
	dirMatrix := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range dirMatrix {
		v := Cell{
			Row: cell.Row + dir[0],
			Col: cell.Col + dir[1],
		}
		if inBounds(v, grid) && grid[v.Row][v.Col] == symbol {
			vs = append(vs, v)
		}
	}
	return vs
}

func inBounds(u Cell, grid [][]byte) bool {
	between := func(lb, x, ub int) bool { return lb <= x && x < ub }
	return between(0, u.Row, len(grid)) && between(0, u.Col, len(grid[0]))
}

func countWalls(region Region) int {
	topLeft, bottomRight := findBoundingCorners(region)
	return countNorthWalls(region, topLeft, bottomRight) +
		countSouthWalls(region, topLeft, bottomRight) +
		getWestWalls(region, topLeft, bottomRight) +
		getEastWalls(region, topLeft, bottomRight)
}

func countNorthWalls(region Region, topLeft, bottomRight Cell) int {
	return countHorizontalWalls(region, topLeft, bottomRight, Cell.North)
}

func countSouthWalls(region Region, topLeft, bottomRight Cell) int {
	return countHorizontalWalls(region, topLeft, bottomRight, Cell.South)
}

func getWestWalls(region Region, topLeft, bottomRight Cell) int {
	return countVerticalWalls(region, topLeft, bottomRight, Cell.West)
}

func getEastWalls(region Region, topLeft, bottomRight Cell) int {
	return countVerticalWalls(region, topLeft, bottomRight, Cell.East)
}

func countHorizontalWalls(region Region, topLeft, bottomRight Cell, direction func(Cell) Cell) int {
	numWalls := 0
	for r := topLeft.Row; r <= bottomRight.Row; r++ {
		prevHadWall := false
		for c := topLeft.Col; c <= bottomRight.Col+1; c++ {
			cell := Cell{r, c}
			currHasWall := region.Has(cell) && !region.Has(direction(cell))
			if !prevHadWall && currHasWall {
				numWalls++
			}
			prevHadWall = currHasWall
		}
	}
	return numWalls
}

func countVerticalWalls(region Region, topLeft, bottomRight Cell, direction func(Cell) Cell) int {
	numWalls := 0
	for c := topLeft.Col; c <= bottomRight.Col; c++ {
		prevHadWall := false
		for r := topLeft.Row; r <= bottomRight.Row+1; r++ {
			cell := Cell{r, c}
			currHasWall := region.Has(cell) && !region.Has(direction(cell))
			if !prevHadWall && currHasWall {
				numWalls++
			}
			prevHadWall = currHasWall
		}
	}
	return numWalls
}

func findBoundingCorners(region Region) (topLeft, bottomRight Cell) {
	for cell := range region {
		// Just grab one of the cells at random to use as a starting point
		topLeft = cell
		bottomRight = cell
		break
	}

	for cell := range region {
		topLeft.Row = min(topLeft.Row, cell.Row)
		topLeft.Col = min(topLeft.Col, cell.Col)
		bottomRight.Row = max(bottomRight.Row, cell.Row)
		bottomRight.Col = max(bottomRight.Col, cell.Col)

	}
	return topLeft, bottomRight
}

func getGrid() [][]byte {
	grid := [][]byte{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	return grid
}
