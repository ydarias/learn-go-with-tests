package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, got, expected)
	})
}
