package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	fmt.Println("Indexing...")
	err := searchEngine.Index(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexd %d documents\n", searchEngine.GetIndexedDocumentsCount())

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("search query: ")
		input.Scan()

		// read and parse query
		queryStr := input.Text()
		query := searchEngine.ParseQuery(queryStr)

		for _, clause := range query.Clauses {
			// get common postings for all [clause.Terms]
			termToPostings := searchEngine.GetMergedPostings(clause.Terms)
			joinedTerms := strings.Join(clause.Terms, " ")

			// no document that has all [clause.Terms]
			if len(termToPostings) == 0 {
				fmt.Printf("[%s] Not found\n", colorText(joinedTerms))
				continue
			}

			// Only get first postings. All postings have same docID
			// Frequency is not the same, but we are not printing it for now, so it's fine
			postings := termToPostings[clause.Terms[0]]

			fmt.Printf("[%s] exists in %s document(s)\n", colorText(joinedTerms), colorInt(len(postings)))
			for _, posting := range postings {
				doc, err := searchEngine.GetDocument(posting.DocID)
				if err != nil {
					continue
				}
				fmt.Printf("  docID=%s path=%s\n", colorInt(posting.DocID), colorText(doc.Path))
			}
		}
	}

}

func colorText(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

func colorInt(i int) string {
	return fmt.Sprintf("\033[32m%d\033[0m", i)
}
