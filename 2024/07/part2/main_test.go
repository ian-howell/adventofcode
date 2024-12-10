package main

import (
	"testing"
)

func TestTrimSuffix(t *testing.T) {
	tests := []struct {
		num, suffix, expected int
	}{
		// Successfully trimmed
		{12345, 5, 1234},
		{12345, 45, 123},
		{12345, 345, 12},
		{12345, 2345, 1},
		{12345, 12345, 0},
		// Should work for negatives too
		{-12345, 5, -1234},
		{-12345, 45, -123},
		{-12345, 345, -12},
		{-12345, 2345, -1},
		{-12345, 12345, 0},

		// Special 0 case
		{0, 0, 0},

		// TrimSuffix returns num when suffix is not actually a suffix
		{123, 45, 123},
		{123, 12345, 123},
		{123, 999123, 123},
		{123, 0, 123},

		// Negative suffixes don't make sense...
		{12345, -5, 12345},
		{12345, -45, 12345},
		{12345, -345, 12345},
		{12345, -2345, 12345},
		{-12345, -5, -12345},
		{-12345, -45, -12345},
		{-12345, -345, -12345},
		{-12345, -2345, -12345},
		// ... Unless num == suffix
		{-12345, -12345, 0},
	}

	for _, test := range tests {
		num := trimSuffix(test.num, test.suffix)
		if num != test.expected {
			t.Errorf("expected (%d), got (%d)", test.expected, num)
		}
	}
}
