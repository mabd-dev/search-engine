package engine

import (
	"strings"
	"unicode"
)

type Tokenizer struct {
}

// Tokenize takes document content and split into tokens
func (t Tokenizer) Tokenize(s string) []string {
	tokens := strings.Fields(s)
	return linguisticPreprocessing(tokens)
}

// TokeizeBy tokenizes string [s] based on [delimeter]
// NOTE: if delimeter is white spacee consider using [Tokenize] token
func (t Tokenizer) TokeizeBy(s string, delimeter string) []string {
	tokens := strings.Split(s, delimeter)
	return linguisticPreprocessing(tokens)
}

// linguisticPreprocessing removes token leading+trailing symbols and punctuations
func linguisticPreprocessing(tokens []string) []string {
	newTokens := []string{}

	for _, token := range tokens {
		cleanedToken := cleanToken(strings.ToLower(strings.TrimSpace(token)))
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
