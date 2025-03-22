package main

import (
	"fmt"
)

/*
// Get the "stationary window"
1. B = A & 0x3      Take the last 3 bits of A

// Calculate where the "volatile window" is
2. B = B xor 5      XOR them with 5 (call the new number x)

// Calculate the "stationary window" xor the "volatile window" xor 3
3. C = A / 2^B      Take the last x bits of A. Notice that we're going to do XORs and AND 3,
		    so we're really saying "Take the 3 bits of A starting at index x"
4. B = B xor 6      x = x xor 6
5. B = B xor C      x = x xor the 3 digits we took earlier (squish this into the above step?)

// Output that value
6. Output B & 0x3

// Sliiiiide to right
7. A = A / 2^3            (A = A / 8)     (A = A >> 3)
8. If A == 0: stop
9. Go to 1


Q: What does all of that mean? Why is it important?
A: Basically, we can now manipulate the outputs.

*/

const debug = true

func main() {
	vm := getInput()
	vm.a = 4095
	vm.Run()
	// for i, done := 0, false; !done; i++ {
	// 	if i%10_000 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	copied := vm
	// 	copied.results = nil
	// 	copied.a = i
	// 	copied.Run()
	// 	if slices.Equal(copied.results, copied.steps) {
	// 		fmt.Println(i)
	// 		done = true
	// 	}
	// }
}

func red(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func yellow(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", s)
}

func underline(s string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", s)
}

func debugln(s string) {
	if debug {
		fmt.Println(s)
	}
}
