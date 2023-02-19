package main

import (
	"testing"
)

// 5
// 1 5 2 4 3

func TestCaktusFlower(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		input    int
	}{
		{
			name:     "sample one",
			expected: 1,
			input:    1,
		},
		{
			name:     "sample two",
			expected: 1,
			input:    5,
		},
		{
			name:     "sample three",
			expected: 2,
			input:    2,
		},
		{
			name:     "sample four",
			expected: 1,
			input:    4,
		},
		{
			name:     "sample five",
			expected: 3,
			input:    3,
		},
	}
	for _, test := range tests {
		t.Run("test everything", func(t *testing.T) {
			got := cactusFlower(test.input)
			assertCactusFlower(t, test.expected, got)
		})
	}
}

func assertCactusFlower(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %d, but instead got %d", want, got)
	}
}
