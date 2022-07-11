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
		w, err := strconv.Atoi(parts[3])
		if err != nil {
			panic(err)
		}
		if parts[2] == "lose" {
			w = -w
		}
		graph.addEdge(parts[0], parts[10][:len(parts[10])-1], w)
	}
	return graph
}

func (g Graph) addEdge(a, b string, weight int) {
	if _, found := g[a]; !found {
		g[a] = map[string]int{}
	}
	g[a][b] = weight
}

func (g Graph) Print() {
	for v, adjList := range g {
		fmt.Printf("%v:\n", v)
		for u, w := range adjList {
			fmt.Printf("  %v: %v\n", u, w)
		}
	}
}
