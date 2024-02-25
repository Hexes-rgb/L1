package main

import "fmt"

func removeElement(slice []int, index int) []int {
	// assign the copied elements after index to the elements starting from index
	copy(slice[index:], slice[index+1:])

	//  cut off the last extra element
	slice = slice[:len(slice)-1]

	return slice
}

func main() {
	numbers := []int{7, 2, 4, 1, 10}
	removeIndex := 3

	numbers = removeElement(numbers, removeIndex)
	fmt.Println(numbers)
}
