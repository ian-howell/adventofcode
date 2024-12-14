package main

import "math"

type LinearEquation struct {
	a, b, s float64
}

func (l LinearEquation) Times(c float64) LinearEquation {
	return LinearEquation{a: c * l.a, b: c * l.b, s: c * l.s}
}

func (l LinearEquation) Plus(m LinearEquation) LinearEquation {
	return LinearEquation{a: m.a + l.a, b: m.b + l.b, s: m.s + l.s}
}

func Solve(first, second LinearEquation) (a, b float64, integerSolutionExists bool) {
	oldFirst := first
	oldSecond := second

	first = first.Times(-second.a / first.a)
	second = second.Plus(first)

	// now we have:
	// a0 + b0 = s0
	//      b1 = s1

	second = second.Times(-first.b / second.b)
	first = first.Plus(second)

	// and now:
	// a0      = s0
	//      b1 = s1

	// This bit just scales both coefficients to 1
	first = first.Times(1 / first.a)
	second = second.Times(1 / second.b)

	a = math.Round(first.s)
	b = math.Round(second.s)
	integerSolutionExists = (a*oldFirst.a+b*oldFirst.b == oldFirst.s) &&
		(a*oldSecond.a+b*oldSecond.b == oldSecond.s)
	if !integerSolutionExists {
		a, b = 0, 0
	}
	return a, b, integerSolutionExists
}
