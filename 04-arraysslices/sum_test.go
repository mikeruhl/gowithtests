package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumSlice(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Summing 3 length", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6
		assertCorrectMessage(t, got, want)
	})

	t.Run("Summing 5 length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlice(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}
