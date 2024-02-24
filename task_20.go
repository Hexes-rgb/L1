package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := "snow dog sun"
	str := "рыба воин краска"
	fmt.Println(reverseSentence(str))
}

func split(str, separator string) []string {
	var splitString []string //use string array for words

	subStrStart := 0 //current word start symbol index

	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i:], separator) { // if i symbol == separator
			splitString = append(splitString, str[subStrStart:i]) //append substr from start to separator to array
			subStrStart = i + 1
		}
	}

	splitString = append(splitString, str[subStrStart:]) //append last word
	return splitString
}

func reverseSentence(sentence string) string {
	splitSentence := split(sentence, " ") //split sentence

	for i, j := len(splitSentence)-1, 0; i > j; i-- { //iterate along the array in both directions
		splitSentence[i], splitSentence[j] = splitSentence[j], splitSentence[i] //swap words
		j++
	}
	return fmt.Sprintf("%s", strings.Join(splitSentence, " ")) //formatted string
}
