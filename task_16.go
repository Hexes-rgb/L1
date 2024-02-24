package main

import "fmt"

func main() {
	arr := []int{2, 1, 4, 0, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low >= high || low < 0 { // сheck if recursion end is reached
		return
	}

	p := partition(arr, low, high) // partition the array and get the index of the pivot element

	quickSort(arr, low, p-1)  // sort left part
	quickSort(arr, p+1, high) //sort the right part
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] //swap current element with element with index i if current element <= pivot
		}
	}

	i++
	arr[i], arr[high] = arr[high], arr[i] // place the pivot element in its final position by swapping it with the element at index i+1
	return i
}
