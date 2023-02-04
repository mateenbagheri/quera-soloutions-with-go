package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Run("try with 0, 0", func(t *testing.T) {
		got := binomial(0, 0)
		want := 1
		assertFactResult(t, got, want)
	})

	t.Run("try with 1, 0", func(t *testing.T) {
		got := binomial(1, 0)
		want := 1
		assertFactResult(t, got, want)
	})

	t.Run("try with actual case", func(t *testing.T) {
		got := binomial(3, 1)
		want := 3
		assertFactResult(t, got, want)
	})
}

func assertFactResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d, but instead, got %d", want, got)
	}
}
