package numbers

import (
	"strconv"
)

func IsValidHexWord( word string ) bool {
	for _, ch := range word {
		if !(( ch >= 48 && ch <= 57 ) || ( ch >= 65 && ch <= 70 ) || ( ch >= 97 && ch <= 102 )){
			return false
		}
	}
	return true
}

func Hex( tokens []string, i int ) []string {
	num := 1
	for j := i - 1; j >= 0; j-- {
		if num > 0 {
			if IsValidHexWord(tokens[j]){
				n, err := strconv.ParseInt(tokens[j], 16, 64)
				if err != nil {
					return tokens
				}
				tokens[j] = strconv.FormatInt(n, 10)
				num--
			}
		}
	}
	return tokens
}

func IsValidBinWord( word string ) bool {
	for _, ch := range word {
		if !( ch == 48 || ch == 49 ) {
			return false
		}
	}
	return true
}

func Bin( tokens []string, i int ) []string {
	num := 1
	for j := i - 1; j >= 0; j-- {
		if num > 0 {
			if IsValidBinWord(tokens[j]){
				n, err := strconv.ParseInt(tokens[j], 2, 64)
				if err != nil {
					return tokens
				}
				tokens[j] = strconv.FormatInt(n, 10)
				num--
			}
		}
	}
	return tokens
}

