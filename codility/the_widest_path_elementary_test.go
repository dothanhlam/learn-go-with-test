package codility

import "testing"

func TheWidestPathElementaryTests(t *testing.T) {
	t.Run("Given X=[1, 8, 7, 3, 4, 1, 8], Y=[6, 4, 1, 8, 5, 1, 7], the function should return 3.", func(t *testing.T) {
		got := WidestPath([]int{1, 8, 7, 3, 4 ,1, 8}, []int{6, 4, 1, 8, 5, 1, 7})
		want := 3
		assertEqual(t, got, want)
	})

	t.Run("Given X=[5, 5, 5, 7, 7, 7], Y=[3, 4, 5, 1, 3, 7], the function should return 2", func(t *testing.T) {
		got := WidestPath([]int{5, 5, 5, 7, 7, 7}, []int{3, 4, 5, 1, 3, 7})
		want := 2
		assertEqual(t, got, want)
	})

	t.Run("Given X=[6, 10, 1, 4, 3], Y=[2, 5, 3, 1, 6], the function should return 4.", func(t *testing.T) {
		got := WidestPath([]int{6, 10, 1, 4, 3}, []int{2, 5, 3, 1, 6})
		want := 4
		assertEqual(t, got, want)
	})

	t.Run("Given X=[4, 1, 5, 4], Y=[4, 5, 1, 3], the function should return 3.", func(t *testing.T) {
		got := WidestPath([]int{4, 1, 5, 4}, []int{4, 5, 1, 3})
		want := 3
		assertEqual(t, got, want)
	})
}
