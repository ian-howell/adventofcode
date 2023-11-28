package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Generator string

func (g Generator) String() string { return string(g) + "G" }

func (g Generator) Pair() Microchip { return Microchip(string(g[0])) }

type Microchip string

func (m Microchip) String() string  { return string(m) + "M" }
func (m Microchip) Pair() Generator { return Generator(string(m[0])) }

type State struct {
	ElevatorFloor int
	Generators    []Set[Generator]
	Microchips    []Set[Microchip]
}

func (s State) String() string {
	sb := strings.Builder{}
	for i := 4; i > 0; i-- {
		sb.WriteString(fmt.Sprintf("F%d ", i))
		if s.ElevatorFloor+1 == i {
			sb.WriteString("E ")
		} else {
			sb.WriteString(". ")
		}
		for g := range s.Generators[i-1] {
			sb.WriteString(fmt.Sprintf("%v ", g))
		}
		for m := range s.Microchips[i-1] {
			sb.WriteString(fmt.Sprintf("%v ", m))
		}
		if i > 1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func NewState() State {
	s := State{
		ElevatorFloor: 0,
		Generators:    make([]Set[Generator], 4),
		Microchips:    make([]Set[Microchip], 4),
	}
	for i := 0; i < 4; i++ {
		s.Generators[i] = NewSet[Generator]()
		s.Microchips[i] = NewSet[Microchip]()
	}
	return s
}

func (s State) Next() []State {
	// The elevator must change floors
	// At least one Generator or Microchip must change floors
	// At most 2 Generators and Microchips may move floors
	//   So effectively, we can take (G), (M), (G, G), (M, M), or (G, M)
	nextStates := []State{}

	for _, floorDiff := range []int{-1, 1} {
		nextFloor := s.ElevatorFloor + floorDiff
		if nextFloor < 0 || nextFloor >= 4 {
			continue
		}

		var s1, s2 State

		// 1 Generator
		for i, g1 := range s.Generators[s.ElevatorFloor].Members() {
			s1 = s.Copy()
			s1.ElevatorFloor = nextFloor
			s1.Generators[s.ElevatorFloor].Remove(g1)
			s1.Generators[s1.ElevatorFloor].Add(g1)
			if s1.IsValid() {
				nextStates = append(nextStates, s1)
			}

			// 2 Generators
			for _, g2 := range s1.Generators[s.ElevatorFloor].Members()[i:] {
				s2 = s1.Copy()
				s2.ElevatorFloor = nextFloor
				s2.Generators[s.ElevatorFloor].Remove(g2)
				s2.Generators[s2.ElevatorFloor].Add(g2)
				if s2.IsValid() {
					nextStates = append(nextStates, s2)
				}
			}

			// 1 Generator and 1 Microchip
			for m2 := range s1.Microchips[s.ElevatorFloor] {
				s2 = s1.Copy()
				s2.ElevatorFloor = nextFloor
				s2.Microchips[s.ElevatorFloor].Remove(m2)
				s2.Microchips[s2.ElevatorFloor].Add(m2)
				if s2.IsValid() {
					nextStates = append(nextStates, s2)
				}
			}
		}

		// 1 Microchip
		for i, m1 := range s.Microchips[s.ElevatorFloor].Members() {
			s1 = s.Copy()
			s1.ElevatorFloor = nextFloor
			s1.Microchips[s.ElevatorFloor].Remove(m1)
			s1.Microchips[s1.ElevatorFloor].Add(m1)
			if s1.IsValid() {
				nextStates = append(nextStates, s1)
			}

			// 2 Microchips
			for _, m2 := range s1.Microchips[s.ElevatorFloor].Members()[i:] {
				s2 = s1.Copy()
				s2.ElevatorFloor = nextFloor
				s2.Microchips[s.ElevatorFloor].Remove(m2)
				s2.Microchips[s2.ElevatorFloor].Add(m2)
				if s2.IsValid() {
					nextStates = append(nextStates, s2)
				}
			}
		}
	}

	return nextStates
}

func (s State) IsValid() bool {
	for i := 0; i < 4; i++ {
		for m := range s.Microchips[i] {
			if !s.Generators[i].Contains(m.Pair()) && len(s.Generators[i]) > 0 {
				return false
			}
		}
	}

	return true
}

func readState() State {
	state := NewState()
	floorNo := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		tokens := tokenize(s.Text())
		for _, token := range tokens {
			element, kind, _ := strings.Cut(token, " ")
			elementSymbol := string(strings.ToUpper(element)[0])
			switch kind {
			case "microchip":
				state.Microchips[floorNo].Add(Microchip(elementSymbol))
			case "generator":
				state.Generators[floorNo].Add(Generator(elementSymbol))
			}
		}
		floorNo++
	}

	return state
}

func tokenize(s string) []string {
	s = strings.TrimRight(s, ".")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, " and", "")
	words := strings.Split(s, " a ")[1:]
	return words
}

func (s State) Copy() State {
	copied := NewState()
	copied.ElevatorFloor = s.ElevatorFloor
	for i := 0; i < 4; i++ {
		copied.Generators[i] = NewSet[Generator]()
		for g := range s.Generators[i] {
			copied.Generators[i].Add(g)
		}

		copied.Microchips[i] = NewSet[Microchip]()
		for g := range s.Microchips[i] {
			copied.Microchips[i].Add(g)
		}
	}
	return copied
}

func (s State) Hash() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%d:", s.ElevatorFloor))
	for i := 0; i < 4; i++ {
		generators := []string{}
		for g := range s.Generators[i] {
			generators = append(generators, g.String())
		}
		sort.Strings(generators)
		sb.WriteString(strings.Join(generators, ","))

		sb.WriteString(";")

		microchips := []string{}
		for m := range s.Microchips[i] {
			microchips = append(microchips, m.String())
		}
		sort.Strings(microchips)
		sb.WriteString(strings.Join(microchips, ","))

		if i < 3 {
			sb.WriteString("-")
		}
	}
	return sb.String()
}

func NewStateFromHash(hash string) State {
	state := NewState()

	floor, hash, _ := strings.Cut(hash, ":")
	state.ElevatorFloor = atoi(floor)

	floors := strings.Split(hash, "-")
	for i, floor := range floors {
		generators, microchips, _ := strings.Cut(floor, ";")
		if len(generators) > 0 {
			for _, g := range strings.Split(generators, ",") {
				state.Generators[i].Add(Generator(g[0]))
			}
		}
		if len(microchips) > 0 {
			for _, m := range strings.Split(microchips, ",") {
				state.Microchips[i].Add(Microchip(m[0]))
			}
		}
	}

	return state
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}
