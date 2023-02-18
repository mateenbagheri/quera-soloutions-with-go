package main

import "testing"

func TestGetBank(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		input    string
	}{
		{
			name:     "first sample: rooz zoj",
			expected: "perspolis",
			input:    "shanbe",
		},
		{
			name:     "second sample: rooz fard",
			expected: "bahman",
			input:    "seshanbe",
		},
		{
			name:     "third sample: tatil day",
			expected: "tatil",
			input:    "jome",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getBank(test.input)
			assertGetBank(t, got, test.expected)
		})
	}
}

func assertGetBank(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q, but got %q", want, got)
	}
}