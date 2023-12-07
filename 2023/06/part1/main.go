package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	times, distances := readInput()

	product := 1
	for i := range times {
		product *= getWins(times[i], distances[i])
	}
	fmt.Println(product)
}

func getDistanceTraveled(timeHeld, timeTotal int) int {
	return (timeTotal - timeHeld) * timeHeld
}

func readInput() ([]int, []int) {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	times := strings.Fields(s.Text())[1:]

	s.Scan()
	distances := strings.Fields(s.Text())[1:]

	return toIntArray(times), toIntArray(distances)
}

func toIntArray(s []string) []int {
	nums := make([]int, 0, len(s))
	for _, a := range s {
		i, _ := strconv.Atoi(a)
		nums = append(nums, i)
	}
	return nums
}

func getWins(time, distance int) int {
	wins := 0
	for i := 1; i < time; i++ {
		distanceTraveled := getDistanceTraveled(i, time)
		if distanceTraveled > distance {
			wins++
		}
		if distanceTraveled <= distance && wins > 1 {
			break
		}
	}
	return wins
}
