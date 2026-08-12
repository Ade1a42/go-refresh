package ai

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

var englishSignals = map[string]struct{}{
	"the": {}, "and": {}, "of": {}, "to": {}, "is": {}, "in": {}, "you": {}, "it": {}, "that": {},
}

var frenchSignals = map[string]struct{}{
	"le": {}, "la": {}, "les": {}, "de": {}, "et": {}, "un": {}, "une": {}, "est": {}, "dans": {}, "vous": {},
}

var keywordStopwords = map[string]struct{}{
	"the": {}, "and": {}, "of": {}, "to": {}, "is": {}, "in": {}, "it": {}, "that": {},
	"a": {}, "an": {}, "for": {}, "on": {}, "with": {}, "as": {}, "by": {}, "at": {}, "from": {},
	"this": {}, "these": {}, "those": {}, "de": {}, "la": {}, "le": {}, "les": {}, "et": {}, "un": {}, "une": {}, "des": {},
}

var languageWordRegexp = regexp.MustCompile(`[\p{L}0-9']+`)
var keywordRegexp = regexp.MustCompile(`[a-z0-9']+`)

func DetectLanguage(text string) string {
	engCount := 0
	frCount := 0

	for _, token := range languageWordRegexp.FindAllString(text, -1) {
		lower := strings.ToLower(token)
		if _, ok := englishSignals[lower]; ok {
			engCount++
		}
		if _, ok := frenchSignals[lower]; ok {
			frCount++
		}
	}

	lowerText := strings.ToLower(text)
	for _, r := range lowerText {
		switch r {
		case 'é', 'è', 'à', 'ç', 'ô', 'î', 'û':
			frCount++
		}
	}

	total := engCount + frCount
	if total == 0 {
		return "Language: Unknown"
	}

	if engCount >= frCount {
		score := int(math.Round(float64(engCount) * 100 / float64(total)))
		return fmt.Sprintf("Language: English (%d%%)", score)
	}

	score := int(math.Round(float64(frCount) * 100 / float64(total)))
	return fmt.Sprintf("Language: French (%d%%)", score)
}

func ExtractKeywords(text string) string {
	lowerText := strings.ToLower(text)
	frequency := map[string]int{}

	for _, token := range keywordRegexp.FindAllString(lowerText, -1) {
		word := strings.Trim(token, "'")
		if len(word) == 0 {
			continue
		}
		if _, ok := keywordStopwords[word]; ok {
			continue
		}
		frequency[word]++
	}

	if len(frequency) == 0 {
		return "Keywords: (none)"
	}

	type entry struct {
		word  string
		count int
	}

	entries := make([]entry, 0, len(frequency))
	for word, count := range frequency {
		entries = append(entries, entry{word: word, count: count})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].word < entries[j].word
	})

	limit := 5
	if len(entries) < limit {
		limit = len(entries)
	}

	keywords := make([]string, limit)
	for i := 0; i < limit; i++ {
		keywords[i] = entries[i].word
	}

	return "Keywords: " + strings.Join(keywords, ", ")
}
