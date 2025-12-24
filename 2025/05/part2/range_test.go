package main

import (
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		ranges   []string
		newRange Range
		expected []string
	}{
		"an empty RangeCollection accepts a new Range": {
			ranges:   []string{},
			newRange: NewRange("5-10"),
			expected: []string{"5-10"},
		},
		"a disjoint Range is added to the back": {
			ranges:   []string{"5-10"},
			newRange: NewRange("15-20"),
			expected: []string{"5-10", "15-20"},
		},
		"a disjoint Range is added to the back of a longer list": {
			ranges:   []string{"5-10", "12-16", "18-20"},
			newRange: NewRange("25-30"),
			expected: []string{"5-10", "12-16", "18-20", "25-30"},
		},
		"a disjoint Range is added to the front": {
			ranges:   []string{"15-20"},
			newRange: NewRange("5-10"),
			expected: []string{"5-10", "15-20"},
		},
		"a disjoint Range is added to the front of a longer list": {
			ranges:   []string{"12-16", "18-20", "25-30"},
			newRange: NewRange("5-10"),
			expected: []string{"5-10", "12-16", "18-20", "25-30"},
		},
		"a disjoint range is added to the middle of a list": {
			ranges:   []string{"5-10", "12-16", "25-30"},
			newRange: NewRange("18-20"),
			expected: []string{"5-10", "12-16", "18-20", "25-30"},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			rc := RangeCollection{}
			for _, r := range test.ranges {
				rc.ranges = append(rc.ranges, NewRange(r))
			}

			expected := RangeCollection{}
			for _, r := range test.expected {
				expected.ranges = append(expected.ranges, NewRange(r))
			}

			rc.Add(test.newRange)
			if !slices.Equal(rc.ranges, expected.ranges) {
				t.Logf("%v failed; expected:\n%v\ngot:\n%v", testName, expected, rc)
				t.Fail()
			}
		})
	}
}
