package main

import (
	"fmt"
	"testing"
)

func TestCalculateChecksum(t *testing.T) {
	tests := []struct {
		files    []File
		expected int
	}{{
		expected: 1928,
		files: []File{{
			ID:     0,
			Blocks: []Block{{Start: 0, End: 2}},
		}, {
			ID:     1,
			Blocks: []Block{{Start: 5, End: 8}},
		}, {
			ID:     2,
			Blocks: []Block{{Start: 11, End: 12}},
		}, {
			ID:     3,
			Blocks: []Block{{Start: 15, End: 18}},
		}, {
			ID:     4,
			Blocks: []Block{{Start: 19, End: 21}},
		}, {
			ID:     5,
			Blocks: []Block{{Start: 22, End: 26}},
		}, {
			ID: 6,
			Blocks: []Block{
				{Start: 18, End: 19},
				{Start: 21, End: 22},
				{Start: 26, End: 28},
			},
		}, {
			ID:     7,
			Blocks: []Block{{Start: 12, End: 15}},
		}, {
			ID: 8,
			Blocks: []Block{
				{Start: 4, End: 5},
				{Start: 8, End: 11},
			},
		}, {
			ID:     9,
			Blocks: []Block{{Start: 2, End: 4}},
		}},
	}}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			got := calculateChecksum(test.files)
			if got != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, got)
			}
		})
	}
}
