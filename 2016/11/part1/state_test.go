package main

import "testing"

func TestNextState(t *testing.T) {
	initialState := NewState()
	initialState.ElevatorFloor = 1
	initialState.Generators[1].Add(Generator("A"))
	initialState.Generators[1].Add(Generator("B"))
	initialState.Microchips[1].Add(Microchip("A"))
	initialState.Microchips[1].Add(Microchip("B"))

	nextStates := initialState.Next()
	if len(nextStates) != 12 {
		t.Errorf("Expected 12 next states, only got %d", len(nextStates))

		t.Logf("Initial State:\n%v", initialState)

		for i, state := range nextStates {
			t.Logf("Next State #%v:\n%v", i, state)
		}
	}

}
