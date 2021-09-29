package go_tf_idf

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math"
)

type Document struct {
	Frequency       map[string]float64
	UniqueTokens    []string
	TotalTokenCount int
}

func (d Document) TermFrequency(term string) float64 {
	if _, ok := d.Frequency[term]; !ok {
		return 0
	}
	return d.Frequency[term] / (float64)(d.TotalTokenCount)
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

type TfIdf struct {
	documents map[string]Document
	stopWords map[string]bool
}

func New(documents []string) *TfIdf {
	tfidf := TfIdf{
		documents: make(map[string]Document),
		stopWords: make(map[string]bool),
	}
	for _, document := range documents {
		tfidf.AddDocument(document)
	}
	return &tfidf
}

type Comparator func(vector1, vector2 []float64) float64

func (i TfIdf) Compare(document1, document2 string, comparator Comparator) (float64, error) {
	doc1 := i.GetDocument(document1)
	doc2 := i.GetDocument(document2)
	if doc1 == nil || doc2 == nil {
		return 0, errors.New("cannot compare with nil document")
	}

	vector1, vector2 := doc1.GetVectors(*doc2)
	return comparator(vector1, vector2), nil
}

func (i TfIdf) GetDocument(document string) *Document {
	hash := md5Hash(document)
	if doc, ok := i.documents[hash]; ok {
		return &doc
	}
	return nil
}

func (i TfIdf) InverseDocumentFrequency(term string) float64 {
	tf := float64(0)
	for _, document := range i.documents {
		if _, ok := document.Frequency[term]; ok {
			tf++
		}
	}

	numerator := (float64)(len(i.documents))
	return math.Log10(numerator / tf)
}

func (i TfIdf) TermFrequencyInverseDocumentFrequency(term string, document string) float64 {
	doc := i.GetDocument(document)
	if doc == nil {
		return 0
	}
	return doc.TermFrequency(term) * i.InverseDocumentFrequency(term)
}

func (i TfIdf) AddDocument(document string) {
	hash := md5Hash(document)
	if _, ok := i.documents[hash]; ok {
		return
	}

	tokens := Tokenize(document)
	if len(tokens) == 0 {
		return
	}

	frequency := make(map[string]float64, len(tokens))
	uniqueTokens := make([]string, 0)
	for _, token := range tokens {
		if _, ok := i.stopWords[token]; ok {
			continue
		}

		frequency[token]++

		if frequency[token] == 1 {
			uniqueTokens = append(uniqueTokens, token)
		}
	}

	i.documents[hash] = Document{
		Frequency:       frequency,
		UniqueTokens:    uniqueTokens,
		TotalTokenCount: len(tokens),
	}
}

func (i TfIdf) AddStopWord(word string) {
	i.stopWords[word] = true
}

func md5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
