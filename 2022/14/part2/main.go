package main

import (
	"fmt"
)

func main() {
	points, lowest := GetInput()

	count := 0
	startPoint := Point{Row: 0, Col: 500}
	for found := false; !found; _, found = points[startPoint] {
		dropSand(points, lowest)
		count++
	}

	fmt.Println(count)
}

func dropSand(points map[Point]struct{}, lowest int) {
	currPoint := Point{
		Row: 0,
		Col: 500,
	}

	falling := true
	for falling {
		currPoint.Row++
		// First check if this sand has fallen out of bounds
		if currPoint.Row >= lowest {
			currPoint.Row--
			break
		}

		if _, found := points[currPoint]; !found {
			continue
		}

		currPoint.Col--
		if _, found := points[currPoint]; !found {
			continue
		}

		currPoint.Col += 2 // first move back to the original column, then move 1 to the right
		if _, found := points[currPoint]; !found {
			continue
		}

		// Revert to the original point
		currPoint.Row--
		currPoint.Col--
		falling = false
	}

	points[currPoint] = struct{}{}
	return
}
