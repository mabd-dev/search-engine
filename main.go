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

		token := input.Text()
		postings := searchEngine.GetPostings(token)

		if len(postings) > 0 {
			fmt.Printf("\033[32m%s\033[0m exists in %d document(s)\n", token, len(postings))

			for _, posting := range postings {
				fmt.Printf("  docID=%d freq=%d\n", posting.DocID, posting.Frequency)
			}
		} else {
			fmt.Println("Not found")
		}
	}

}
