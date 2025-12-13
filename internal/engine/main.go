package engine

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func NewEngine() SearchEngine {
	index := Index{
		postings: map[string][]Posting{},
		docs:     map[int]Document{},
	}
	return SearchEngine{
		index: index,
	}
}

// Index walk through path if dir, and index all files in it. Or if path is file index that file only
func (e *SearchEngine) Index(path string) error {
	return filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fInfo, err := d.Info()
		if err != nil {
			return err
		}

		docID := len(e.index.docs) + 1
		doc := Document{
			ID:   docID,
			Name: fInfo.Name(),
			Path: p,
		}
		if err := e.indexDocument(doc); err != nil {
			return err
		}
		e.index.docs[doc.ID] = doc

		return nil
	})
}

func (e SearchEngine) indexDocument(doc Document) error {
	bytes, err := os.ReadFile(doc.Path)
	if err != nil {
		return err
	}
	content := string(bytes)
	tokens := tokenizeDocumentContent(content)
	tokens = linguisticPreprocessing(tokens)
	tokens = removeDuplicates(tokens)

	for _, token := range tokens {
		postings := e.index.postings[token]
		postings = append(postings, Posting{DocID: doc.ID})
		e.index.postings[token] = postings
	}

	return nil
}

// tokenizeDocumentContent takes document content and split into tokens
func tokenizeDocumentContent(s string) []string {
	return strings.Fields(s)
}

// linguisticPreprocessing removes token leading+trailing symbols and punctuations
func linguisticPreprocessing(tokens []string) []string {
	newTokens := []string{}

	for _, token := range tokens {
		cleanedToken := cleanToken(strings.ToLower(token))
		if len(cleanedToken) > 0 {
			newTokens = append(newTokens, cleanedToken)
		}
	}

	return newTokens
}

func cleanToken(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	})
}

func removeDuplicates(tokens []string) []string {
	set := map[string]struct{}{}

	for _, token := range tokens {
		if _, exist := set[token]; !exist {
			set[token] = struct{}{}
		}
	}

	uniqueTokens := []string{}
	for token := range set {
		uniqueTokens = append(uniqueTokens, token)
	}
	return uniqueTokens
}

func (e SearchEngine) GetIndexedDocumentsCount() int {
	return len(e.index.docs)
}

func (e SearchEngine) GetPostings(token string) []Posting {
	postings := e.index.postings[strings.ToLower(token)]
	return postings
}

func (e SearchEngine) PrintAllIndexedDocuments() {
	fmt.Println("Indexed Documents:")
	for _, doc := range e.index.docs {
		fmt.Println(doc.Path)
	}
	fmt.Println("--------------")
}
