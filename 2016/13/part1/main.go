package main

import "fmt"

type Pair struct {
	x int
	y int
}

func NewPair(x, y int) Pair {
	return Pair{x: x, y: y}
}

var NullPair = NewPair(-1, -1)

// var Target = NewPair(7, 4)

var Target = NewPair(31, 39)

func main() {
	var favNumber int
	fmt.Scan(&favNumber)

	path := bfs(NewPair(1, 1), favNumber)

	// for i := len(path) - 1; i >= 0; i-- {
	// 	fmt.Println(path[i])
	// }

	fmt.Println(len(path) - 1)
}

func bfs(initialPosition Pair, favNumber int) []Pair {
	q := []Pair{initialPosition}
	cameFrom := map[Pair]Pair{initialPosition: NullPair}
	floorPlan := map[Pair]byte{initialPosition: 'O'}

	for len(q) > 0 {
		var u Pair
		u, q = q[0], q[1:]

		if u == Target {
			path := []Pair{}
			for u != NullPair {
				path = append(path, u)
				u = cameFrom[u]
			}
			return path
		}

		for _, v := range getNeighbors(u, favNumber, floorPlan) {
			if _, ok := cameFrom[v]; !ok {
				q = append(q, v)
				cameFrom[v] = u
			}
		}
	}

	return nil
}

func getNeighbors(u Pair, favNumber int, floorPlan map[Pair]byte) []Pair {
	neighbors := []Pair{}
	for _, m := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		candidate := NewPair(u.x+m[0], u.y+m[1])
		_, ok := floorPlan[candidate]
		if !ok {
			floorPlan[candidate] = getSymbol(candidate, favNumber)
		}
		if floorPlan[candidate] == 'O' {
			neighbors = append(neighbors, candidate)
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
