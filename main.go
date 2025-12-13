package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mabd-dev/search-engine/internal/engine"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "Folder/file path to index")
	flag.Parse()

	if len(strings.TrimSpace(path)) == 0 {
		panic("Path is invalid")
	}

	searchEngine := engine.NewEngine()

	err := searchEngine.Index(path)
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
	fmt.Printf("\033[32m%s\033[0m exists in %d document(s)\n", token, len(postings))
	for _, posting := range postings {
		fmt.Printf("  docID=%d freq=%d\n", posting.DocID, posting.Frequency)
	}
	fmt.Println(strings.Repeat(" ̶", 30))
}
