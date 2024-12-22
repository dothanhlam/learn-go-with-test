package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("given a and 5, then got aaaaa", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assert(t, repeated, expected)
	})

	t.Run("given a and 0, then got a", func (t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "a"
		assert(t, repeated, expected)
	})

	t.Run("given b and 2, then got bb", func(t *testing.T) {
		repeated := Repeat("b", 2)
		expected := "bb"
		assert(t, repeated, expected)
	})

	t.Run("given c and 3, then got ccc", func(t *testing.T) {
		repeated := IterationAsWhile("c", 3)
		expected := "ccc"
		assert(t, repeated, expected)
	})
}


func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected got %q want %q", got, want)
	}
}