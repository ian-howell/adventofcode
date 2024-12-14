package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

type Velocity struct {
	X int
	Y int
}

type Robot struct {
	Position Position
	Velocity Velocity
}

const (
	gridWidth  = 101
	gridHeight = 103

	// Use this for the small_input
	// gridWidth  = 11
	// gridHeight = 7
)

func main() {
	robots := getRobots()
	newRobots := make([]Robot, 0, len(robots))
	for _, robot := range robots {
		newRobots = append(
			newRobots,
			move(robot, 100),
		)
	}

	fmt.Println(safetyFactor(newRobots))
}

func move(robot Robot, distance int) Robot {
	x := (robot.Position.X + distance*robot.Velocity.X) % gridWidth
	if x < 0 {
		x += gridWidth
	}
	y := (robot.Position.Y + distance*robot.Velocity.Y) % gridHeight
	if y < 0 {
		y += gridHeight
	}
	return Robot{
		Position: Position{x, y},
		Velocity: robot.Velocity,
	}
}

func safetyFactor(robots []Robot) int {
	halfWidth := gridWidth / 2
	halfHeight := gridHeight / 2
	return robotsInArea(0, halfWidth, 0, halfHeight, robots) *
		robotsInArea(halfWidth+1, gridWidth, 0, halfHeight, robots) *
		robotsInArea(0, halfWidth, halfHeight+1, gridHeight, robots) *
		robotsInArea(halfWidth+1, gridWidth, halfHeight+1, gridHeight, robots)
}

func between(lb, x, ub int) bool {
	return lb <= x && x < ub
}

func robotsInArea(lx, ux, ly, uy int, robots []Robot) int {
	total := 0
	for _, robot := range robots {
		if between(lx, robot.Position.X, ux) && between(ly, robot.Position.Y, uy) {
			total++
		}
	}
	return total
}

func getRobots() []Robot {
	robots := []Robot{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		posStr, velStr, _ := strings.Cut(s.Text(), " ")
		_, posStr, _ = strings.Cut(posStr, "=")
		_, velStr, _ = strings.Cut(velStr, "=")

		posXStr, posYStr, _ := strings.Cut(posStr, ",")
		velXStr, velYStr, _ := strings.Cut(velStr, ",")

		robots = append(
			robots,
			Robot{
				Position: Position{X: atoi(posXStr), Y: atoi(posYStr)},
				Velocity: Velocity{X: atoi(velXStr), Y: atoi(velYStr)},
			},
		)
	}
	return robots
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func printGrid(robots []Robot) {
	for y := range gridHeight {
		for x := range gridWidth {
			n := 0
			for _, robot := range robots {
				if robot.Position.X == x && robot.Position.Y == y {
					n++
				}
			}
			if n == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", n)
			}
		}
		fmt.Println()
	}
}
