package main

import "fmt"

func main() {
	grid := GetGrid()
	fmt.Println(BestScenicScore(grid))
}

func BestScenicScore(grid Grid) int {
	best := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid)-1; c++ {
			score := ScenicScore(grid, r, c)
			if score > best {
				best = score
			}
		}
	}
	return best
}

func ScenicScore(grid Grid, row, col int) int {
	return lookWest(grid, row, col) *
		lookEast(grid, row, col) *
		lookNorth(grid, row, col) *
		lookSouth(grid, row, col)
}

func lookWest(grid Grid, row, col int) int {
	score := 0
	for c := col - 1; c >= 0; c-- {
		score++
		if grid[row][c] >= grid[row][col] {
			break
		}
	}
	return score
}

func lookEast(grid Grid, row, col int) int {
	score := 0
	for c := col + 1; c < len(grid); c++ {
		score++
		if grid[row][c] >= grid[row][col] {
			break
		}
	}
	return score
}

func lookSouth(grid Grid, row, col int) int {
	score := 0
	for r := row + 1; r < len(grid); r++ {
		score++
		if grid[r][col] >= grid[row][col] {
			break
		}
	}
	return score
}

func lookNorth(grid Grid, row, col int) int {
	score := 0
	for r := row - 1; r >= 0; r-- {
		score++
		if grid[r][col] >= grid[row][col] {
			break
		}
	}
	return score
}
