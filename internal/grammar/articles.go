package grammar

import (
	"strings"
)

func NeedArticleChange(word string) bool {
	if len(word) == 0 {
		return false
	}

	ch := strings.ToLower(string(word[0]))
	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" || ch == "h" {
		return true
	}

	return false
}

func FixArticle(tokens []string, i int) []string {
	if i+1 >= len(tokens) {
		return tokens
	}

	if NeedArticleChange(tokens[i+1]) {
		tokens[i] = tokens[i] + "n"
	}
	return tokens
}
