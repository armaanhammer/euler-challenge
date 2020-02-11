package main

import (
	"fmt"
)

func main() {

	num := 1000
	//natural := 1

	// 1) While n is divisible by 2, divide n by 2.
	for num%2 == 0 {
		fmt.Println(num)
		num = num / 2
	}

	fmt.Println(num % 2)
}
