package main

import (
	"fmt"

	"github.com/mabd-dev/search-engine/internal/engine"
)

func main() {
	searchEngine := engine.NewEngine()

	err := searchEngine.Index("/Users/mabd/Documents/mind-map/Books/")
	if err != nil {
		panic(err)
	}

	// searchEngine.PrintAllIndexedDocuments()

	// err := searchEngine.Index("/Users/mabd/Documents/mind-map/Books/Introduction to Information Retrievals.md")
	// if err != nil {
	// 	panic(err)
	// }

	// err = searchEngine.Index("/Users/mabd/Documents/mind-map/Books/testing.md")
	// if err != nil {
	// 	panic(err)
	// }

	tokens := []string{
		"document",
		"shit",
		"it",
	}

	for _, token := range tokens {
		postings := searchEngine.GetPostings(token)
		printPostings(token, postings)
	}

}

func printPostings(token string, postings []engine.Posting) {
	fmt.Printf("[%s] exists in %d document(s)", token, len(postings))
	// for _, posting := range postings {
	// 	fmt.Printf("docID=%d ", posting.DocID)
	// }
	fmt.Println("")
}
