package main

import (
	"bufio"
	"os"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func GetInput() (map[Point]struct{}, int) {
	lowest := 0
	points := map[Point]struct{}{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for point := range parseLine(s.Text()) {
			if point.Row > lowest {
				lowest = point.Row
			}
			points[point] = struct{}{}
		}
	}
	return points, lowest + 2
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
