package main

import (
	"fmt"
	"testing"
)

func TestIsInvalid(t *testing.T) {
	tests := []struct {
		input     int
		isInvalid bool
	}{
		{1, false},
		{11, true},
		{111, false},
		{1111, true},
		{11111, false},
		{111111, true},

		{11, true},
		{22, true},
		{99, true},
		{1010, true},
		{1188511885, true},
		{222222, true},
		{446446, true},
		{38593859, true},

		{12, false},
		{1234, false},
		{123456, false},
		{12345678, false},
		{1010101410, false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("input-%d-expect-%v", test.input, test.isInvalid), func(t *testing.T) {
			if got := isInvalid(test.input); got != test.isInvalid {
				t.Errorf("expected isInvalid(%d) == %v, got %v", test.input, test.isInvalid, got)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{15, 24, 25},
		{192_837, 918_273, 998_877},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("a-%d-b-%d-expect-%d", test.a, test.b, test.expected), func(t *testing.T) {
			if got := merge(test.a, test.b); got != test.expected {
				t.Errorf("expected merge(%d, %d) == %d, got %d", test.a, test.b, test.expected, got)
			}
		})
	}
}
