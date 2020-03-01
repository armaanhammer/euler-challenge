// Program for finding the difference between the sum of the squares of
// the first one hundred natural numbers and the square of the sum.
//
// Armaan Roshani

package main

import (
	"fmt"
	"math"
)

func main() {
	num := 100.0 // number of naturals to process
	sumOfSquares := 0.0
	squareOfSum := 0.0
	diff := 0.0

	for i := 1.0; i <= num; i++ {
		sumOfSquares += math.Pow(i, 2)
		squareOfSum += i
	}

	fmt.Println(sumOfSquares)
	fmt.Println(squareOfSum)

	squareOfSum = math.Pow(squareOfSum, 2)
	fmt.Println(squareOfSum)

	diff = squareOfSum - sumOfSquares

	fmt.Printf("%f is the difference between the sum of the squares of the first %f natural numbers and the square of the sum.\n", diff, num)
	//fmt.Printf("%d is the difference between the sum of the squares of the first %d natural numbers and the square of the sum.\n", diff, num)

}
