package structs_methods_interfaces

import ("testing")


func TestSuite(t *testing.T) {
	t.Run("Perimeter", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		assertEqual(t, got, want)
	})

	t.Run("Area", func(t *testing.T) {
		got := Area(12.0, 6.0)
		want := 72.0

		assertEqual(t, got, want)
	})

	t.Run("Area", func(t *testing.T) {
		got := Area(10)
		want := 314.1592653589793

		assertEqual(t, got, want)
	}) 
}

func assertEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("expected got %f want %f", got, want)
	}
}