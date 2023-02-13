package main

import "testing"

func TestKeypadResult(t *testing.T) {
	t.Run("first sample", func(t *testing.T) {
		want := "dANGerY"
		got := mergeKeys([]string{
			"d",
			"CAPS",
			"a",
			"n",
			"g",
			"CAPS",
			"e",
			"r",
			"CAPS",
			"y",
		})
		assertCorrectKeyWordOutput(t, want, got)
	})

	t.Run("second sample", func(t *testing.T) {
		want := "zju"
		got := mergeKeys([]string{
			"z",
			"j",
			"u",
		})
		assertCorrectKeyWordOutput(t, want, got)
	})
}

func assertCorrectKeyWordOutput(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q but instead, we got %q", want, got)
	}
}
