package cap4quicksort

import "testing"

func TestSum(t *testing.T)  {
	list := []int{1, 2, 3, 4, 5}

	got := Sum(list)
	want := 15

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
