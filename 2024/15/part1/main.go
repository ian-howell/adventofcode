package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const animate = false

type Direction struct {
	RowDelta, ColDelta int
}

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
			if val == 'O' {
				total += 100*r + c
			}
		}
	}
	fmt.Println(total)
}

func walk(robot *Cell, grid *Grid, instructions []byte) {
	if animate {
		grid.Print()
	}
	for _, step := range instructions {
		doStep(robot, grid, step)
		if animate {
			time.Sleep(25 * time.Millisecond)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			grid.Print()
		}
	}
}

func doStep(robot *Cell, grid *Grid, step byte) {
	dir := DirectionMap[step]
	nextSpace := robot.Get(dir)

	for grid.At(nextSpace) == 'O' {
		nextSpace = nextSpace.Get(dir)
	}

	// After all of those boxes, there was a wall. Nothing to do
	if grid.At(nextSpace) == '#' {
		return
	}

	// Push the boxes. If there were no boxes, this is a no-op
	grid.Swap(robot.Get(dir), nextSpace)

	// Move the robot
	grid.Swap(robot.Get(dir), *robot)
	*robot = robot.Get(dir)
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
		*grid = append(*grid, []byte(s.Text()))
	}

	instructions := []byte{}
	for s.Scan() {
		instructions = append(instructions, []byte(s.Text())...)
	}

	return grid, instructions
}
