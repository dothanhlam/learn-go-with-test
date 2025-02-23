package codility

import (
	"testing"
)

func CountBananasTests(t *testing.T) {
	t.Run("Given S = 'NAABXXAN', the function should return 1.", func(t *testing.T) {
		got := CountBananas("NAABXXAN")
		want := 1
		assertEqual(t, got, want)
	})

	t.Run("Given S = 'NAANAAXNABABYNNBZ', the function should return 2", func(t *testing.T) {
		got := CountBananas("NAANAAXNABABYNNBZ")
		want := 2
		assertEqual(t, got, want)
	})

	t.Run("Given S = 'QABAAAWOBL', the function should return 0", func(t *testing.T) {
		got := CountBananas("QABAAAWOBL")
		want := 0
		assertEqual(t, got, want)
	})
}

