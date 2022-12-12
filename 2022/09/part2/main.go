package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func (p Point) RowDiff(o Point) int {
	return abs(p.Row - o.Row)
}

func (p Point) ColDiff(o Point) int {
	return abs(p.Col - o.Col)
}

func (p Point) UnitVector(o Point) [2]int {
	rd := p.Row - o.Row
	if rd != 0 {
		rd /= abs(rd)
	}

	cd := p.Col - o.Col
	if cd != 0 {
		cd /= abs(cd)
	}

	return [2]int{rd, cd}
}

type Rope struct {
	Knots [10]Point
}

func (r Rope) Head() Point {
	return r.Knots[0]
}

func (r Rope) Tail() Point {
	return r.Knots[9]
}

func (r *Rope) Move(dir string, dist int) map[Point]struct{} {
	points := map[Point]struct{}{}
	for i := 0; i < dist; i++ {
		points[r.Step(dir)] = struct{}{}
	}
	return points
}

func (r *Rope) Step(dir string) Point {
	m := map[string][2]int{
		"R": {+0, +1},
		"L": {+0, -1},
		"D": {+1, +0},
		"U": {-1, +0},
	}
	// Update the "head"
	r.Knots[0] = Point{
		Row: r.Knots[0].Row + m[dir][0],
		Col: r.Knots[0].Col + m[dir][1],
	}

	for i := 1; i <= 9; i++ {
		if r.Knots[i-1].RowDiff(r.Knots[i]) >= 2 || r.Knots[i-1].ColDiff(r.Knots[i]) >= 2 {
			v := r.Knots[i-1].UnitVector(r.Knots[i])
			r.Knots[i] = Point{
				Row: r.Knots[i].Row + v[0],
				Col: r.Knots[i].Col + v[1],
			}
		}
	}

	// fmt.Printf("Moved %q: %+v\n", dir, r)
	// printGrid(r.Knots)

	return r.Knots[9]
}

func main() {
	var rope Rope
	for i := 0; i < 10; i++ {
		rope.Knots[i] = Point{Row: 15, Col: 11}
	}
	// printGrid(rope.Knots)
	trail := map[Point]struct{}{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		parts := strings.Split(s.Text(), " ")
		dir, dist := parts[0], atoi(parts[1])
		for point := range rope.Move(dir, dist) {
			trail[point] = struct{}{}
		}
	}

	// for point := range trail {
	// 	fmt.Println(point)
	// }

	fmt.Println(len(trail))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func printGrid(points [10]Point) {
	start := Point{Row: 15, Col: 11}
	for r := 0; r <= 22; r++ {
		for c := 0; c <= 27; c++ {
			p := Point{r, c}
			if p == points[0] {
				fmt.Print("H ")
			} else {
				isPoint := false
				for i := 1; i <= 9; i++ {
					if p == points[i] {
						fmt.Printf("%d ", i)
						isPoint = true
						break
					}
				}
				if !isPoint {
					if p == start {
						fmt.Print("s ")
					} else {
						fmt.Print(". ")
					}
				}
			}

		}
		fmt.Println()
	}
}
