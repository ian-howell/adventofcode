package main

import (
	"fmt"
	"os"
)

func main() {
	scroller := NewScroller(os.Stdin)

	fmt.Println("GRID")
	for i, row := range scroller.grid {
		fmt.Printf("%v: %s\n", i, row)
	}
	fmt.Printf("operations: %q\n", scroller.operations)

	total := 0
	for scroller.Scroll() {
		total += scroller.Value()
	}
	fmt.Println(total)
}
