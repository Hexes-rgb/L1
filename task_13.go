package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	a, b := 5, 10

	a, b = b, a //use multiple assignment

	fmt.Println(a, b)
}

func example2() {
	a, b := 5, 10

	a = a + b //10 + 5 = 15
	b = a - b //15 - 10 = 5
	a = a - b //15 - 5 = 10

	fmt.Println(a, b)
}

func example3() {
	a, b := 5, 10 // 101 and 1010

	a = a ^ b // 0101 ^ 1010 = 1111
	b = a ^ b // 1111 ^ 1010 = 101
	a = a ^ b // 1111 ^ 0101 = 1010

	fmt.Println(a, b)
}

func example4() {
	a, b := 5, 10

	a = a * b //5 * 10 = 50
	b = a / b //50 / 10 = 5
	a = a / b //50 / 5 = 10

	fmt.Println(a, b)
}
