package cap2simplesort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Helper()
	t.Run("sort list", func(t *testing.T) {
		list := []int{3, 1, 7, 0}

		got := Sort(list)
		want := []int{0, 1, 3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}

func TestFindSmaller(t *testing.T) {
	t.Helper()
	t.Run("find smaller", func(t *testing.T) {
		list := []int{7, 2, 0, 8, 1}

		got := FindSmaller(list)
		want := []int{0, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}