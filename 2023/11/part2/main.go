package main

import "fmt"

type Point struct {
	Row int
	Col int
}

type Edge struct {
	A Point
	B Point
	W int
}

func (p Point) String() string {
	return fmt.Sprintf("<%v %v>", p.Row, p.Col)
}

const expansion = 1_000_000

func main() {
	grid := getInput()
	shortestPaths := getShortestPaths(grid)

	emptyRows := getEmptyRows(grid)
	emptyCols := getEmptyCols(grid)

	total := 0
	for _, path := range shortestPaths {
		total += path.W
		lower, higher := order(path.A.Row, path.B.Row)
		for r := lower + 1; r < higher; r++ {
			if _, ok := emptyRows[r]; ok {
				total += expansion - 1
			}
		}
		lower, higher = order(path.A.Col, path.B.Col)
		for c := lower + 1; c < higher; c++ {
			if _, ok := emptyCols[c]; ok {
				total += expansion - 1
			}
		}
	}
	fmt.Println(total)
}

func getShortestPaths(grid [][]byte) []Edge {
	edges := []Edge{}
	galaxies := getGalaxies(grid)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := getDistance(galaxies[i], galaxies[j])
			edges = append(edges, Edge{
				A: galaxies[i],
				B: galaxies[j],
				W: distance,
			})
		}
	}
	return edges
}

func getGalaxies(grid [][]byte) []Point {
	galaxies := []Point{}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '#' {
				galaxies = append(galaxies, Point{r, c})
			}
		}
	}
	return galaxies
}

func getDistance(a, b Point) int {
	return abs(a.Row-b.Row) + abs(a.Col-b.Col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getEmptyRows(grid [][]byte) map[int]struct{} {
	emptyRows := map[int]struct{}{}

	isRowEmpty := func(r int) bool {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '#' {
				return false
			}
		}
		return true
	}

	for r := 0; r < len(grid); r++ {
		if isRowEmpty(r) {
			emptyRows[r] = struct{}{}
		}
	}

	return emptyRows
}

func getEmptyCols(grid [][]byte) map[int]struct{} {
	emptyCols := map[int]struct{}{}

	isColEmpty := func(c int) bool {
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == '#' {
				return false
			}
		}
		return true
	}

	for c := 0; c < len(grid[0]); c++ {
		if isColEmpty(c) {
			emptyCols[c] = struct{}{}
		}
	}
	return emptyCols
}

func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
