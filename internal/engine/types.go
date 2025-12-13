package engine

type Document struct {
	ID            int
	Name          string
	Path          string
	FileExtension string
}

type Posting struct {
	DocID int
}

type Index struct {
	postings map[string][]Posting
	docs     map[int]Document
}

type SearchEngine struct {
	index Index
}
