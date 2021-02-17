// Program for finding the 10 001st prime number
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
	num := 10001 // number of primes to process
	prime := 2   // specific prime

	// i tracks current number under test
	// j is number of primes found
	j := 2 // accounts for 2
	for i := 3; j <= num; i = i + 2 {
		if isPrime(i) {
			prime = i
			//fmt.Printf("%d is the %dth prime number.\n", prime, j) // debug
			j++
		} else {
			//fmt.Printf("%d is NOT the %dth prime number.\n", prime, j) // debug
		}
	}

	fmt.Printf("%d is the %dth prime number.\n", prime, num)
}
