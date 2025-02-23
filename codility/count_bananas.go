package codility

import "math"

func CountBananas(s string) int {
	B := 0
	A := 0
	N := 0
	for _, char := range s {
		switch char {
		case 'B':
			B++
		case 'A':
			A++
		case 'N':
			N++
		}
	}

	return int(math.Min(float64(B), math.Min(math.Floor(float64(A)/3), math.Floor(float64(N)/2))))
}