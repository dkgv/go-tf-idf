package go_tf_idf

import (
	"testing"
)

func TestCosineComparator(t *testing.T) {
	tests := []struct {
		vec1 []float64
		vec2 []float64
		want float64
	}{
		{
			vec1: []float64{1, 1, 1, 1, 0, 1, 1, 2},
			vec2: []float64{1, 0, 0, 1, 1, 0, 1, 0},
			want: 0.4743416490252569,
		},
		{
			vec1: []float64{0.2, 0.2, 0.4, 0.2, 0, 0},
			vec2: []float64{0.14285714285714285, 0.14285714285714285, 0, 0, 0.2857142857142857, 0.42857142857142855},
			want: 0.19518001458970657,
		},
		{
			vec1: []float64{3, 4},
			vec2: []float64{4, 3},
			want: 0.96,
		},
		{
			vec1: []float64{4, 2, 0, 0, 1, 1, 1, 1, 1, 0, 3, 1, 0, 3},
			vec2: []float64{6, 2, 2, 2, 3, 3, 3, 3, 3, 2, 5, 8, 3, 2},
			want: 0.7694486109646872,
		},
	}
	for _, tt := range tests {
		t.Run("cosine comparison", func(t *testing.T) {
			if got := CosineComparator(tt.vec1, tt.vec2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
