// program for finding the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20
//
// Armaan Roshani

package main

import "fmt"

func main() {

	testCond := 0   // if all five conditions met, solution found
	initVal := 2520 // since obviously will be larger than example number

	// there has got to be a better way to iterate through [2:19]
	testDivs := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	for ; testCond < len(testDivs); initVal += 20 {
		testCond = 0 // reset

		// iterate through all test divisors
		for j := range testDivs {
			//fmt.Println(testDivs[j], initVal/testDivs[j], initVal%testDivs[j])
			if initVal%testDivs[j] == 0 {
				testCond += 1
			}
		}
	}

	initVal -= 20 // doh!

	fmt.Printf("%d is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20\n", initVal)

}
