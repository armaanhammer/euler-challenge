// Program for finding the product abc such that a, b, c are a Pythagorean triplet
// and a + b + c = 1000.
//
// Armaan Roshani

package main

import (
	"fmt"
)

func loop(num int) int {
	var a, b, c int

	// reasoning for num/3 is that a < b < c and a + b + c = num
	for a = 1; a < (num / 3); a++ {
		for b = a + 1; b < (num / 2); b++ {
			c = num - (a + b)
			//fmt.Println("a: ", a, "\tb: ", b, "\tc: ", c) // debug
			if a*a+b*b == c*c {
				return a * b * c
			}
		}
	}

	return -1 // error
}

func main() {
	num := 1000
	var product int

	product = loop(num)

	fmt.Println("Product is: ", product)
}
