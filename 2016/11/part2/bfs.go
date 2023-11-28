package main

func BFS(s0 State) []string {
	path := []string{}

	q := []string{s0.Hash()}
	cameFrom := map[string]string{s0.Hash(): ""}

	for len(q) > 0 {
		var u string
		u, q = q[0], q[1:]

		uState := NewStateFromHash(u)

		if isTargetState(uState) {
			current := u
			for current != "" {
				path = append(path, current)
				current = cameFrom[current]
			}
			return path[:len(path)-1]
		}

		for _, v := range uState.Next() {
			vHash := v.Hash()
			if _, ok := cameFrom[vHash]; !ok {
				cameFrom[vHash] = u
				q = append(q, vHash)
			}
		}

	}

	return path
}

func isTargetState(s State) bool {
	if s.ElevatorFloor != 3 {
		return false
	}

	for i := 0; i < 3; i++ {
		if len(s.Generators[i]) > 0 {
			return false
		}
		if len(s.Microchips[i]) > 0 {
			return false
		}
	}
	return true
}
