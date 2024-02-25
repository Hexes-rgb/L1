package main

import (
	"fmt"
	"math/big"
)

func main() {
	// use big.Int to work with large numbers
	a := big.NewInt(2)
	b := big.NewInt(2)

	a.Exp(a, big.NewInt(23), nil) // a = 2^23
	b.Exp(b, big.NewInt(21), nil) // b = 2^21

	// multiply
	multiply_result := new(big.Int).Mul(a, b)

	//use big.Float to get accurate division results
	// convert a and and to big.Float and perform the division
	floatA := new(big.Float).SetInt(a)
	floatB := new(big.Float).SetInt(b)
	divide_result := new(big.Float).Quo(floatA, floatB)

	// adding
	add_result := new(big.Int).Add(a, b)

	// subtract
	subtract_result := new(big.Int).Sub(a, b)

	// results
	fmt.Printf("a: %v\nb: %v\n", a, b)
	fmt.Println("Multiply a * b:", multiply_result)
	fmt.Println("Divide a / b:", divide_result)
	fmt.Println("Add a + b:", add_result)
	fmt.Println("Subtract a - b:", subtract_result)
}
