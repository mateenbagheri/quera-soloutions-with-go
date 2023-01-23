package main

import "testing"

func TestGetBackTimeCalc(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		trafficLights := [][]int{
			{3, 5, 5},
			{5, 2, 2},
		}
		l := 10

		want := 12
		got := getBackTimeCalc(l, trafficLights)
		assertGetBackTimeDuration(t, want, got)
	})
	t.Run("second example", func(t *testing.T) {
		trafficLights := [][]int{
			{7, 13, 5},
			{14, 4, 4},
			{15, 3, 10},
			{25, 1, 1},
		}
		l := 30

		want := 36
		got := getBackTimeCalc(l, trafficLights)
		assertGetBackTimeDuration(t, want, got)
	})
}

func assertGetBackTimeDuration(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %d but instead, we got %d", want, got)
	}
}
