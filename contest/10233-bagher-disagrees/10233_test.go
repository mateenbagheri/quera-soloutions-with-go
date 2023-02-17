package main

import (
	"reflect"
	"testing"
)

func TestHasBiggerSmall(t *testing.T) {
	t.Run("has bigger small sample", func(t *testing.T) {
		wanted := "165"
		got, _ := hasBiggerSmall("156")
		assertHasBiggerSmallValue(t, wanted, *got)
	})

	t.Run("does not have bigger small sample", func(t *testing.T) {
		wanted := "0"
		got, _ := hasBiggerSmall("330")
		assertHasBiggerSmallValue(t, wanted, *got)
	})

	t.Run("third sample", func(t *testing.T) {
		wanted := "71127"
		got, _ := hasBiggerSmall("27711")
		assertHasBiggerSmallValue(t, wanted, *got)
	})

}

func TestGeneratePerms(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "single character",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "two characters",
			input:    "ab",
			expected: []string{"ab", "ba"},
		},
		{
			name:     "three characters",
			input:    "abc",
			expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := generatePerms(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func assertHasBiggerSmallValue(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %s but instead got %s", want, got)
	}
}
