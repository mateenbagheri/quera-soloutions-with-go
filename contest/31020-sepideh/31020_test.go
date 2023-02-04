package main

import "testing"

func TestMaxDoorEntryCalculator(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		got := maxDoorEntryCalculator(9, 3)
		want := 3
		assertDoorEntryNumber(t, got, want)
	})

	t.Run("test case 2", func(t *testing.T) {
		got := maxDoorEntryCalculator(7, 4)
		want := 2
		assertDoorEntryNumber(t, got, want)
	})
}

func assertDoorEntryNumber(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %d but instead, we got %d", want, got)
	}
}
