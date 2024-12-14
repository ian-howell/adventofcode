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
)

// 7131 is too low

func main() {
	// So I'm just gonna be honest. I watched these frame by frame until I noticed a pattern. At t=61,
	// most of the robots appeared to cluster in the center of the the screen. This pattern occurred at
	// every 101 seconds, so I adjusted my loop to jump by that interval. And then I just literally
	// watched it until the tree appeared. That just happened to occur at t=7132
	t := 7132
	robots := getRobots()
	newRobots := make([]Robot, 0, len(robots))
	for _, robot := range robots {
		newRobots = append(newRobots, move(robot, t))
	}
	fmt.Printf("T = %d\n", t)
	printGrid(newRobots)
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
