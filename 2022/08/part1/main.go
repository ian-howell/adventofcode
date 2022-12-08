package main

import "fmt"

func main() {
	grid := GetGrid()

	west := FromWest(grid)
	east := FromEast(grid)
	north := FromNorth(grid)
	south := FromSouth(grid)

	all := west.Or(east).Or(north).Or(south)

	fmt.Println(all.Count())
}
