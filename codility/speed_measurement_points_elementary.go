package codility

func SpeedMeasurementPointElementary(S string) int {
	n := len(S)
	right := make([]int, n) // Create a slice (dynamically sized array)
	left := make([]int, n)  // of integers, initialized to 0
	rCars := 0
	lCars := 0

	for i := 0; i < n; i++ {
		if S[i] == '.' {
			right[i] = rCars
		} else if S[i] == '>' {
			rCars++
		}
	}

	for i := n - 1; i >= 0; i-- {
		if S[i] == '.' {
			left[i] = lCars
		} else if S[i] == '<' {
			lCars++
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += left[i] + right[i]
	}
	return ans
}