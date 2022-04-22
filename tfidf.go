package go_tf_idf

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math"
)

type Option func(idf *TfIdf)

func WithStopWords(stopWords []string) Option {
	return func(tfIdf *TfIdf) {
		tfIdf.StopWords.AddWords(stopWords)
	}
}

func WithDefaultStopWords() Option {
	return func(tfIdf *TfIdf) {
		for s, _ := range DefaultList {
			tfIdf.StopWords.AddWord(s)
		}
	}
}

func WithDocuments(documents []string) Option {
	return func(tfIdf *TfIdf) {
		for _, document := range documents {
			tfIdf.AddDocument(document)
		}
	}
}

func WithComparator(comparator Comparator) Option {
	return func(tfIdf *TfIdf) {
		tfIdf.comparator = comparator
	}
}

type TfIdf struct {
	Documents              map[string]Document
	StopWords              *StopWords
	comparator             Comparator
	termToIndex            map[string]int
	documentsWithTermCount map[string]int
}

func DefaultOptions() *TfIdf {
	return &TfIdf{
		Documents:              make(map[string]Document, 0),
		StopWords:              NewEmptyStopWords(),
		comparator:             CosineComparator,
		termToIndex:            make(map[string]int, 0),
		documentsWithTermCount: make(map[string]int, 0),
	}
}

func New(opts ...Option) *TfIdf {
	tfIdf := DefaultOptions()
	for _, opt := range opts {
		opt(tfIdf)
	}

	return tfIdf
}

type Document struct {
	AllTokens    []string
	TermCount    map[string]int
	UniqueTokens []string
}

func (d Document) TermFrequency(term string) float64 {
	if _, ok := d.TermCount[term]; !ok {
		return 0
	}
	return float64(d.TermCount[term]) / float64(len(d.AllTokens))
}

func (d Document) GetVectors(other Document) ([]float64, []float64) {
	visited := make(map[string]bool, 0)
	terms := make([]string, 0)
	for _, token := range d.UniqueTokens {
		if _, ok := visited[token]; !ok {
			terms = append(terms, token)
			visited[token] = true
		}
	}
	for _, token := range other.UniqueTokens {
		if _, ok := visited[token]; !ok {
			terms = append(terms, token)
			visited[token] = true
		}
	}

	vector1 := make([]float64, len(visited))
	vector2 := make([]float64, len(visited))
	index := 0
	for _, term := range terms {
		vector1[index] = d.TermFrequency(term)
		vector2[index] = other.TermFrequency(term)
		index += 1
	}

	return vector1, vector2
}

type Comparator func(vector1, vector2 []float64) float64

func (i TfIdf) Compare(document1, document2 string) (float64, error) {
	doc1 := i.GetDocument(document1)
	doc2 := i.GetDocument(document2)
	if doc1 == nil || doc2 == nil {
		return 0, errors.New("cannot compare with nil document")
	}

	vector1, vector2 := doc1.GetVectors(*doc2)
	return i.comparator(vector1, vector2), nil
}

func (i TfIdf) GetDocument(document string) *Document {
	hash := md5Hash(document)
	if doc, ok := i.Documents[hash]; ok {
		return &doc
	}

	return nil
}

func (i TfIdf) InverseDocumentFrequency(term string) float64 {
	termCount := i.documentsWithTermCount[term]
	documentCount := len(i.Documents)
	return math.Log10(float64(documentCount) / float64(termCount))
}

func (i TfIdf) TermFrequencyInverseDocumentFrequencyForTerm(term string, document string) float64 {
	doc := i.GetDocument(document)
	if doc == nil {
		return 0
	}
	return doc.TermFrequency(term) * i.InverseDocumentFrequency(term)
}

func (i TfIdf) TermFrequencyInverseDocumentFrequencyForDocument(document string) []float64 {
	vec := make([]float64, len(i.termToIndex))
	doc := i.GetDocument(document)
	if doc == nil {
		return vec
	}

	for _, term := range doc.AllTokens {
		idf := i.InverseDocumentFrequency(term)
		tfidf := doc.TermFrequency(term) * idf
		vec[i.termToIndex[term]] = tfidf
	}

	return vec
}

func (i TfIdf) AddDocument(document string) {
	hash := md5Hash(document)
	if _, ok := i.Documents[hash]; ok {
		return
	}

	allTokens := Tokenize(document)
	if len(allTokens) == 0 {
		return
	}

	termCount := make(map[string]int, 0)
	uniqueTokens := make([]string, 0)
	for _, token := range allTokens {
		if i.StopWords.Matches(token) {
			continue
		}

		termCount[token]++

		if termCount[token] == 1 {
			uniqueTokens = append(uniqueTokens, token)
			i.documentsWithTermCount[token]++
		}

		if _, ok := i.termToIndex[token]; !ok {
			i.termToIndex[token] = len(i.termToIndex)
		}
	}

	i.Documents[hash] = Document{
		AllTokens:    allTokens,
		UniqueTokens: uniqueTokens,
		TermCount:    termCount,
	}
}

func md5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
