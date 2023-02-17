package main

import (
	"testing"
)

func TestCheckHandsfree(t *testing.T) {
	tests := []struct {
		name     string
		input    [2]handsfree
		expected bool
	}{
		{
			name:     "first sample",
			input:    [2]handsfree{{"A", "A"}, {"B", "B"}},
			expected: true,
		},
		{
			name:     "second sample",
			input:    [2]handsfree{{"H", "U"}, {"I", "H"}},
			expected: true,
		},
		{
			name:     "third sample",
			input:    [2]handsfree{{"K", "F"}, {"F", "J"}},
			expected: true,
		},
		{
			name:     "forth sample",
			input:    [2]handsfree{{"O", "R"}, {"O", "G"}},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkHandsfree(test.input[0], test.input[1])
			assertCheckHandsfreeResult(t, test.expected, result)
		})
	}
}

func assertCheckHandsfreeResult(t testing.TB, want, got bool) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %v but instead, got %v", want, got)
	}
}
