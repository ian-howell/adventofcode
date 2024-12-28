package main

import "fmt"

const debug = false

func main() {
	vm := getInput()
	vm.Run()
	fmt.Println(vm.Results())
}

func red(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func debugln(s string) {
	if debug {
		fmt.Println(s)
	}
}
