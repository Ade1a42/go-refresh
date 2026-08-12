package pipeline

import (
	"goRefresh/internal/cases"
	"goRefresh/internal/grammar"
	"goRefresh/internal/numbers"
	"goRefresh/internal/punct"
	"goRefresh/internal/quotes"
	"regexp"
	"strconv"
	"strings"
)

// making tokens based on (), punctuation, quote and words
func Tokenizer(content string) []string {
	re := regexp.MustCompile(`\([^)]+\)|[.,!?:;]{2,}|[.,!?:;]|'|[\w-]+`)
	return re.FindAllString(content, -1)
}

func IsTag(word string) bool {
	if !strings.HasPrefix(word, "(") || !strings.HasSuffix(word, ")") {
		return false
	}

	if len(word) < 4 {
		return false
	}

	tag := strings.ReplaceAll(word[1:len(word)-1], " ", "")
	tags := strings.Split(tag, ",")
	tag_word := tags[0]

	// if tag with number
	if len(tags) == 2 {
		if tag_word == "up" || tag_word == "low" || tag_word == "cap" {
			num, err := strconv.Atoi(tags[1])
			if err != nil || num <= 0 {
				return false
			}
			return true
		} else {
			return false
		}

		// if tag without number
	} else if len(tags) == 1 {
		if tag_word == "hex" || tag_word == "bin" || tag_word == "up" || tag_word == "low" || tag_word == "cap" {
			return true
		}
	}

	return false
}

func ApplyTag(word string, tokens []string, i int) []string {
	tag := strings.ReplaceAll(word[1:len(word)-1], " ", "")
	tags := strings.Split(tag, ",")
	tag_word := tags[0]

	if len(tags) == 2 {
		// handling number
		num, err := strconv.Atoi(tags[1])
		if err != nil {
			return tokens
		}
		if num > len(tokens[:i]) {
			num = len(tokens[:i])
		}

		switch tag_word {
		case "up":
			tokens = cases.UpN(tokens, num, i)
			return tokens
		case "low":
			tokens = cases.LowN(tokens, num, i)
			return tokens
		case "cap":
			tokens = cases.CapN(tokens, num, i)
			return tokens
		}
	} else {
		switch tag_word {
		case "hex":
			tokens = numbers.Hex(tokens, i)
			return tokens
		case "bin":
			tokens = numbers.Bin(tokens, i)
			return tokens
		case "up":
			tokens = cases.UpN(tokens, 1, i)
			return tokens
		case "low":
			tokens = cases.LowN(tokens, 1, i)
			return tokens
		case "cap":
			tokens = cases.CapN(tokens, 1, i)
			return tokens
		}
	}

	return tokens
}

func Process(tokens []string) []string {
	// loop for tags
	for i := 0; i < len(tokens); i++ {
		word := tokens[i]
		if IsTag(word) {
			if i > 0 {
				tokens = ApplyTag(word, tokens, i)
			}

			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}

	// loop for quote
	for i := 0; i < len(tokens); i++ {
		if quotes.IsQuote(tokens[i]) {
			tokens = quotes.FixQuotes(tokens, i)
			i--
		}
	}

	// loop for punctuation
	for i := 0; i < len(tokens); i++ {
		if punct.IsPunctuations(tokens[i]) && i > 0 {
			tokens = punct.FixSpacing(tokens, i)
			i--
		}
	}

	// loop for article
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "A" || tokens[i] == "a" {
			tokens = grammar.FixArticle(tokens, i)
		}
	}

	return tokens
}
