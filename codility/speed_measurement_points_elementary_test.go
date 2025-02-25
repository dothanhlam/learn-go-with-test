package codility

import "testing"

func SpeedMeasurementPointElementaryTests(t *testing.T) {
	t.Run("Given S = '.>...', the function should return 3. The car will pass by three speed cameras to the right of it.", func(t *testing.T) {
		got := SpeedMeasurementPointElementary(".>...")
		want := 3
		assertEqual(t, got, want)
	})

	t.Run("Given S = '.>.<.>', the function should return 4. The first two cars will pass by two speed cameras each, and the third car will not pass by any.", func(t *testing.T) {
		got := SpeedMeasurementPointElementary(".>.<.>")
		want := 4
		assertEqual(t, got, want)
	})

	t.Run("Given S = '>>>.<<<', the function should return 6. Each car will pass one speed camera.", func(t *testing.T) {
		got := SpeedMeasurementPointElementary(">>>.<<<")
		want := 6
		assertEqual(t, got, want)
	})
}
