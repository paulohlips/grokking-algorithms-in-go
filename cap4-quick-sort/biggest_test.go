package cap4quicksort

import "testing"

func TestBiggest(t *testing.T) {
	list := []int{1, -3, 9, 2, -100}

	got := Biggest(list)
	want := 9

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}