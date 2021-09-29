# go-tf-idf

A small Go implementation of [tf-idf](https://en.wikipedia.org/wiki/Tf%E2%80%93idf) (term frequency-inverse document frequency) with support for comparing documents using [cosine similarities](https://en.wikipedia.org/wiki/Cosine_similarity).

## Usage
Install with `go get github.com/dkgv/go-tf-idf`.

```go
package main

import (
	"fmt"
	go_tf_idf "github.com/dkgv/go-tf-idf"
)

func main() {
	// Initializing a tf-idf container 
	doc1 := "this is a document"
	doc2 := "and this is another document"
	tfidf := go_tf_idf.New([]string{
		doc1,
		doc2,
	})

	// Calculating tf-idf for a term
	term := "document"
	res1 := tfidf.TermFrequencyInverseDocumentFrequency(term, doc1)
	fmt.Printf("res1 %f", res1)

	// Comparing two documents via cosine similarity
	comparator := go_tf_idf.CosineComparator
	similarity, err := tfidf.Compare(doc1, doc2, comparator)
	if err != nil {
		panic("comparison failed")
	}
	fmt.Printf("similarity %f", similarity)
}
```
