package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	visited := map[int]map[int]struct{}{}

	r, c := 0, 0
	addToVisited(visited, r, c)

	for _, v := range input {
		switch v {
			case '^':
				r++
			case 'v':
				r--
			case '>':
				c++
			case '<':
				c--
		}
		addToVisited(visited, r, c)
	}

	fmt.Println(sizeof(visited))
}

func addToVisited(visited map[int]map[int]struct{}, r, c int) {
	if _, found := visited[r]; found {
		visited[r][c] = struct{}{}
	} else {
		visited[r] = map[int]struct{}{c: {}}
	}
}

func sizeof(visited map[int]map[int]struct{}) int {
	size := 0
	for _, v := range visited {
		size += len(v)
	}
	return size
}
