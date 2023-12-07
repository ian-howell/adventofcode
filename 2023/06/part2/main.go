package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	time, distance := readInput()
	fmt.Println(getWins(time, distance))
}

func getDistanceTraveled(timeHeld, timeTotal int) int {
	return (timeTotal - timeHeld) * timeHeld
}

func readInput() (int, int) {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	time, _ := strconv.Atoi(strings.Join(strings.Fields(s.Text())[1:], ""))

	s.Scan()
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(s.Text())[1:], ""))

	return time, distance
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
