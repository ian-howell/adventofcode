package main

import "fmt"

func main() {
	s1 := readState()

	path := BFS(s1)
	// for i := len(path) - 1; i >= 0; i-- {
	// 	step := path[i]
	// 	fmt.Println(NewStateFromHash(step))
	// 	fmt.Println("----------------------------------------")
	// }
	fmt.Println(len(path))
}
