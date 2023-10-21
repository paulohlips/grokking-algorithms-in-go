package cap3recursion

import "testing"

func TestFactorial(t *testing.T)  {
	t.Run("5! = 120", func(t *testing.T) {
		got := Factorial(5)
		want := 120

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	})
}