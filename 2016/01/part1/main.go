package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	steps := strings.Split(s.Text(), ", ")

	r, c := 0, 0
	facing := North
	for _, step := range steps {
		direction := step[:1]
		howMany, _ := strconv.Atoi(step[1:])
		facing = turn(direction, facing)
		r, c = takeSteps(r, c, facing, howMany)
	}
	fmt.Println(abs(r) + abs(c))
}

func takeSteps(r, c int, d DirectionType, howMany int) (int, int) {
	delta := map[DirectionType][2]int{
		North: {0, 1},
		West:  {1, 0},
		South: {0, -1},
		East:  {-1, 0},
	}[d]
	return r + delta[0]*howMany, c + delta[1]*howMany
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
