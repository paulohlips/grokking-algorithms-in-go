package cap4quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		list := []int{1}

		got := QuickSort(list)
		want := []int{1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("random case", func(t *testing.T) {
		list := []int{1, 3, 2, 0, 80, 6}

		got := QuickSort(list)
		want := []int{0, 1, 2, 3, 6, 80}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}