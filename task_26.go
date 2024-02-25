package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ContainsOnlyUniqueSymbols("abcd"))
	fmt.Println(ContainsOnlyUniqueSymbols("abCdefAaf"))
	fmt.Println(ContainsOnlyUniqueSymbols("aabcd"))
}

// Checking that all symbols in string are unique
func ContainsOnlyUniqueSymbols(str string) bool {
	lowerStr := strings.ToLower(str) //get lower case string

	for _, symbol := range lowerStr {
		if strings.Count(lowerStr, string(symbol)) > 1 { //count current symbol in string
			return false
		}
	}

	return true
}
