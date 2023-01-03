package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func GetInput() map[Point]struct{} {
	points := map[Point]struct{}{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for point := range parseLine(s.Text()) {
			points[point] = struct{}{}
		}
	}
	return points
}

func parseLine(line string) map[Point]struct{} {
	parts := strings.Split(line, " -> ")
	currPoint := str2point(parts[0])

	points := map[Point]struct{}{currPoint: {}}

	for i := 1; i < len(parts); i++ {
		endPoint := str2point(parts[i])
		dr := getDiff(currPoint.Row, endPoint.Row)
		dc := getDiff(currPoint.Col, endPoint.Col)
		for currPoint != endPoint {
			currPoint = Point{
				Row: currPoint.Row + dr,
				Col: currPoint.Col + dc,
			}
			points[currPoint] = struct{}{}
		}
	}

	return points
}
func points2grid(points map[Point]struct{}) ([][]byte, int) {
	var (
		lr int = 0
		hr int = -1
		lc int = -1
		hc int = -1
	)

	for point := range points {
		if point.Row < lr {
			lr = point.Row
		}
		if hr == -1 || point.Row > hr {
			hr = point.Row
		}
		if lc == -1 || point.Col < lc {
			lc = point.Col
		}
		if hc == -1 || point.Col > hc {
			hc = point.Col
		}
	}

	var grid [][]byte
	for r := lr; r <= hr; r++ {
		var row []byte
		for c := lc; c <= hc; c++ {
			p := Point{Row: r, Col: c}
			if _, found := points[p]; found {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
		}
		grid = append(grid, row)
	}

	return grid, 500 - lc
}

func printGrid(grid [][]byte, startCol int) {
	fmt.Println("    0123456789")
	for r, row := range grid {
		fmt.Printf("% 3d ", r)
		for c, cell := range row {
			if r == 0 && c == startCol {
				fmt.Print("+")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}
	fmt.Println("    0123456789")
}
