package codility

import "testing"

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected got %d want %d", got, want)
	}
}