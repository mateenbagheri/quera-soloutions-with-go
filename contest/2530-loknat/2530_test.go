package main

import "testing"

func TestLoknatCaseCounter(t *testing.T) {
	t.Run("sample case 1", func(t *testing.T) {
		wanted := 4
		got := loknatCaseCounter("FILIPEK")
		assertCorrectResult(t, wanted, got)
	})
}

func assertCorrectResult(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d but instead we got %d", want, got)
	}
}
