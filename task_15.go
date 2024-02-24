package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] //This code will cause a memory leak. A slice is created that references the same array v.
	//justString itself cannot store a reference to the array. As a result, the memory allocated for v cannot be freed.
}

func someFuncCorrect() {
	v := createHugeString(1 << 10)
	justString = string(v[:100]) //it would be more correct to copy the slice data to a string so that the memory occupied
	// by v can be freed, since there will be no references to it
}

func main() {
	someFunc()
	someFuncCorrect()
}
