package ai

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var englishSignals = map[string]struct{}{
	"the": {}, "and": {}, "of": {}, "to": {}, "is": {}, "in": {}, "you": {}, "it": {}, "that": {},
}

var frenchSignals = map[string]struct{}{
	"le": {}, "la": {}, "les": {}, "de": {}, "et": {}, "un": {}, "une": {}, "est": {}, "dans": {}, "vous": {},
}


var languageWordRegexp = regexp.MustCompile(`[\p{L}0-9']+`)

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

