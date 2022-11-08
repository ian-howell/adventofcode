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

	visited := map[[2]int]struct{}{{0, 0}: {}}

	r, c := 0, 0
	facing := North
outerLoop:
	for _, step := range steps {
		direction := step[:1]
		facing = turn(direction, facing)
		for howMany, _ := strconv.Atoi(step[1:]); howMany > 0; howMany-- {
			r, c = takeSteps(r, c, facing)
			fmt.Println(r, c)
			if _, ok := visited[[2]int{r, c}]; ok {
				break outerLoop
			}
			visited[[2]int{r, c}] = struct{}{}
		}
	}
	fmt.Println(abs(r) + abs(c))
}

func takeSteps(r, c int, d DirectionType) (int, int) {
	delta := map[DirectionType][2]int{
		North: {0, 1},
		West:  {1, 0},
		South: {0, -1},
		East:  {-1, 0},
	}[d]
	return r + delta[0], c + delta[1]
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
