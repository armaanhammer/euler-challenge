// Program for finding the greatest product of four adjacent numbers in the same
// direction (up, down, left, right, or diagonally) in a 20×20 grid.
//
// Armaan Roshani

package main

import (
	"fmt"
)

func isPrime(toTest int) bool {
	// 1 is not prime
	if toTest == 1 {
		return false
	}

	// loop from 2 to int(sqrt(x))
	i := 2
	for i*i <= toTest {
		//fmt.Printf("Value of i: %d\n", i) // debug
		//time.Sleep(500 * time.Millisecond) // debug
		if toTest%i == 0 {
			return false
		}
		i++
	}

	// if the above loop found no factors, then number to test is prime
	return true
}

func main() {
	num := 2000000 // max number to increment up to
	sum := 2       // accounts for 2 being prime

	// i tracks current number under test
	for i := 3; i <= num; i = i + 2 {
		if isPrime(i) {
			sum = sum + i
			//fmt.Printf("%d is the %dth prime number.\n", prime, j) // debug
		} else {
			//fmt.Printf("%d is NOT the %dth prime number.\n", prime, j) // debug
		}
	}

	fmt.Printf("%d is the sum of all the primes below %d.\n", sum, num)
}
