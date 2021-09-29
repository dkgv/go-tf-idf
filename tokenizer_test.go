package go_tf_idf

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "Lowercase to two parts",
			s:    "two parts",
			want: []string{"two", "parts"},
		},
		{
			name: "Uppercase to lowercase two parts",
			s:    "TWO PARTS",
			want: []string{"two", "parts"},
		},
		{
			name: "One part",
			s:    "one",
			want: []string{"one"},
		},
		{
			name: "",
			s:    "",
			want: []string{},
		},
		{
			name: "One sentence with commas",
			s:    "sentence, one, sentence, two.",
			want: []string{"sentence", "one", "sentence", "two"},
		},
		{
			name: "Two sentences with period",
			s:    "sentence one. sentence two.",
			want: []string{"sentence", "one", "sentence", "two"},
		},
		{
			name: "One period",
			s:    ".",
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tokenize(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
