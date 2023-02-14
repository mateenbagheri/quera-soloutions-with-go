package main

import (
	"reflect"
	"testing"
)

func TestDateDiff(t *testing.T) {
	t.Run("first sample", func(t *testing.T) {
		wanted := 1
		got := dateDiff(date{1, 1}, date{2, 1})
		assertDateDiffResult(t, wanted, got)
	})

	t.Run("second sample", func(t *testing.T) {
		wanted := 31
		got := dateDiff(date{30, 6}, date{30, 7})
		assertDateDiffResult(t, wanted, got)
	})

	t.Run("second sample", func(t *testing.T) {
		wanted := 0
		got := dateDiff(date{30, 2}, date{30, 2})
		assertDateDiffResult(t, wanted, got)
	})
}

func TestSortDate(t *testing.T) {
	t.Run("first is the result sample", func(t *testing.T) {
		got, _ := sortDate(date{12, 2}, date{2, 1})
		want := date{12, 2}
		assertSortDateResult(t, want, got)
	})

	t.Run("seond is the result sample", func(t *testing.T) {
		got, _ := sortDate(date{1, 2}, date{2, 1})
		want := date{1, 2}
		assertSortDateResult(t, want, got)
	})

	t.Run("equal sample", func(t *testing.T) {
		got, _ := sortDate(date{12, 2}, date{12, 2})
		want := date{12, 2}
		assertSortDateResult(t, want, got)
	})
}

func TestMonthToDay(t *testing.T) {
	t.Run("first 5 month", func(t *testing.T) {
		want := 62
		got := monthToDay(3)
		assertMonthToDayResult(t, want, got)
	})

	t.Run("between 6th and last", func(t *testing.T) {
		want := 216
		got := monthToDay(7)
		assertMonthToDayResult(t, want, got)
	})

	t.Run("last", func(t *testing.T) {
		want := (31 * 6) + (30 * 5)
		got := monthToDay(12)
		assertMonthToDayResult(t, want, got)
	})
}

func assertMonthToDayResult(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %d but instead, we got %d", want, got)
	}
}

func assertDateDiffResult(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %d but instead, we got %d", want, got)
	}
}

func assertSortDateResult(t testing.TB, want, got date) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("there is differance in dates. wanted %+v and got %+v", want, got)
	}
}
