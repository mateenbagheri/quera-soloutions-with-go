package main

import "testing"

func TestIsXLarge(t *testing.T) {
	tests := []struct {
		name     string
		input    [4]int
		expected bool
	}{
		{
			name:     "first sample: true",
			input:    [4]int{30, 34, 15, 34},
			expected: true,
		},
		{
			name:     "second sample: false",
			input:    [4]int{30, 34, 68, 33},
			expected: false,
		},
	}
	for _, test := range tests {
		got := isXLarge(
			test.input[0],
			test.input[1],
			test.input[2],
			test.input[3],
		)
		if got != test.expected {
			t.Errorf("wanted %v but instead got %v", test.expected, got)
		}
	}
}
