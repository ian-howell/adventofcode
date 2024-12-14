package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		x, y                  LinearEquation
		a, b                  float64
		integerSolutionExists bool
	}{
		{
			x: LinearEquation{94, 22, 8400},
			y: LinearEquation{34, 67, 5400},
			a: 80, b: 40,
			integerSolutionExists: true,
		},
		{
			x: LinearEquation{26, 67, 12748},
			y: LinearEquation{66, 21, 12176},
			a: 0, b: 0,
			integerSolutionExists: false,
		},
		{
			x: LinearEquation{17, 84, 7870},
			y: LinearEquation{86, 37, 6450},
			a: 38, b: 86,
			integerSolutionExists: true,
		},
		{
			x: LinearEquation{69, 27, 18641},
			y: LinearEquation{23, 71, 10279},
			a: 0, b: 0,
			integerSolutionExists: false,
		},
	}

	for _, test := range tests {
		a, b, exists := Solve(test.x, test.y)
		if exists != test.integerSolutionExists || a != test.a || b != test.b {
			t.Errorf("Expected Solve(%+v, %+v) = (%f, %f, %v), got (%f, %f, %v)",
				test.x, test.y, test.a, test.b, test.integerSolutionExists,
				a, b, exists)
		}
	}
}
