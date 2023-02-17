package main

import "testing"

func TestGetIsOutside(t *testing.T) {
	t.Run("first sample", func(t *testing.T) {
		wanted := true
		got := getIsOutSide(point{0, 5}, 5, point{0, 0})
		assertBoolean(t, wanted, got)
	})

	t.Run("second sample", func(t *testing.T) {
		wanted := false
		got := getIsOutSide(point{0, 5}, 5, point{5, 6})
		assertBoolean(t, wanted, got)
	})

	t.Run("third sample", func(t *testing.T) {
		wanted := false
		got := getIsOutSide(point{0, 5}, 5, point{-5, 3})
		assertBoolean(t, wanted, got)
	})

	t.Run("four sample", func(t *testing.T) {
		wanted := true
		got := getIsOutSide(point{0, 5}, 5, point{3, 3})
		assertBoolean(t, wanted, got)
	})
}

func assertBoolean(t testing.TB, want, got bool) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %t but instead we got %t", want, got)
	}
}
