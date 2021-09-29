package go_tf_idf

import "math"

func CosineComparator(vec1, vec2 []float64) float64 {
	dot := float64(0)
	for i := range vec1 {
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
