package main

import "testing"

func TestPronunciationCombinationCalculator(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := pronunciationCombinationCalculator("mammad")
		want := 4
		assertCalculatorResults(t, got, want)
	})

	t.Run("case 2", func(t *testing.T) {
		got := pronunciationCombinationCalculator("anvari")
		want := 8
		assertCalculatorResults(t, got, want)
	})
	
	t.Run("case 3", func(t *testing.T) {
		got := pronunciationCombinationCalculator("sghrwq")
		want := 1
		assertCalculatorResults(t, got, want)
	})
}

func assertCalculatorResults(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d, but instead, got %d", want, got)
	}
}
