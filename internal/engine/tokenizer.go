package engine

import "strings"

type Tokenizer struct {
}

// Tokenize takes document content and split into tokens
func (t Tokenizer) Tokenize(s string) []string {
	tokens := strings.Fields(s)
	terms := linguisticPreprocessing(tokens)
	return terms
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
