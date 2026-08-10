package quotes


func IsQuote(word string) bool {
	if len(word) == 1  {
		if word == "'"  {
			return true
		}
	}
	return false
}

func FixQuotes(tokens []string, i int) []string {

	isOpenQuo := true

	if isOpenQuo && (len(tokens) != i+1) {
		tokens[i+1] = tokens[i] + tokens[i+1]
		tokens = append(tokens[:i], tokens[i+1:]...)
		isOpenQuo = false
		i--

		for j := i+1; j < len(tokens); j++ {
			if IsQuote(tokens[j]) && !(isOpenQuo){
				tokens[j-1] = tokens[j-1] + tokens[j]

				if j+1 == len(tokens){
					tokens = tokens[:j]
				} else {
					tokens = append(tokens[:j], tokens[j+1:]...)
				}
			}
		}
	} 

	return tokens
}
