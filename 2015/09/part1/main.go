package main

import "fmt"

func main() {
	graph := readGraph()
	fmt.Println(tsp(graph))
}

func tsp(g Graph) int {
	smallest := -1
	for u := range g {
		distance := dfs(u, g, map[string]struct{}{u: {}})
		if smallest > distance || smallest == -1 {
			smallest = distance
		}
	}
	return smallest
}

func dfs(start string, g Graph, visited map[string]struct{}) int {
	if len(visited) == len(g) {
		return 0
	}

	smallest := -1
	for v, w := range g[start] {
		if _, found := visited[v]; !found {
			copied := copySet(visited)
			copied[v] = struct{}{}
			distance := dfs(v, g, copied) + w
			if smallest > distance || smallest == -1 {
				smallest = distance
			}
		}
	}
	return smallest
}

func copySet(s map[string]struct{}) map[string]struct{} {
	c := map[string]struct{}{}
	for k := range s {
		c[k] = struct{}{}
	}
	return c
}