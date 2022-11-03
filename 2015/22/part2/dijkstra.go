package main

func Dijkstra(start State) []State {
	cameFrom := map[State]*State{start: nil}
	distanceTo := map[State]int{start: 0}
	queue := MinHeap[State]{}
	queue.Push(start, 0)

	for queue.Len() != 0 {

		u, du, _ := queue.Pop()

		if u.BossHP <= 0 {
			return createPath(cameFrom, u)
		}

		for _, v := range Neighbors(u) {
			if u.PlayerHP > 0 {
				if dv, ok := distanceTo[v]; !ok || du+Weight(v) < dv {
					distanceTo[v] = du + Weight(v)
					cameFrom[v] = &u
					queue.Push(v, distanceTo[v])
				}
			}
		}
	}

	return nil
}

func createPath[T comparable](cameFrom map[T]*T, end T) []T {
	path := []T{end}
	for next := cameFrom[end]; next != nil; next = cameFrom[end] {
		path = append([]T{*next}, path...)
		end = *next
	}
	return path
}
