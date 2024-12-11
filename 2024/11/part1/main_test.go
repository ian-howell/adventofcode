package main

import "testing"

func TestSplit(t *testing.T) {
	tests := []struct {
		input, left, right int
	}{
		// I only really care about the cases where the input has an even number of digits
		{12, 1, 2},
		{1234, 12, 34},
		{123456, 123, 456},

		// The prompt specifically calls this scenario out
		{1000, 10, 0},
	}

	for _, test := range tests {
		l, r := split(test.input)
		if l != test.left || r != test.right {
			t.Errorf("expected (%d, %d), got (%d, %d)",
				test.left, test.right, l, r)
		}
	}
}
