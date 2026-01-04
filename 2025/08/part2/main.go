package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// kruskal's algorithm

func main() {
	points := getPoints()
	edges := findAllEdges(points)
	slices.SortFunc(edges, func(u, v Edge) int {
		if u.weight < v.weight {
			return -1
		}
		return 1
	})

	totalUnions := 0
	uf := NewUnionFind(len(points))
	for _, edge := range edges {
		if !uf.Union(edge.u, edge.v) {
			// already connected...
			continue
		}

		totalUnions++
		// There's 1000 nodes in the input, so to fully connect it,
		// we need at least 999 edges.
		if totalUnions == 999 {
			fmt.Println(points[edge.u].x * points[edge.v].x)
			return
		}
	}

}

type Edge struct {
	u, v   int
	weight float64
}

func findAllEdges(points []Point) []Edge {
	n := len(points)
	edges := make([]Edge, 0, n*n/2)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				u: i, v: j,
				weight: calcDistance(points[i], points[j]),
			})
		}
	}
	return edges
}

func calcDistance(a, b Point) float64 {
	deltaX := float64(a.x - b.x)
	deltaY := float64(a.y - b.y)
	deltaZ := float64(a.z - b.z)
	return math.Pow(deltaX, 2) + math.Pow(deltaY, 2) + math.Pow(deltaZ, 2)
}

type Point struct {
	x, y, z int
}

func getPoints() []Point {
	points := make([]Point, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		strVals := strings.Split(s.Text(), ",")
		points = append(points, Point{
			x: atoi(strVals[0]),
			y: atoi(strVals[1]),
			z: atoi(strVals[2]),
		})
	}
	return points
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type UnionFind struct {
	parent []int
	rank   []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Root(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Root(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) (joined bool) {
	rootOfX := uf.Root(x)
	rootOfY := uf.Root(y)

	if rootOfX == rootOfY {
		return false
	}

	switch {
	case uf.rank[rootOfX] > uf.rank[rootOfY]:
		uf.parent[rootOfY] = rootOfX
		uf.size[rootOfX] += uf.size[rootOfY]
	case uf.rank[rootOfX] < uf.rank[rootOfY]:
		uf.parent[rootOfX] = rootOfY
		uf.size[rootOfY] += uf.size[rootOfX]
	default:
		uf.parent[rootOfY] = rootOfX
		uf.rank[rootOfX]++
		uf.size[rootOfX] += uf.size[rootOfY]
	}
	return true
}

func (uf *UnionFind) Groups() map[int][]int {
	groups := make(map[int][]int)
	for i := range uf.parent {
		root := uf.Root(i)
		groups[root] = append(groups[root], i)
	}
	return groups
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Root(x)]
}
