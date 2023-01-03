package main

import (
	"fmt"
)

func main() {
	points := GetInput()

	grid, startCol := points2grid(points)
	// printGrid(grid, startCol)

	// fmt.Println("----------------------------------------")
	count := -1
	done := false
	for !done {
		done = dropSand(grid, startCol)
		count++
	}

	// printGrid(grid, startCol)
	fmt.Println(count)
}

func dropSand(grid [][]byte, startCol int) bool {
	r, c := 0, startCol

	falling := true
	for falling {
		r++
		// fmt.Printf("Checking (%v, %v)...\n", r, c)
		// First check if this sand has fallen out of bounds
		if r >= len(grid) {
			// fmt.Println("Fell off the bottom!")
			return true
		}

		if grid[r][c] == '.' {
			// fmt.Println("There's nothing immediately below this!")
			continue
		}

		c--
		// fmt.Printf("Checking (%v, %v)...\n", r, c)
		if c < 0 {
			// fmt.Println("Fell off the left!")
			return true
		}
		if grid[r][c] == '.' {
			// fmt.Println("There's something immediately below this, but there's nothing down left of it!")
			continue
		}

		c += 2 // first move back to the original column, then move 1 to the right
		// fmt.Printf("Checking (%v, %v)...\n", r, c)
		if c >= len(grid[0]) {
			// fmt.Println("Fell off the right!")
			return true
		}
		if grid[r][c] == '.' {
			// fmt.Println("There's something immediately below this and down left, but there's nothing down right of it!")
			continue
		}

		// fmt.Println("Holy bananas! There's something down right as well! This sand can't fall any further!")
		// Revert to the original point
		r--
		c--
		falling = false
	}

	grid[r][c] = 'o'
	return false
}
