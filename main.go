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
			fmt.Printf("%s exists in %s document(s)\n", colorText(token), colorInt(len(postings)))

			for _, posting := range postings {
				doc, err := searchEngine.GetDocument(posting.DocID)
				if err != nil {
					continue
				}
				fmt.Printf("  docID=%s freq=%s path=%s\n", colorInt(posting.DocID), colorInt(posting.Frequency), colorText(doc.Path))
			}
		} else {
			fmt.Println("Not found")
		}
	}

}

func colorText(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

func colorInt(i int) string {
	return fmt.Sprintf("\033[32m%d\033[0m", i)
}
