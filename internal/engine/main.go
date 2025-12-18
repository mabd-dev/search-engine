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

		ext := filepath.Ext(p)
		if !isSupportedFileExtensions(ext) {
			return nil
		}

		fInfo, err := d.Info()
		if err != nil {
			return err
		}

		docID := len(e.index.docs) + 1
		doc := Document{
			ID:            docID,
			Name:          fInfo.Name(),
			Path:          p,
			FileExtension: ext,
		}
		if err := e.indexDocument(doc); err != nil {
			return err
		}
		e.index.docs[doc.ID] = doc

		return nil
	})
}

func isSupportedFileExtensions(extension string) bool {
	return extension == ".md"
}

func (e SearchEngine) indexDocument(doc Document) error {
	bytes, err := os.ReadFile(doc.Path)
	if err != nil {
		return err
	}
	content := decodeFileContent(bytes, doc.FileExtension)
	tokens := e.Tokenizer.Tokenize(content)
	tokensFreq := getTokensFrequencies(tokens)

	for token, freq := range tokensFreq {
		posting := Posting{
			DocID:     doc.ID,
			Frequency: freq,
		}
		e.index.postings[token] = append(e.index.postings[token], posting)
	}

	return nil
}

func cleanToken(word string) string {
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	})
}

func getTokensFrequencies(tokens []string) map[string]int {
	frequencies := map[string]int{}

	for _, token := range tokens {
		prevFreq := frequencies[token]
		frequencies[token] = prevFreq + 1
	}
	return frequencies
}

func (e SearchEngine) GetIndexedDocumentsCount() int {
	return len(e.index.docs)
}

func (e SearchEngine) GetDocument(id int) (Document, error) {
	doc, found := e.index.docs[id]
	if !found {
		return Document{}, fmt.Errorf("Document not found")
	}
	return doc, nil
}

func (e SearchEngine) GetPostings(term string) []Posting {
	postings := e.index.postings[strings.ToLower(term)]
	return postings
}

func (e SearchEngine) PrintAllIndexedDocuments() {
	fmt.Println("Indexed Documents:")
	for _, doc := range e.index.docs {
		fmt.Printf("%d, path=%s\n", doc.ID, doc.Path)
	}
	fmt.Println("--------------")
}
