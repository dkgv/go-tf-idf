package go_tf_idf

import (
	"reflect"
	"strings"
	"testing"
)

var doc1Hash = "83e4b1789306d3d1c99140df3827d600"
var doc1Content = "this is a a sample"
var doc1Freq = map[string]int{
	"this":   1,
	"is":     1,
	"a":      2,
	"sample": 1,
}
var allTokens1 = strings.Split(doc1Content, " ")
var uniqueTokens1 = []string{"this", "is", "a", "sample"}
var doc1 = Document{AllTokens: allTokens1, TermCount: doc1Freq, UniqueTokens: uniqueTokens1}

var doc2Hash = "271559ec25268bb9bb2ad7fd8b4cf71a"
var doc2Content = "this is another another example example example"
var doc2Freq = map[string]int{
	"this":    1,
	"is":      1,
	"another": 2,
	"example": 3,
}
var allTokens2 = strings.Split(doc2Content, " ")
var uniqueTokens2 = []string{"this", "is", "another", "example"}
var doc2 = Document{AllTokens: allTokens2, TermCount: doc2Freq, UniqueTokens: uniqueTokens2}

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

func TestDocument_TermFrequency(t *testing.T) {
	tests := []struct {
		name         string
		frequency    map[string]int
		term         string
		allTokens    []string
		uniqueTokens []string
		want         float64
	}{
		{
			name:         "'this' doc1 document frequency",
			frequency:    doc1Freq,
			term:         "this",
			allTokens:    allTokens1,
			uniqueTokens: uniqueTokens1,
			want:         0.2,
		},
		{
			name:         "'this' doc2 document frequency",
			frequency:    doc2Freq,
			term:         "this",
			allTokens:    allTokens2,
			uniqueTokens: uniqueTokens2,
			want:         0.14285714285714285,
		},
		{
			name:         "'example' doc1 document frequency",
			frequency:    doc1Freq,
			term:         "example",
			allTokens:    allTokens1,
			uniqueTokens: uniqueTokens1,
			want:         0,
		},
		{
			name:         "'example' doc2 document frequency",
			frequency:    doc2Freq,
			term:         "example",
			allTokens:    allTokens2,
			uniqueTokens: uniqueTokens2,
			want:         0.42857142857142855,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Document{AllTokens: tt.allTokens, TermCount: tt.frequency, UniqueTokens: tt.uniqueTokens}
			if got := d.TermFrequency(tt.term); got != tt.want {
				t.Errorf("TermFrequency() = %v, want %v", got, tt.want)
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
			getDocument: doc1Content,
			want:        &doc1,
		},
		{
			name:        "doc2",
			getDocument: doc2Content,
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
			i := New(
				WithDocuments([]string{doc1Content, doc2Content}),
			)
			got := i.GetDocument(tt.getDocument)
			if got != tt.want && !reflect.DeepEqual(got.UniqueTokens, tt.want.UniqueTokens) {
				t.Errorf("GetDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTfIdf_InverseDocumentFrequency(t *testing.T) {
	tests := []struct {
		name      string
		documents []string
		term      string
		want      float64
	}{
		{
			name:      "'this' inverse document frequency",
			documents: []string{doc1Content, doc2Content},
			term:      "this",
			want:      0,
		},
		{
			name:      "'example' inverse document frequency",
			documents: []string{doc1Content, doc2Content},
			term:      "example",
			want:      0.3010299956639812,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := New(
				WithDocuments(tt.documents),
			)
			if got := i.InverseDocumentFrequency(tt.term); got != tt.want {
				t.Errorf("InverseDocumentFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTfIdf_TermFrequencyInverseDocumentFrequencyForTerm(t *testing.T) {
	type fields struct {
		documents []string
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
				documents: []string{doc1Content, doc2Content},
			},
			args: args{
				term:     "this",
				document: doc1Content,
			},
			want: 0,
		},
		{
			name: "'this' term frequency inverse document frequency doc2",
			fields: fields{
				documents: []string{doc1Content, doc2Content},
			},
			args: args{
				term:     "this",
				document: doc2Content,
			},
			want: 0,
		},
		{
			name: "'example' term frequency inverse document frequency doc1",
			fields: fields{
				documents: []string{doc1Content, doc2Content},
			},
			args: args{
				term:     "example",
				document: doc1Content,
			},
			want: 0,
		},
		{
			name: "'example' term frequency inverse document frequency doc2",
			fields: fields{
				documents: []string{doc1Content, doc2Content},
			},
			args: args{
				term:     "example",
				document: doc2Content,
			},
			want: 0.12901285528456335,
		},
		{
			name: "non existent document",
			fields: fields{
				documents: []string{doc1Content, doc2Content},
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
			i := New(
				WithDocuments(tt.fields.documents),
			)
			if got := i.TermFrequencyInverseDocumentFrequencyForTerm(tt.args.term, tt.args.document); got != tt.want {
				t.Errorf("TermFrequencyInverseDocumentFrequencyForTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTfIdf_AddDocument(t *testing.T) {
	tests := []struct {
		name        string
		documents   []string
		stopWords   []string
		wantNumDocs int
	}{
		{
			name:        "add one document",
			documents:   []string{doc1Content},
			wantNumDocs: 1,
		},
		{
			name:        "add one document",
			documents:   []string{doc1Content},
			stopWords:   []string{"this"},
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
			i := New(
				WithStopWords(tt.stopWords),
			)
			for _, document := range tt.documents {
				i.AddDocument(document)
			}

			if got := len(i.Documents); got != tt.wantNumDocs {
				t.Errorf("len(documents) = %v, want %v", got, tt.wantNumDocs)
			}
		})
	}
}

func TestTfIdf_Compare(t *testing.T) {
	tests := []struct {
		name      string
		documents []string
		doc1      string
		doc2      string
		want      float64
		wantErr   bool
	}{
		{
			name:      "cosine compare doc1 with doc1",
			documents: []string{doc1Content},
			doc1:      doc1Content,
			doc2:      doc1Content,
			want:      1,
			wantErr:   false,
		},
		{
			name:      "cosine compare doc1 with nil",
			documents: []string{doc1Content},
			doc1:      doc1Content,
			doc2:      "",
			want:      0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := New(
				WithDocuments(tt.documents),
				WithComparator(CosineComparator),
			)
			result, err := i.Compare(tt.doc1, tt.doc2)
			if !tt.wantErr && err != nil {
				t.Errorf("got err %v, want none", err)
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

func TestNewWithDocuments(t *testing.T) {
	i := New(
		WithDocuments([]string{doc1Content, doc2Content}),
	)
	if len(i.Documents) != 2 {
		t.Error("invalid # of docs")
	}
}

func TestTfIdf_TermFrequencyInverseDocumentFrequencyForDocument(t *testing.T) {
	type fields struct {
		Options []Option
	}
	type args struct {
		document string
	}
	documents := []string{
		doc1Content,
		doc2Content,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []float64
	}{
		{
			name: "doc1",
			fields: fields{
				Options: []Option{
					WithDocuments(documents),
				},
			},
			args: args{
				document: documents[0],
			},
			want: []float64{0, 0, 0.12041199826559248, 0.06020599913279624, 0, 0},
		},
		{
			name: "doc1 with stopwords",
			fields: fields{
				Options: []Option{
					WithDocuments(documents),
					WithDefaultStopWords(),
				},
			},
			args: args{
				document: documents[0],
			},
			want: []float64{0, 0, 0.12041199826559248, 0.06020599913279624, 0, 0},
		},
		{
			name: "doc2",
			fields: fields{
				Options: []Option{
					WithDocuments(documents),
				},
			},
			args: args{
				document: documents[1],
			},
			want: []float64{0, 0, 0, 0, 0.08600857018970891, 0.12901285528456335},
		},
		{
			name: "nil document",
			fields: fields{
				Options: []Option{
					WithDocuments(documents),
				},
			},
			args: args{
				document: "asdfasdf",
			},
			want: []float64{0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := New(tt.fields.Options...)
			if got := i.TermFrequencyInverseDocumentFrequencyForDocument(tt.args.document); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TermFrequencyInverseDocumentFrequencyForDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
