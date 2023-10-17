package main

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("should finish with 1 iteration", func(t *testing.T) {
		list := []int{1 ,2, 3, 4, 5, 6, 7, 8}
		x := 4
		numIterations := 0

		got := BinarySearch(list, x)
		expected := numIterations

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})

	t.Run("should finish with 3 iteration to 8", func(t *testing.T) {
		list := []int{1 ,2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		x := 3
		numIterations := 3

		got := BinarySearch(list, x)
		expected := numIterations

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})
	t.Run("should finish with 3 iterations to 6", func(t *testing.T) {
		list := []int{0, 1 ,2, 3, 4, 5, 6, 7}
		x := 6
		numIterations := 2

		got := BinarySearch(list, x)
		expected := numIterations

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})

	t.Run("should finish with 4 iteration to 16", func(t *testing.T) {
		list := []int{1 ,2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		x := 16
		numIterations := 4

		got := BinarySearch(list, x)
		expected := numIterations

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})


	t.Run("should finish with 3 iterations to 0", func(t *testing.T) {
		list := []int{0, 1 ,2, 3, 4, 5, 6, 7}
		x := 0
		numIterations := 2

		got := BinarySearch(list, x)
		expected := numIterations

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})
}