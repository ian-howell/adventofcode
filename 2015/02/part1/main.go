package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	total := 0
	for {
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		total += calculateRequiredPaper(input)
	}

	fmt.Println(total)
}

func calculateRequiredPaper(dimensions string) int {
	dim := strings.Split(dimensions, "x")
	h, _ := strconv.Atoi(dim[0])
	l, _ := strconv.Atoi(dim[1])
	w, _ := strconv.Atoi(dim[2])

	s1 := h * l
	s2 := h * w
	s3 := l * w

	smallest := s1
	if s2 < smallest {
		smallest = s2
	}
	if s3 < smallest {
		smallest = s3
	}
	return 2*(s1+s2+s3) + smallest
}
