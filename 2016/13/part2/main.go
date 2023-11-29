package main

import "fmt"

const distance = 50

type Pair struct {
	x int
	y int
}

func NewPair(x, y int) Pair {
	return Pair{x: x, y: y}
}

var NullPair = NewPair(-1, -1)

func main() {
	var favNumber int
	fmt.Scan(&favNumber)

	d := bfs(NewPair(1, 1), favNumber)

	// for dv := range d {
	// 	fmt.Println(dv)
	// }

	fmt.Println(len(d))
}

func bfs(initialPosition Pair, favNumber int) map[Pair]int {
	q := []Pair{initialPosition}
	floorPlan := map[Pair]byte{initialPosition: 'O'}
	d := map[Pair]int{initialPosition: 0}

	for len(q) > 0 {
		var u Pair
		u, q = q[0], q[1:]

		for _, v := range getNeighbors(u, favNumber, floorPlan) {
			if _, ok := d[v]; !ok {
				d[v] = d[u] + 1
				if d[v] < distance {
					q = append(q, v)
				}
			}
		}
	}

	return d
}

func getNeighbors(u Pair, favNumber int, floorPlan map[Pair]byte) []Pair {
	neighbors := []Pair{}
	for _, m := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		candidate := NewPair(u.x+m[0], u.y+m[1])
		if candidate.x >= 0 && candidate.y >= 0 {
			_, ok := floorPlan[candidate]
			if !ok {
				floorPlan[candidate] = getSymbol(candidate, favNumber)
			}
			if floorPlan[candidate] == 'O' {
				neighbors = append(neighbors, candidate)
			}
		}
	}
	return neighbors
}

func f(p Pair) int {
	x, y := p.x, p.y
	return x*x + 3*x + 2*x*y + y + y*y
}

func countBits(x int) int {
	n := 0
	for x != 0 {
		x &= (x - 1)
		n++
	}
	return n
}

func getSymbol(p Pair, favNumber int) byte {
	if countBits(f(p)+favNumber)%2 == 0 {
		return 'O'
	}
	return '#'
}
