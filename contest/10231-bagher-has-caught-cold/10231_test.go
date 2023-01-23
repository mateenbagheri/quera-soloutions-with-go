package main

import "testing"

func TestHasacceptedWords(t *testing.T) {
	t.Run("true case", func(t *testing.T) {
		want := true
		got := hasAcceptedWords("MOLANA_RAND", acceptedWords)
		assertHasAcceptedWords(t, want, got)
	})

	t.Run("false case", func(t *testing.T) {
		want := false
		got := hasAcceptedWords("ENGSLVLWOE", acceptedWords)
		assertHasAcceptedWords(t, want, got)
	})
}

func TestCountHasacceptedWords(t *testing.T) {
	t.Run("third example", func(t *testing.T) {
		want := []int{1}
		got := getAcceptedWords([]string{
			"N-MOLANA1",
			"9A-UKOK",
			"SAYE-NTERP",
			"G-SANAEI",
			"RF-MOLLASADRA",
		}, acceptedWords)
		asserCountAcceptedWords(t, want, got)
	})

	t.Run("second example", func(t *testing.T) {
		want := []int{}
		got := getAcceptedWords([]string{
			"N321-HAFEEZ",
			"F3-B12I",
			"F-BI-12",
			"OVO-JE-FE",
			"KASHANI",
		}, acceptedWords)
		asserCountAcceptedWords(t, want, got)
	})

	t.Run("third example", func(t *testing.T) {
		want := []int{1, 3, 5}
		got := getAcceptedWords([]string{
			"Z7-MOLANA",
			"ZOND-007",
			"ZF-MOLANA18",
			"ZARICA-13",
			"Z3A-HAFEZLL",
		}, acceptedWords)
		asserCountAcceptedWords(t, want, got)
	})
}

func assertHasAcceptedWords(t testing.TB, want, got bool) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %t but instead I got %t", want, got)
	}
}

func asserCountAcceptedWords(t testing.TB, want, got []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Error("want and got have different lengths")
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("somethings wrong")
		}
	}
}
