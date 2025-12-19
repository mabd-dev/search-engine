package engine

func intersectPostingsSorted(termToPosting map[string][]Posting) map[string][]Posting {
	if len(termToPosting) == 0 {
		return nil
	}

	// Collect terms and find the shortest posting list
	terms := make([]string, 0, len(termToPosting))
	shortestIdx := 0
	shortestLen := -1
	for term, postings := range termToPosting {
		if shortestLen == -1 || len(postings) < shortestLen {
			shortestLen = len(postings)
			shortestIdx = len(terms)
		}
		terms = append(terms, term)
	}
	// Move shortest to front
	terms[0], terms[shortestIdx] = terms[shortestIdx], terms[0]

	// Start with DocIDs from shortest list
	commonDocIDs := make([]int, len(termToPosting[terms[0]]))
	for i, p := range termToPosting[terms[0]] {
		commonDocIDs[i] = p.DocID
	}

	// Intersect with each other sorted list
	for i := 1; i < len(terms) && len(commonDocIDs) > 0; i++ {
		commonDocIDs = intersectSorted(commonDocIDs, termToPosting[terms[i]])
	}

	if len(commonDocIDs) == 0 {
		return map[string][]Posting{}
	}

	// Build result by filtering each posting list
	result := make(map[string][]Posting, len(termToPosting))
	for term, postings := range termToPosting {
		result[term] = filterPostings(postings, commonDocIDs)
	}

	return result
}

// intersectSorted finds DocIDs present in both the sorted slice and sorted posting list
func intersectSorted(docIDs []int, postings []Posting) []int {
	result := make([]int, 0, min(len(docIDs), len(postings)))
	i, j := 0, 0

	for i < len(docIDs) && j < len(postings) {
		if docIDs[i] == postings[j].DocID {
			result = append(result, docIDs[i])
			i++
			j++
		} else if docIDs[i] < postings[j].DocID {
			i++
		} else {
			j++
		}
	}
	return result
}

// filterPostings returns postings whose DocID exists in the sorted commonDocIDs slice
func filterPostings(postings []Posting, commonDocIDs []int) []Posting {
	result := make([]Posting, 0, len(commonDocIDs))
	i, j := 0, 0

	for i < len(postings) && j < len(commonDocIDs) {
		if postings[i].DocID == commonDocIDs[j] {
			result = append(result, postings[i])
			i++
			j++
		} else if postings[i].DocID < commonDocIDs[j] {
			i++
		} else {
			j++
		}
	}
	return result
}
