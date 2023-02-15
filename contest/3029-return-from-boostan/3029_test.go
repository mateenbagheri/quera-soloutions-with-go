package main

import "testing"

func TestLeastDistanceDirection(t *testing.T) {
	t.Run("first case", func(t *testing.T) {
		wanted := rightDir
		got := leastDistanceDirection(location{2, 2}, location{3, 1})
		assertDirection(t, wanted, got)
	})
}

func assertDirection(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q but instead we got %q", want, got)
	}
}
