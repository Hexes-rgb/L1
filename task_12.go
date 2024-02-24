package main

import (
	"fmt"
	"strings"
)

type Set map[string]struct{} //define type for Set

func NewSet(elements []string) Set { // create and return set
	set := make(Set)

	for _, elem := range elements {
		set[elem] = struct{}{}
	}

	return set
}

func (s Set) String() string { //formatting the set string representation
	var elements []string
	for elem := range s {
		elements = append(elements, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"} //strings sequence

	set := NewSet(strings)

	fmt.Println(set)
}
