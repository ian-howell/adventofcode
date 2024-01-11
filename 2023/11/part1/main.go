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

func main() {
	grid := getInput()
	expandedGrid := expandGrid(grid)
	shortestPaths := getShortestPaths(expandedGrid)

	total := 0
	for _, path := range shortestPaths {
		total += path.W
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
