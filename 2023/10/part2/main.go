package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func (p Point) Up() Point    { return Point{p.Row + Up.Row, p.Col + Up.Col} }
func (p Point) Down() Point  { return Point{p.Row + Down.Row, p.Col + Down.Col} }
func (p Point) Left() Point  { return Point{p.Row + Left.Row, p.Col + Left.Col} }
func (p Point) Right() Point { return Point{p.Row + Right.Row, p.Col + Right.Col} }

func (p Point) String() string {
	if p == Down {
		return "Down"
	}
	if p == Up {
		return "Up"
	}
	if p == Left {
		return "Left"
	}
	if p == Right {
		return "Right"
	}
	return fmt.Sprintf("<%v %v>", p.Row, p.Col)
}

var (
	NilPoint = Point{-1, -1}
	Down     = Point{1, 0}
	Up       = Point{-1, 0}
	Right    = Point{0, 1}
	Left     = Point{0, -1}

	AllDirections = []Point{Up, Down, Left, Right}
)

func main() {
	grid := getInput()

	path := dfs(grid)
	path = startAtTopLeft(path)
	startDirection := getDirection(path[0], path[1])
	if startDirection == Right || startDirection == Up {
		path = reverse(path)
		// Move the topleft position to the beginning of the path
		path = append([]Point{path[len(path)-1]}, path[:len(path)-1]...)
	}

	enclosedTiles := getEnclosedTiles(grid, path)

	// for fun
	printRedGreen(grid, toSet(path), enclosedTiles)

	fmt.Println(len(enclosedTiles))
}

func printHighlightedPath(grid []string, path []Point) {
	pathSet := toSet(path)
	for r := range grid {
		for c := range grid[0] {
			if _, ok := pathSet[Point{r, c}]; ok {
				fmt.Printf("\033[32m%v\033[0m", prettify(grid[r][c]))
			} else {
				fmt.Print(prettify(grid[r][c]))
			}
		}
		fmt.Println()
	}
}

func printRedGreen(grid []string, greenPoints map[Point]struct{}, redPoints map[Point]struct{}) {
	for r := range grid {
		for c := range grid[0] {
			if _, ok := greenPoints[Point{r, c}]; ok {
				fmt.Printf("\033[32m%v\033[0m", prettify(grid[r][c]))
			} else if _, ok := redPoints[Point{r, c}]; ok {
				fmt.Printf("\033[31m%v\033[0m", prettify(grid[r][c]))
			} else {
				fmt.Print(prettify(grid[r][c]))
			}
		}
		fmt.Println()
	}
}

func dfs(grid []string) []Point {
	startPoint := findAndFixStart(grid)
	visited := map[Point]struct{}{startPoint: {}}

	q := []Point{startPoint}

	path := []Point{}
	for len(q) > 0 {
		var u Point
		u, q = q[len(q)-1], q[:len(q)-1]
		path = append(path, u)

		for _, v := range getNeighbors(grid, u) {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				q = append(q, v)
			}
		}
	}

	return path
}

func bfs(grid []string, pathSet map[Point]struct{}, enclosedTiles map[Point]struct{}, startPoint Point) {

	getNonPathNeighbors := func(grid []string, point Point) []Point {
		r, c := point.Row, point.Col
		neighbors := []Point{}
		for _, d := range AllDirections {
			newPoint := Point{r + d.Row, c + d.Col}
			if _, ok := pathSet[newPoint]; !ok {
				neighbors = append(neighbors, newPoint)
			}
		}
		return neighbors
	}

	visited := map[Point]struct{}{startPoint: {}}

	q := []Point{startPoint}

	for len(q) > 0 {
		var u Point
		u, q = q[0], q[1:]

		enclosedTiles[u] = struct{}{}

		for _, v := range getNonPathNeighbors(grid, u) {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				q = append(q, v)
			}
		}
	}

}

