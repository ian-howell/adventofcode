package main

func expandGrid(grid [][]byte) [][]byte {
	grid = expandRows(grid)
	grid = expandCols(grid)
	return grid
}

func expandRows(grid [][]byte) [][]byte {
	newGrid := [][]byte{}
	emptyRows := getEmptyRows(grid)
	for r := 0; r < len(grid); r++ {
		newGrid = append(newGrid, grid[r])
		if _, ok := emptyRows[r]; ok {
			newGrid = append(newGrid, grid[r])
		}
	}
	return newGrid
}

func expandCols(grid [][]byte) [][]byte {
	newGrid := [][]byte{}
	emptyCols := getEmptyCols(grid)
	for r := 0; r < len(grid); r++ {
		row := []byte{}
		for c := 0; c < len(grid[0]); c++ {
			row = append(row, grid[r][c])
			if _, ok := emptyCols[c]; ok {
				row = append(row, '.')
			}
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
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
