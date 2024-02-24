package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 9, 10, 14, 17, 45, 56, 78} //define array
	key := 10                                        //define key

	result, err := binarySearch(arr, key, 0, len(arr)-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func binarySearch(arr []int, key, min, max int) (int, error) {
	if max < min {
		return 0, fmt.Errorf("Element not found") // element not found
	}

	pivotIndex := (min + max) / 2
	pivot := arr[pivotIndex]

	if key < pivot {
		return binarySearch(arr, key, min, pivotIndex-1) // search in left part
	} else if key > pivot {
		return binarySearch(arr, key, pivotIndex+1, max) // seRCH in right part
	} else {
		return pivotIndex, nil // return result
	}
}
