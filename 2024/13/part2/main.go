package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ClawMachine struct {
	ax, ay, bx, by, prizeX, prizeY int
}

const unitConversionError = 10000000000000

func main() {
	total := 0
	clawMachines := getInput()
	for _, clawMachine := range clawMachines {
		aPresses, bPresses, exists := Solve(
			LinearEquation{
				a: float64(clawMachine.ax),
				b: float64(clawMachine.bx),
				s: float64(clawMachine.prizeX) + unitConversionError,
			},
			LinearEquation{

				a: float64(clawMachine.ay),
				b: float64(clawMachine.by),
				s: float64(clawMachine.prizeY) + unitConversionError,
			},
		)

		if exists {
			total += int(aPresses)*3 + int(bPresses)
		}
	}

	fmt.Println(total)
}

func getInput() []ClawMachine {
	clawMachines := []ClawMachine{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ax, ay := parseButton(s.Text())
		s.Scan()
		bx, by := parseButton(s.Text())
		s.Scan()
		px, py := parsePosition(s.Text())
		s.Scan()
		clawMachines = append(
			clawMachines,
			ClawMachine{ax, ay, bx, by, px, py},
		)
	}
	return clawMachines
}

func parseButton(s string) (int, int) {
	s = strings.ReplaceAll(s, ",", "")
	fields := strings.Fields(s)
	xStr, yStr := fields[2], fields[3]
	_, xStr, _ = strings.Cut(xStr, "+")
	_, yStr, _ = strings.Cut(yStr, "+")
	return atoi(xStr), atoi(yStr)
}

func parsePosition(s string) (int, int) {
	s = strings.ReplaceAll(s, ",", "")
	fields := strings.Fields(s)
	xStr, yStr := fields[1], fields[2]
	_, xStr, _ = strings.Cut(xStr, "=")
	_, yStr, _ = strings.Cut(yStr, "=")
	return atoi(xStr), atoi(yStr)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
