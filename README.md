# go-tf-idf
[![Coverage Status](https://coveralls.io/repos/github/dkgv/go-tf-idf/badge.svg?branch=master)](https://coveralls.io/github/dkgv/go-tf-idf?branch=master)

A small Go implementation of [tf-idf](https://en.wikipedia.org/wiki/Tf%E2%80%93idf) (term frequency-inverse document frequency) with support for comparing documents using [cosine similarities](https://en.wikipedia.org/wiki/Cosine_similarity).

## Usage
Install with `go get -u github.com/dkgv/go-tf-idf`.

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
    tfidf := go_tf_idf.New(
        go_tf_idf.WithDocuments([]string{doc1, doc2}),
        go_tf_idf.WithDefaultStopWords(),
    )

    // Calculating tf-idf for a term
    term := "document"
    res1 := tfidf.TermFrequencyInverseDocumentFrequencyForTerm(term, doc1)
    fmt.Printf("res1 %f", res1)

    // Comparing two documents via cosine similarity
    comparator := go_tf_idf.CosineComparator
    similarity, err := tfidf.Compare(doc1, doc2, comparator)
    if err != nil {
        // ...
    }
	
    fmt.Printf("similarity %f", similarity)
}
```
