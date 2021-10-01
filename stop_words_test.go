package go_tf_idf

import "testing"

func TestStopWords_Matches(t *testing.T) {
	type fields struct {
		List    map[string]bool
		Filters []Filter
	}
	tests := []struct {
		name            string
		fields          fields
		wordAndExpected map[string]bool
	}{
		{
			name: "test list matches",
			fields: fields{
				List: map[string]bool{
					"word": true,
				},
				Filters: []Filter{},
			},
			wordAndExpected: map[string]bool{
				"word": true,
				"asdf": false,
			},
		},
		{
			name: "test filter matches",
			fields: fields{
				List: map[string]bool{},
				Filters: []Filter{
					func(s string) bool {
						return len(s) == 1
					},
				},
			},
			wordAndExpected: map[string]bool{
				"":   false,
				"a":  true,
				"ab": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := StopWords{
				List:    tt.fields.List,
				Filters: tt.fields.Filters,
			}
			for word, want := range tt.wordAndExpected {
				if got := w.Matches(word); got != want {
					t.Errorf("Matches() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestStopWords_AddWord(t *testing.T) {
	w := StopWords{List: map[string]bool{}}
	w.AddWord("test")
	if _, ok := w.List["test"]; !ok {
		t.Errorf("AddWord() = %v, want %v", false, true)
	}
}

func TestStopWords_AddWords(t *testing.T) {
	w := StopWords{List: map[string]bool{}}
	w.AddWords([]string{"test1", "test2"})
	if _, ok := w.List["test1"]; !ok {
		t.Errorf("AddWord() = %v, want %v", false, true)
	}
	if _, ok := w.List["test2"]; !ok {
		t.Errorf("AddWord() = %v, want %v", false, true)
	}
}

func TestStopWords_AddFilter(t *testing.T) {
	w := StopWords{Filters: []Filter{}}
	filter := func(s string) bool {
		return true
	}
	w.AddIgnoreFilter(filter)
	if len(w.Filters) != 1 {
		t.Errorf("AddIgnoreFilter() = %v, want %v", len(w.Filters), 1)
	}
}
