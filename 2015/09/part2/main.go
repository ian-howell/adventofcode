package main

import "fmt"

func main() {
	graph := readGraph()
	fmt.Println(tsp(graph))
}

func tsp(g Graph) int {
	longest := -1
	for u := range g {
		distance := dfs(u, g, map[string]struct{}{u: {}})
		if longest < distance || longest == -1 {
			longest = distance
		}
	}
	return longest
}

func dfs(start string, g Graph, visited map[string]struct{}) int {
	if len(visited) == len(g) {
		return 0
	}

	longest := -1
	for v, w := range g[start] {
		if _, found := visited[v]; !found {
			copied := copySet(visited)
			copied[v] = struct{}{}
			distance := dfs(v, g, copied) + w
			if longest < distance || longest == -1 {
				longest = distance
			}
		}
	}
	return longest
}

func copySet(s map[string]struct{}) map[string]struct{} {
	c := map[string]struct{}{}
	for k := range s {
		c[k] = struct{}{}
	}
	return c
}
