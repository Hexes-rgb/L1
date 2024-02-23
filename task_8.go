package main

import "fmt"

func main() {
	var num int64
	var pos, bit byte

	fmt.Print("Enter number, position, 0 or 1: ")
	fmt.Scan(&num, &pos, &bit)

	if bit == 1 {
		num |= 1 << pos //set a bit in position using logical OR with assignment
	} else {
		num &^= 1 << pos //reset a bit in position using logical XOR with assignment
	}

	fmt.Println(num)

}
