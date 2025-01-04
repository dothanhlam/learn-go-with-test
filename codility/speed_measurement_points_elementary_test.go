package codility

import "testing"

func TestSuite(t *testing.T) {
	t.Run("speed measurement points", func(t *testing.T) {
		got := 0
		want := 0
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected got %d want %d", got, want)
	}
}