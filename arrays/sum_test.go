package main

import "testing"

func TestSum(t *testing.T) {
	assertSameNumber := func(t *testing.T, got, expected int, numbers []int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d want %d given, %v", got, expected, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		assertSameNumber(t, got, expected, numbers)
	})
}
