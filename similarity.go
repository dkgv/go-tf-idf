package go_tf_idf

import "math"

func CosineComparator(vec1, vec2 []float64) float64 {
	// Instead of padding with 0s we trim
	min := int(math.Min(float64(len(vec1)), float64(len(vec2))))
	dot := float64(0)
	for i := 0; i < min; i++ {
		dot += vec1[i] * vec2[i]
	}

	magnitude := func(vec []float64) float64 {
		sum := float64(0)
		for _, xi := range vec {
			sum += xi * xi
		}
		return math.Sqrt(sum)
	}

	m1 := magnitude(vec1)
	m2 := magnitude(vec2)
	alpha := dot / (m1 * m2)
	return alpha
}
