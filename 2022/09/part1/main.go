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

type Rope struct {
	Head Point
	Tail Point
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
	r.Head = Point{
		Row: r.Head.Row + m[dir][0],
		Col: r.Head.Col + m[dir][1],
	}

	if r.Head.RowDiff(r.Tail) >= 2 {
		if r.Head.Col != r.Tail.Col {
			r.Tail.Col = r.Head.Col
		}
		r.Tail = Point{
			Row: r.Tail.Row + m[dir][0],
			Col: r.Tail.Col + m[dir][1],
		}
	} else if r.Head.ColDiff(r.Tail) >= 2 {
		if r.Head.Row != r.Tail.Row {
			r.Tail.Row = r.Head.Row
		}
		r.Tail = Point{
			Row: r.Tail.Row + m[dir][0],
			Col: r.Tail.Col + m[dir][1],
		}
	}

	// fmt.Printf("Moved %q: %+v\n", dir, r)
	// printGrid(r.Head, r.Tail)

	return r.Tail
}

func main() {
	rope := Rope{
		Head: Point{Row: 4},
		Tail: Point{Row: 4},
	}
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

func printGrid(head, tail Point) {
	for r := 0; r <= 4; r++ {
		for c := 0; c <= 5; c++ {
			p := Point{r, c}
			if p == head {
				fmt.Print("H ")
			} else if p == tail {
				fmt.Print("T ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
