package main

func dijkstras(grid Grid) int {
	start := State{
		Dir:  East,
		Cell: grid.Find('S'),
	}
	endPos := grid.Find('E')

	distanceTo := map[State]int{start: 0}
	queue := PriorityQueue{}
	queue.Push(Item{State: start, Priority: 0})

	for queue.Len() > 0 {
		u := queue.Pop()

		if u.State.Cell == endPos {
			return distanceTo[u.State]
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
				distanceTo[neighbor] = distanceThroughU
				queue.Push(Item{Priority: distanceThroughU, State: neighbor})
			} else if distanceThroughU < distanceTo[neighbor] {
				// We've already seen it.
				// The new path is shorter than the path we already had, so we should update
				// it and requeue it... It kinda sucks that I'm cropping up a bunch of extra
				// items at the back of the queue. TODO Fix this if I need more speed.
				distanceTo[neighbor] = min(distanceTo[neighbor], distanceThroughU)
				queue.Push(Item{Priority: distanceThroughU, State: neighbor})
			}
		}

	}

	return 0
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
