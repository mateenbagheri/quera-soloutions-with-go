package main

import "testing"

func TestAverageFinder(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    []int
	}{
		{
			name:     "sample one",
			expected: 3,
			input:    []int{1, 2, 3, 6},
		},
	}

	for _, test := range tests {
		got := averageFinder(test.input)
		if got != test.expected {
			t.Errorf("wanted %d but instead got %d", test.expected, got)
		}
	}
}

func TestTotalDistanceCalculator(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    []int
	}{
		{
			name:     "sample one",
			expected: 3,
			input:    []int{1, 2, 3, 6},
		},
	}

	for _, test := range tests {
		got := totalDistanceCalculator(test.input)
		if got != test.expected {
			t.Errorf("wanted %d but instead %d got", test.expected, got)
		}
	}
}
