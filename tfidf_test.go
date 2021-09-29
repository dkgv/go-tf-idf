package go_tf_idf

import (
	"reflect"
	"testing"
)

var doc1Hash = "83e4b1789306d3d1c99140df3827d600"
var doc1Content = "this is a a sample"
var doc1Freq = map[string]float64{
	"this":   1,
	"is":     1,
	"a":      2,
	"sample": 1,
}
var tokens1 = []string{"this", "is", "a", "sample"}
var tokenCount1 = 5
var doc1 = Document{Frequency: doc1Freq, UniqueTokens: tokens1, TotalTokenCount: tokenCount1}

var doc2Hash = "271559ec25268bb9bb2ad7fd8b4cf71a"
var doc2Content = "this is another another example example example"
var doc2Freq = map[string]float64{
	"this":    1,
	"is":      1,
	"another": 2,
	"example": 3,
}
var tokens2 = []string{"this", "is", "another", "example"}
var tokenCount2 = 7
var doc2 = Document{Frequency: doc2Freq, UniqueTokens: tokens2, TotalTokenCount: tokenCount2}

var docs = map[string]Document{doc1Hash: doc1, doc2Hash: doc2}

func TestDocument_GetVectors(t *testing.T) {
	tests := []struct {
		name  string
		doc   Document
		other Document
		want1 []float64
		want2 []float64
	}{
		{
			name:  "vector from doc0 and doc1",
			doc:   doc1,
			other: doc2,
			want1: []float64{0.2, 0.2, 0.4, 0.2, 0, 0},
			want2: []float64{0.14285714285714285, 0.14285714285714285, 0, 0, 0.2857142857142857, 0.42857142857142855},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.doc.GetVectors(tt.other)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("vec1 = %v, want %v", got1, tt.want1)
			}

			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("vec2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestTfIdf_GetDocument(t *testing.T) {
	tests := []struct {
		name        string
		getDocument string
		want        *Document
	}{
		{
			name:        "doc1",
			getDocument: "doc1",
			want:        &doc1,
		},
		{
			name:        "doc2",
			getDocument: "doc2",
			want:        &doc2,
		},
		{
			name:        "asdf = nil",
			getDocument: "asdf",
			want:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := TfIdf{documents: docs}
			if got := i.GetDocument(tt.getDocument); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_TermFrequency(t *testing.T) {
	tests := []struct {
		name       string
		frequency  map[string]float64
		term       string
		tokens     []string
		tokenCount int
		want       float64
	}{
		{
			name:       "'this' doc1 document frequency",
			frequency:  doc1Freq,
			term:       "this",
			tokens:     tokens1,
			tokenCount: tokenCount1,
			want:       0.2,
		},
		{
			name:       "'this' doc2 document frequency",
			frequency:  doc2Freq,
			term:       "this",
			tokens:     tokens2,
			tokenCount: tokenCount2,
			want:       0.14285714285714285,
		},
		{
			name:       "'example' doc1 document frequency",
			frequency:  doc1Freq,
			term:       "example",
			tokens:     tokens1,
			tokenCount: tokenCount1,
			want:       0,
		},
		{
			name:       "'example' doc2 document frequency",
			frequency:  doc2Freq,
			term:       "example",
			tokens:     tokens2,
			tokenCount: tokenCount2,
			want:       0.42857142857142855,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Document{Frequency: tt.frequency, UniqueTokens: tt.tokens, TotalTokenCount: tt.tokenCount}
			if got := d.TermFrequency(tt.term); got != tt.want {
				t.Errorf("TermFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTfIdf_InverseDocumentFrequency(t *testing.T) {
	tests := []struct {
		name      string
		documents map[string]Document
		term      string
		want      float64
	}{
		{
			name:      "'this' inverse document frequency",
			documents: docs,
			term:      "this",
			want:      0,
		},
		{
			name:      "'example' inverse document frequency",
			documents: docs,
			term:      "example",
			want:      0.3010299956639812,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := TfIdf{
				documents: tt.documents,
				stopWords: make(map[string]bool, 0),
			}
			if got := i.InverseDocumentFrequency(tt.term); got != tt.want {
				t.Errorf("InverseDocumentFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTfIdf_TermFrequencyInverseDocumentFrequency(t *testing.T) {
	type fields struct {
		documents map[string]Document
		stopWords map[string]bool
	}
	type args struct {
		term     string
		document string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "'this' term frequency inverse document frequency doc1",
			fields: fields{
				documents: docs,
				stopWords: map[string]bool{},
			},
			args: args{
				term:     "this",
				document: "doc1",
			},
			want: 0,
		},
		{
			name: "'this' term frequency inverse document frequency doc2",
			fields: fields{
				documents: docs,
				stopWords: map[string]bool{},
			},
			args: args{
				term:     "this",
				document: "doc2",
			},
			want: 0,
		},
		{
			name: "'example' term frequency inverse document frequency doc1",
			fields: fields{
				documents: docs,
				stopWords: map[string]bool{},
			},
			args: args{
				term:     "example",
				document: "doc1",
			},
			want: 0,
		},
		{
			name: "'example' term frequency inverse document frequency doc2",
			fields: fields{
				documents: docs,
				stopWords: map[string]bool{},
			},
			args: args{
				term:     "example",
				document: "doc2",
			},
			want: 0.12901285528456335,
		},
		{
			name: "non existent document",
			fields: fields{
				documents: docs,
				stopWords: map[string]bool{},
			},
			args: args{
				term:     "example",
				document: "asdf",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := TfIdf{
				documents: tt.fields.documents,
				stopWords: tt.fields.stopWords,
			}
			if got := i.TermFrequencyInverseDocumentFrequency(tt.args.term, tt.args.document); got != tt.want {
				t.Errorf("TermFrequencyInverseDocumentFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTfIdf_AddDocument(t *testing.T) {
	tests := []struct {
		name        string
		documents   []string
		wantNumDocs int
	}{
		{
			name:        "add one document",
			documents:   []string{doc1Content},
			wantNumDocs: 1,
		},
		{
			name:        "try add two identical documents",
			documents:   []string{doc1Content, doc1Content},
			wantNumDocs: 1,
		},
		{
			name:        "try add empty document",
			documents:   []string{""},
			wantNumDocs: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := TfIdf{documents: make(map[string]Document)}
			for _, document := range tt.documents {
				i.AddDocument(document)
			}

			if got := len(i.documents); got != tt.wantNumDocs {
				t.Errorf("len(documents) = %v, want %v", got, tt.wantNumDocs)
			}
		})
	}
}

func TestTfIdf_Compare(t *testing.T) {
	tests := []struct {
		name      string
		documents map[string]Document
		doc1      string
		doc2      string
		want      float64
		wantErr   error
	}{
		{
			name: "compare doc1 with doc1",
			documents: map[string]Document{
				"ebfd60f5f708658b4b5ff376d33d3393": doc1,
			},
			doc1:    doc1Content,
			doc2:    doc1Content,
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := TfIdf{documents: tt.documents}
			result, err := i.Compare(tt.doc1, tt.doc2, CosineComparator)
			if err != tt.wantErr {
				t.Errorf("got err %v, want %v", err, tt.wantErr)
			}

			if result != tt.want {
				t.Errorf("got comparison value %v, want %v", result, tt.want)
			}
		})
	}
}

func Test_md5Hash(t *testing.T) {
	if got := md5Hash("doc1"); got != doc1Hash {
		t.Errorf("md5Hash() = %v, want %v", got, doc1Hash)
	}

	if got := md5Hash("doc2"); got != doc2Hash {
		t.Errorf("md5Hash() = %v, want %v", got, doc2Hash)
	}
}

func TestNew(t *testing.T) {
	i := New([]string{
		doc1Content, doc2Content,
	})
	if len(i.documents) != 2 {
		t.Error("invalid # of docs")
	}
}
