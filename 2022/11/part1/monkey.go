package main

import (
	"bufio"
	"os"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation func(int) int
	Test      int
	True      int
	False     int
}

func getMonkeys() []Monkey {
	var monkeys []Monkey
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var monkey Monkey
		s.Scan() // Eat the "Monkey i" line

		// Process the "Starting items"
		line := strings.ReplaceAll(s.Text(), ",", " ")
		parts := strings.Fields(line)
		for _, part := range parts[2:] {
			monkey.Items = append(monkey.Items, atoi(part))
		}
		s.Scan()

		// Process the "Operation"
		parts = strings.Fields(s.Text())
		val := atoi(parts[5])
		if parts[4] == "*" {
			if parts[5] == "old" {
				monkey.Operation = func(x int) int { return x * x }
			} else {
				monkey.Operation = func(x int) int { return x * val }
			}
		} else {
			if parts[5] == "old" {
				monkey.Operation = func(x int) int { return x + x }
			} else {
				monkey.Operation = func(x int) int { return x + val }
			}
		}
		s.Scan()

		// Process the "Test"
		parts = strings.Fields(s.Text())
		monkey.Test = atoi(parts[3])
		s.Scan()

		// Process the "If true"
		parts = strings.Fields(s.Text())
		monkey.True = atoi(parts[5])
		s.Scan()

		// Process the "If false"
		parts = strings.Fields(s.Text())
		monkey.False = atoi(parts[5])
		s.Scan()

		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func (m Monkey) HasItems() bool {
	return len(m.Items) > 0
}

func (m *Monkey) Pop() int {
	val := m.Items[0]
	m.Items = m.Items[1:]
	return val
}

func (m *Monkey) Push(x int) {
	m.Items = append(m.Items, x)
}
