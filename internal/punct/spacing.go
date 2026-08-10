package punct

func IsPunctuations(word string) bool {
	if len(word) == 1 {
		if ( word == "." || word == "," || word == "!" || word == "?" || word == ":" ){
			return true
		}
	} else if len(word) > 1 {
		for i := 0; i < len(word); i++ {
			if !( string(word[i]) == "." || string(word[i]) == "," || string(word[i]) == "!" || string(word[i]) == "?" || string(word[i]) == ":" ){
				return false
			}
		}
		return true
	}
	return false
}

func FixSpacing(tokens []string, i int) []string {
	tokens[i-1] = tokens[i-1] + tokens[i]

	tokens = append(tokens[:i], tokens[i+1:]...)

	return tokens
}
