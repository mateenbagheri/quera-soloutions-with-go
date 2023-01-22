package main

import "testing"

func TestDiffLetterCounter(t *testing.T) {
	t.Run("string with no repeated word", func(t *testing.T) {
		got := diffLetterCounter("ali")
		want := 3
		assertDifferenceCorrectness(t, got, want)
	})

	t.Run("string with repeated words", func(t *testing.T) {
		got := diffLetterCounter("abbas")
		want := 3
		assertDifferenceCorrectness(t, got, want)
	})

	t.Run("string with all letters the same", func(t *testing.T) {
		got := diffLetterCounter("aaaa")
		want := 1
		assertDifferenceCorrectness(t, got, want)
	})
}

func TestMostDifference(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		got := mostDifference([]string{"ali", "karim", "abbas", "mohammad"})
		want := 5
		assertMostDifferenceCorrectness(t, got, want)
	})
}

func assertDifferenceCorrectness(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q but instead, got %q", want, got)
	}
}

func assertMostDifferenceCorrectness(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d but instead, got %d", want, got)
	}
}