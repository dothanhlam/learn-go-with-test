package codility

import (
	"sort"
	"math"
	)	

type Tree struct {
	X int
	Y int
}

func NewTree(x, y int) *Tree {
	return &Tree{x, y}
}


func WidestPath(X []int, Y []int) int {
	n := len(X)
	if n <= 1 {
		return 0
	}

	trees := make([]*Tree, n)
	for i := 0; i < n; i++ {
		trees[i] = NewTree(X[i], Y[i])
	}

	sort.Slice(trees, func(i, j int) bool {	
		return trees[i].X < trees[j].X
	})

	maxWidth := 0
	for i := 0; i < n - 1; i++ {
		width := trees[i+1].X - trees[i].X
		maxY := int(math.Max(float64(trees[i].Y), float64(trees[i+1].Y)))
		minY := int(math.Min(float64(trees[i].Y), float64(trees[i+1].Y)))

		if minY <= maxY && width > maxWidth {
			maxWidth = width
		}	
	}
	return maxWidth
}