func getEnclosedTiles(grid []string, path []Point) map[Point]struct{} {
	enclosedTiles := map[Point]struct{}{}
	pathSet := toSet(path)
	for i := 1; i < len(path); i++ {
		candidates := []Point{}
		direction := getDirection(path[i-1], path[i])
		switch direction {
		case Down: // look right
			switch grid[path[i].Row][path[i].Col] {
			case 'J':
				candidates = append(candidates, path[i].Down())
				fallthrough
			case '|':
				candidates = append(candidates, path[i].Right())
			}
		case Up: // look left
			switch grid[path[i].Row][path[i].Col] {
			case 'F':
				candidates = append(candidates, path[i].Up())
				fallthrough
			case '|':
				candidates = append(candidates, path[i].Left())
			}
		case Right: // look up
			switch grid[path[i].Row][path[i].Col] {
			case '7':
				candidates = append(candidates, path[i].Right())
				fallthrough
			case '-':
				candidates = append(candidates, path[i].Up())
			}
		case Left: // look down
			switch grid[path[i].Row][path[i].Col] {
			case 'L':
				candidates = append(candidates, path[i].Left())
				fallthrough
			case '-':
				candidates = append(candidates, path[i].Down())
			}
		}
		for _, candidate := range candidates {
			if _, ok := pathSet[candidate]; !ok {
				bfs(grid, pathSet, enclosedTiles, candidate)
			}
		}
	}
	return enclosedTiles
}

func getDirection(a, b Point) Point {
	return Point{b.Row - a.Row, b.Col - a.Col}
}

func toSet(list []Point) map[Point]struct{} {
	m := make(map[Point]struct{}, len(list))
	for _, item := range list {
		m[item] = struct{}{}
	}
	return m
}

func getNeighbors(grid []string, point Point) []Point {
	r, c := point.Row, point.Col
	directions := getDirectionsForSymbol(grid[r][c])
	neighbors := []Point{}
	for _, d := range directions {
		nr, nc := r+d.Row, c+d.Col
		if isValid(grid, nr, nc) {
			neighbors = append(neighbors, Point{nr, nc})
		}
	}
	return neighbors
}

func getDirectionsForSymbol(symbol byte) []Point {
	switch symbol {
	case '|':
		return []Point{Up, Down}
	case '-':
		return []Point{Right, Left}
	case 'J':
		return []Point{Left, Up}
	case 'F':
		return []Point{Down, Right}
	case '7':
		return []Point{Down, Left}
	case 'L':
		return []Point{Up, Right}
	}
	// Should be unreachable
	return nil
}

func findAndFixStart(grid []string) Point {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 'S' {
				startPoint := Point{r, c}
				startSymbol := getRealStartSymbol(grid, startPoint)
				grid[r] = strings.ReplaceAll(grid[r], "S", startSymbol)
				return startPoint
			}
		}
	}
	return NilPoint
}

func getRealStartSymbol(grid []string, startPoint Point) string {
	r, c := startPoint.Row, startPoint.Col

	connections := 0
	if isValid(grid, r-1, c) && bytes.ContainsAny([]byte{grid[r-1][c]}, "F7|") {
		connections |= 8
	}
	if isValid(grid, r+1, c) && bytes.ContainsAny([]byte{grid[r+1][c]}, "JL|") {
		connections |= 4
	}
	if isValid(grid, r, c-1) && bytes.ContainsAny([]byte{grid[r][c-1]}, "FL-") {
		connections |= 2
	}
	if isValid(grid, r, c+1) && bytes.ContainsAny([]byte{grid[r][c+1]}, "J7-") {
		connections |= 1
	}

	connectionMap := map[int]string{
		0b1100: "|",
		0b0011: "-",
		0b1010: "J",
		0b1001: "L",
		0b0110: "7",
		0b0101: "F",
	}

	return connectionMap[connections]
}

func isValid(grid []string, r int, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}

func prettify(x byte) string {
	switch x {
	case 'F':
		return string('┏')
	case '7':
		return string('┓')
	case 'L':
		return string('┗')
	case 'J':
		return string('┛')
	case '|':
		return string('┃')
	case '-':
		return string('━')
	}
	return "x"
}

func reverse(path []Point) []Point {
	revPath := make([]Point, 0, len(path))
	for i := len(path) - 1; i >= 0; i-- {
		revPath = append(revPath, path[i])
	}
	return revPath
}

func startAtTopLeft(path []Point) []Point {
	topLeft := NilPoint
	for _, step := range path {
		if topLeft == NilPoint || (step.Row <= topLeft.Row && step.Col <= topLeft.Col) {
			topLeft = step
		}
	}

	// Rotate the path until the first step is the top left
	// for path[0].Row != topLeft.Row && path[0].Col != topLeft.Col {
	for path[0] != topLeft {
		path = append(path, path[0])
		path = path[1:]
	}
	return path
}
