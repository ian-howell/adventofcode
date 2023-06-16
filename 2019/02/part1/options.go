package main

type IntCodeOption func(*IntCode)

func WithDebug() IntCodeOption {
	return func(ic *IntCode) {
		ic.Debug = true
	}
}

func WithProgram(program []int) IntCodeOption {
	return func(ic *IntCode) {
		ic.Memory = program
	}
}
