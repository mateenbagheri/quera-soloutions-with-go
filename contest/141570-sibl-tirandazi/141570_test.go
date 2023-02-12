package main

import "testing"

func TestEvaluateHitpoint(t *testing.T) {
	t.Run("A white case", func(t *testing.T) {
		got := evaluateHit(3)
		want := whitePoint
		assertHitPointResult(t, want, got)
	})

	t.Run("A black case", func(t *testing.T) {
		got := evaluateHit(10)
		want := blackPoint
		assertHitPointResult(t, want, got)
	})

	t.Run("A close black case", func(t *testing.T) {
		got := evaluateHit(7)
		want := blackPoint
		assertHitPointResult(t, want, got)
	})

	t.Run("An out case", func(t *testing.T) {
		got := evaluateHit(0)
		want := outPoint
		assertHitPointResult(t, want, got)
	})
}

func assertHitPointResult(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q but instead we got %q", want, got)
	}
}
