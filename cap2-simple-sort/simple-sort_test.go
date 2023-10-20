package cap2simplesort

import (
	"testing"
)

/* func TestSort(t *testing.T) {
	t.Run("sort list", func(t *testing.T) {
		list := []int{3, 1, 7, 0}

		got := Sort(list)
		want := []int{0, 1, 3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})
}
*/
func TestFindSmaller(t *testing.T) {
	list := []int{7, 2, 0, 8, 1}

	got := FindSmaller(list)
	want := 0

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}