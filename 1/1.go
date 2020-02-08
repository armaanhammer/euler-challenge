package main

import "fmt"

func main() {
	upperBound := 1000
	nums := []int{3, 5}
	sum := 0

	// set a for loop
	for i := 0; i < upperBound; i++ {
		for j := range nums {
			if i%nums[j] == 0 {
				sum += i
				fmt.Println("num:", i, "sum:", sum)
				break // to prevent multiples being caught multiple times
			}
		}
	}

	fmt.Println("Final sum:", sum)
}
