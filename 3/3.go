package main

import (
	"fmt"
	"math"
)

func main() {

	num := 600851475143

	// 1) While n is divisible by 2, divide n by 2.
	for num%2 == 0 {
		num = num / 2
	}

	// 2) After step 1, n must be odd. Now start a loop from i = 3 to square root of n.
	// While i divides n, print i and divide n by i. After i fails to divide n,
	// increment i by 2 and continue.
	sqrtNum := int(math.Round(math.Sqrt(float64(num))))

	result := 0
	for odds := 3; odds <= sqrtNum; odds += 2 {
		for num%odds == 0 {
			result = odds
			num = num / odds
		}
	}

	fmt.Println("Largest prime factor:", result)
}
