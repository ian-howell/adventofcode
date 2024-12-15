package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"time"
)

const animate = false

type Direction struct {
	RowDelta, ColDelta int
}

func (d Direction) Scale(x int) Direction { return Direction{x * d.RowDelta, x * d.ColDelta} }

var DirectionMap = map[byte]Direction{'^': {-1, 0}, 'v': {+1, 0}, '>': {0, +1}, '<': {0, -1}}

type Cell struct {
	Row, Col int
}

func (c Cell) Get(dir Direction) Cell {
	return Cell{c.Row + dir.RowDelta, c.Col + dir.ColDelta}
}

type Grid [][]byte

func (g Grid) At(cell Cell) byte { return g[cell.Row][cell.Col] }
func (g *Grid) Swap(a, b Cell) {
	(*g)[a.Row][a.Col], (*g)[b.Row][b.Col] = (*g)[b.Row][b.Col], (*g)[a.Row][a.Col]
}
func (g Grid) Print() {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

func main() {
	grid, instructions := getInput()

	robot := findRobot(*grid)
	walk(robot, grid, instructions)

	total := 0
	for r, row := range *grid {
		for c, val := range row {
			if val == '[' {
				total += 100*r + c
			}
		}
	}
	fmt.Println(total)
}

func walk(robot *Cell, grid *Grid, instructions []byte) {
	for i, step := range instructions {
		if animate {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Printf("Moves: %d/%d\n", i, len(instructions))
			fmt.Println(string(step))
			grid.Print()
			time.Sleep(10 * time.Millisecond)
		}
		doStep(robot, grid, step)
	}
}

func doStep(robot *Cell, grid *Grid, step byte) {
	if step == '<' || step == '>' {
		doHorizontalStep(robot, grid, step)
		return
	}

	doVerticalStep(robot, grid, step)
}

func doHorizontalStep(robot *Cell, grid *Grid, step byte) {
	dir := DirectionMap[step]
	nextSpace := robot.Get(dir)

	for grid.At(nextSpace) == '[' || grid.At(nextSpace) == ']' {
		nextSpace = nextSpace.Get(dir.Scale(2))
	}

	// After all of those boxes, there was a wall. Nothing to do
	if grid.At(nextSpace) == '#' {
		return
	}

	// Turn around - we need to move the boxes starting from the back
	dir = dir.Scale(-1)

	// Push the boxes. If there were no boxes, this is a no-op
	for nextSpace.Get(dir) != *robot {
		grid.Swap(nextSpace.Get(dir), nextSpace)
		nextSpace = nextSpace.Get(dir)
	}

	// Move the robot
	grid.Swap(nextSpace, *robot)
	*robot = nextSpace
}

func doVerticalStep(robot *Cell, grid *Grid, step byte) {
	dir := DirectionMap[step]
	currentLayer := map[Cell]struct{}{*robot: {}}
	layers := []map[Cell]struct{}{currentLayer}
	currentLayer = getNextLayer(currentLayer, grid, dir)
	for containsBox(currentLayer, grid) {
		if containsWall(currentLayer, grid) {
			return
		}
		layers = append(layers, currentLayer)
		currentLayer = getNextLayer(currentLayer, grid, dir)
	}

	// This happens when there's no boxes in the final layer, but there *is* a wall
	if containsWall(currentLayer, grid) {
		return
	}

	// There is *not* a wall blocking the boxes, so we can push them
	slices.Reverse(layers)
	for _, layer := range layers {
		for cell := range layer {
			grid.Swap(cell, cell.Get(dir))
		}
	}

	*robot = robot.Get(dir)
}

func getNextLayer(layer map[Cell]struct{}, grid *Grid, dir Direction) map[Cell]struct{} {
	nextLayer := map[Cell]struct{}{}
	for cell := range layer {
		// if grid.At(cell) == '.' {
		// 	continue
		// }
		nextCell := cell.Get(dir)
		if grid.At(nextCell) == '.' {
			continue
		}
		nextLayer[nextCell] = struct{}{}
		if grid.At(nextCell) == '[' {
			nextLayer[nextCell.Get(DirectionMap['>'])] = struct{}{}
		} else if grid.At(nextCell) == ']' {
			nextLayer[nextCell.Get(DirectionMap['<'])] = struct{}{}
		}
	}
	return nextLayer
}

func containsBox(layer map[Cell]struct{}, grid *Grid) bool {
	for cell := range layer {
		if grid.At(cell) == '[' || grid.At(cell) == ']' {
			return true
		}
	}
	return false
}

func containsWall(layer map[Cell]struct{}, grid *Grid) bool {
	for cell := range layer {
		if grid.At(cell) == '#' {
			return true
		}
	}
	return false
}

func findRobot(grid Grid) *Cell {
	for r, row := range grid {
		for c, val := range row {
			if val == '@' {
				return &Cell{r, c}
			}
		}
	}
	return nil
}

func getInput() (*Grid, []byte) {
	grid := &Grid{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan(); s.Text() != ""; s.Scan() {
		row := make([]byte, 0, 2*len(s.Text()))
		for i := 0; i < len(s.Text()); i++ {
			switch s.Text()[i] {
			case '#':
				row = append(row, "##"...)
			case 'O':
				row = append(row, "[]"...)
			case '.':
				row = append(row, ".."...)
			case '@':
				row = append(row, "@."...)
			}
		}
		*grid = append(*grid, row)
	}

	instructions := []byte{}
	for s.Scan() {
		instructions = append(instructions, []byte(s.Text())...)
	}

	return grid, instructions
	// g := Grid([][]byte{
	// 	[]byte("####################"),
	// 	[]byte("##[]..[]......[][]##"),
	// 	[]byte("##[]...........[].##"),
	// 	[]byte("##...........@[][]##"),
	// 	[]byte("##..........[].[].##"),
	// 	[]byte("##..##[]..[].[]...##"),
	// 	[]byte("##...[]...[]..[]..##"),
	// 	[]byte("##.....[]..[].[][]##"),
	// 	[]byte("##........[]......##"),
	// 	[]byte("####################"),
	// })
	//
	// instructions = []byte{'v'}
	// return &g, instructions
}
