package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph map[string]map[string]int

func readGraph() Graph {
	graph := Graph{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := s.Text()
		parts := strings.Split(input, " ")
		w, err := strconv.Atoi(parts[4])
		if err != nil {
			panic(err)
		}
		graph.addEdge(parts[0], parts[2], w)
	}
	return graph
}

func (g Graph) addEdge(a, b string, weight int) {
	if _, found := g[a]; !found {
		g[a] = map[string]int{}
	}
	if _, found := g[b]; !found {
		g[b] = map[string]int{}
	}
	g[a][b] = weight
	g[b][a] = weight
}

func (g Graph) Print() {
	for v, adjList := range g {
		fmt.Printf("%v:\n", v)
		for u, w := range adjList {
			fmt.Printf("  %v: %v\n", u, w)
		}
	}
}