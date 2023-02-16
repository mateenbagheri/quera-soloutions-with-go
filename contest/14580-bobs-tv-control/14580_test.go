package main

import "testing"

func TestFinalChannelCalculator(t *testing.T) {
	t.Run("first sample", func(t *testing.T) {
		want := "carl"
		got := finalChannelCalculator(5, 2, 5, []string{
			"bob",
			"carl",
			"kevin",
			"phil",
			"tim",
		})
		// carl > kevin > phil > tim > bob > carl
		assertFinalChannel(t, want, got)
	})

	t.Run("second sample", func(t *testing.T) {
		want := "bob"
		got := finalChannelCalculator(5, 1, 10, []string{
			"bob",
			"carl",
			"kevin",
			"phil",
			"tim",
		})
		// carl > kevin > phil > tim > bob > carl
		assertFinalChannel(t, want, got)
	})
}

func assertFinalChannel(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q but instead got %s", want, got)
	}
}
