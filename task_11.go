package main

import (
	"fmt"
	"strings"
)

type Set map[interface{}]struct{} //define type for Set

func NewSet(elements []interface{}) Set { // create and return set
	set := make(Set)

	for _, elem := range elements {
		set[elem] = struct{}{}
	}

	return set
}

func (s1 Set) intersection(s2 Set) Set {
	intersection := make(Set)

	for val := range s1 {
		if _, exists := s2[val]; exists { //Checking that an element from the first set is in the second set
			intersection[val] = struct{}{}
		}
	}
	return intersection
}

func (s Set) String() string { //formatting the set string representation
	var elements []string
	for elem := range s {
		elements = append(elements, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}

func main() {
	set1 := NewSet([]interface{}{1, "apple", 3, 6, "gold"})
	set2 := NewSet([]interface{}{15, "gold", 3})
	set3 := set1.intersection(set2)

	fmt.Println("Intersection:", set3)
}
