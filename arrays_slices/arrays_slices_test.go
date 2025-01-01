package arrays_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertInt(t, got, want)
	})

	t.Run("collection of any numbers with SumWithRange", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumWithRange(numbers)
		want := 6
		assertInt(t, got, want)
	})

}


func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected got %q want %q", got, want)
	}
}