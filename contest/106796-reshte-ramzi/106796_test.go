package main

import "testing"

func TestEncrypt(t *testing.T) {
	testCases := []struct {
		name        string
		inputK      int
		inputN      int
		inputString string
		expected    string
	}{
		{
			name:        "test case no 1",
			inputN:      3,
			inputK:      1,
			inputString: "abz",
			expected:    "abc",
		},
		{
			name:        "test case no 2",
			inputN:      4,
			inputK:      5,
			inputString: "abcd",
			expected:    "ifgh",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := encrypt(tc.inputN, tc.inputK, tc.inputString)
			if got != tc.expected {
				t.Errorf("got %q but expected %q", got, tc.expected)
			}
		})
	}
}
