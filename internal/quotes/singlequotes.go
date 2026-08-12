package quotes

func IsQuote(word string) bool {
	if len(word) == 1 {
		if word == "'" {
			return true
		}
	}
	return false
}

func FixQuotes(tokens []string, i int) []string {
	if i+1 >= len(tokens) {
		return tokens
	}

	tokens[i+1] = tokens[i] + tokens[i+1]
	tokens = append(tokens[:i], tokens[i+1:]...)

	for j := i; j < len(tokens); j++ {
		if IsQuote(tokens[j]) {
			if j == 0 {
				break
			}
			tokens[j-1] = tokens[j-1] + tokens[j]
			if j+1 == len(tokens) {
				tokens = tokens[:j]
			} else {
				tokens = append(tokens[:j], tokens[j+1:]...)
			}
			break
		}
	}

	return tokens
}
