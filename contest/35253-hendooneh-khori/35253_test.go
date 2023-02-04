package main

import "testing"

func TestChooseLeast(t *testing.T) {
	t.Run("choose least 1 variable", func(t *testing.T) {
		gotIdx, got := chooseLeast(10)
		wantIdx, want := 0, 10
		assertChooseLeastResult(t, got, gotIdx, want, wantIdx)
	})

	t.Run("choose least more than 1 variables", func(t *testing.T) {
		gotIdx, got := chooseLeast(10, 2, 40)
		wantIdx, want := 1, 2
		assertChooseLeastResult(t, got, gotIdx, want, wantIdx)
	})
}

func TestLastWatermelonCalc(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		got := lastWatermelonCalc([]int{4, 3, 1, 5, 2}, 5)
		want := 4
		assertIndexReturned(t, got, want)
	})

	t.Run("sample 2", func(t *testing.T) {
		got := lastWatermelonCalc([]int{2, 4, 5, 1, 3}, 5)
		want := 3
		assertIndexReturned(t, got, want)
	})
}

func assertChooseLeastResult(t testing.TB, got, gotIdx, want, wantIdx int) {
	t.Helper()
	if got != want {
		t.Errorf("value err: wanted %d but instead, got %d", want, got)
	}
	if gotIdx != wantIdx {
		t.Errorf("index err: wanted %d but instead got %d", wantIdx, gotIdx)
	}
}

func assertIndexReturned(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d, but instead we got %d", want, got)
	}
}
