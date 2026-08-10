package cases

import (
	"strings"
	"unicode"
)

func isValidWord( word string ) bool {
	for _, ch := range word {
		if !(unicode.IsLetter(ch)){
			return false
		}
	}
	return true
}


func UpN( tokens []string, num int, i int ) []string {
	for j := i - 1; j >= 0; j-- {
		if num > 0 {
			if isValidWord(tokens[j]){
				tokens[j] = strings.ToUpper(tokens[j])
				num--
			}
		}
	}
	return tokens
}

func LowN( tokens []string, num int, i int ) []string {
	for j := i - 1; j >= 0; j-- {
		if num > 0 {
			if isValidWord(tokens[j]){
				tokens[j] = strings.ToLower(tokens[j])
				num--
			}
		}
	}
	return tokens
}

func CapN( tokens []string, num int, i int ) []string {
	for j := i - 1; j >= 0; j-- {
		if num > 0 {
			if isValidWord(tokens[j]){
				word := strings.ToLower(tokens[j])
				tokens[j] = strings.ToUpper(string(word[0]))+word[1:]
				num--
			}
		}
	}
	return tokens
}

