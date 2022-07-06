package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	if len(input)%2 == 1 {
		input += " "
	}

	santaPath, roboPath := []byte{}, []byte{}
	for i := 0; i < len(input); i += 2 {
		santaPath = append(santaPath, input[i])
		roboPath = append(roboPath, input[i+1])
	}

	visited := map[int]map[int]struct{}{}
	deliver(visited, santaPath)
	deliver(visited, roboPath)

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

func deliver(visited map[int]map[int]struct{}, path []byte) {

	r, c := 0, 0
	addToVisited(visited, r, c)

	for _, v := range path {
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
}
