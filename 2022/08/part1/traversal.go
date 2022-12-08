package main

type RelativeDirection int

const (
	DownRight RelativeDirection = 1
	UpLeft                      = -1
)

type CompassDirection int

const (
	West CompassDirection = iota + 1
	East
	North
	South
)

func FromWest(grid Grid) Grid {
	return FromCompassDirection(grid, West)
}

func FromEast(grid Grid) Grid {
	return FromCompassDirection(grid, East)
}

func FromNorth(grid Grid) Grid {
	return FromCompassDirection(grid, North)
}

func FromSouth(grid Grid) Grid {
	return FromCompassDirection(grid, South)
}

func FromCompassDirection(grid Grid, from CompassDirection) Grid {
	switch from {
	case West:
		return rowWise(grid, DownRight)
	case East:
		return rowWise(grid, UpLeft)
	case North:
		return colWise(grid, DownRight)
	case South:
		return colWise(grid, UpLeft)
	}
	return nil
}

func rowWise(grid Grid, dir RelativeDirection) Grid {
	start, end := getEndpoints(dir, len(grid)-1)
	result := InitGrid(grid)
	for r := 0; r < len(grid[0]); r++ {
		tallest := -1
		for c := start; c != end+int(dir); c += int(dir) {
			tallest, result[r][c] = getTallest(grid[r][c], tallest)
		}
	}
	return result
}

func colWise(grid Grid, dir RelativeDirection) Grid {
	start, end := getEndpoints(dir, len(grid)-1)
	result := InitGrid(grid)
	for c := 0; c < len(grid); c++ {
		tallest := -1
		for r := start; r != end+int(dir); r += int(dir) {
			tallest, result[r][c] = getTallest(grid[r][c], tallest)
		}
	}
	return result
}

func getEndpoints(dir RelativeDirection, max int) (int, int) {
	switch dir {
	case DownRight:
		return 0, max
	case UpLeft:
		return max, 0
	}
	return -1, -1
}

func getTallest(candidate, tallest int) (int, int) {
	if candidate > tallest {
		return candidate, 1
	}
	return tallest, 0
}
