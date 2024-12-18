package main

import "fmt"

func dijkstras(grid Grid) int {
	start := State{
		Dir:  East,
		Cell: grid.Find('S'),
	}
	endPos := grid.Find('E')

	distanceTo := map[State]int{start: 0}
	queue := PriorityQueue{}
	queue.Push(Item{State: start, Priority: 0})

	parents := map[State]Set{start: {}}

	bestPathCells := map[Cell]struct{}{endPos: {}}

	bestDistance := -1

	for queue.Len() > 0 {
		u := queue.Pop()

		if (bestDistance == -1 || bestDistance == distanceTo[u.State]) && u.State.Cell == endPos {
			bestDistance = distanceTo[u.State]
			ancestors := getAncestors(parents, u.State)
			for _, ancestor := range ancestors {
				bestPathCells[ancestor.Cell] = struct{}{}
			}
		}

		for _, neighbor := range u.State.Neighbors(grid) {
			// Assume we're turning
			cost := 1000
			// If we didn't turn, it means we stepped forward
			if neighbor.Cell != u.State.Cell {
				cost = 1
			}
			distanceThroughU := distanceTo[u.State] + cost
			if _, ok := distanceTo[neighbor]; !ok {
				// We've never seen this state before
				parents[neighbor] = map[State]struct{}{u.State: {}}
				distanceTo[neighbor] = distanceThroughU
				queue.Push(Item{Priority: distanceThroughU, State: neighbor})
			} else if distanceThroughU < distanceTo[neighbor] {
				// We have seen this state before, and the new path is shorter than the path
				// we already knew about. We should update it and requeue it.
				parents[neighbor][u.State] = struct{}{}
				distanceTo[neighbor] = distanceThroughU
				queue.Push(Item{Priority: distanceThroughU, State: neighbor})
			} else if distanceThroughU == distanceTo[neighbor] {
				// This state has the same distance through u as it does through one of its
				// other parents.
				parents[neighbor][u.State] = struct{}{}
			}
		}

	}

	printBestPaths(grid, bestPathCells)

	return len(bestPathCells)
}

func getAncestors(tree map[State]Set, s State) []State {
	parents, ok := tree[s]
	if !ok {
		return nil
	}

	ancestors := []State{}
	for parent := range parents {
		ancestors = append(ancestors, parent)
		ancestors = append(ancestors, getAncestors(tree, parent)...)
	}

	return ancestors
}

func printBestPaths(grid [][]byte, cells map[Cell]struct{}) {
	for r, row := range grid {
		for c, val := range row {
			if _, ok := cells[Cell{r, c}]; ok {
				fmt.Print("0")
			} else {
				fmt.Printf("%v", string(val))
			}
		}
		fmt.Println()
	}
}

type State struct {
	Cell Cell
	Dir  Direction
}

func (s State) Neighbors(grid Grid) []State {
	neighbors := []State{
		s.TurnRight(),
		s.TurnRight().TurnRight().TurnRight(),
	}

	stepForward := State{
		Cell: s.Cell.Neighbor(s.Dir),
		Dir:  s.Dir,
	}
	if grid.At(stepForward.Cell) != '#' {
		neighbors = append(neighbors, stepForward)
	}

	return neighbors
}

func (s State) TurnRight() State {
	s.Dir = (s.Dir + 1) % Direction(len(Directions))
	return s
}

type Set map[State]struct{}
