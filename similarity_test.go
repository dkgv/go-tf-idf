package go_tf_idf

import (
	"strconv"
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
			vec1: []float64{3, 4},
			vec2: []float64{4, 3},
			want: 0.96,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := CosineComparator(tt.vec1, tt.vec2); got != tt.want {
				t.Errorf("Cosine() = %v, want %v", got, tt.want)
			}
		})
	}
}
