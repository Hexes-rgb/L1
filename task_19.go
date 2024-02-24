package main

import "fmt"

func main() {
	str := "абырвалг"

	revertString := reverseString(str)
	fmt.Println(str, revertString)
}

func reverseString(str string) string {
	runeSlice := []rune(str) //convert the string into a slice of runes so that you can work with unicode characters

	for i, j := len(runeSlice)-1, 0; i > j; i-- { //iterate along the string in both directions
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i] //swap runes
		j++
	}
	return string(runeSlice)
}
