package main

import "fmt"

func main() {
	upperBound := 4000000
	curr := 2
	prev := 1
	sum := 0

	for curr < upperBound {
		if curr%2 == 0 {
			sum += curr
			fmt.Println("curr", curr, "prev", prev, "sum", sum)
		}
		curr += prev
		prev = curr - prev
	}

	fmt.Println("Final sum:", sum)
}
