// Golang program to check whether a number is palindrome or not
// from: https://www.golangprograms.com/go-program-to-check-whether-a-number-is-palindrome-or-not.html

package main

import "fmt"

var digits int = 4 // size of ints to multiply

func main() {
	bigInt := 0

	// create ints to multiply together
	// populate with 9s to start at largest possible product
	for i := 0; i < digits; i++ {
		bigInt *= 10
		bigInt += 9
	}

	//fmt.Println(int1, int2)

	var number, remainder, temp, final int
	var reverse int = 0

	//fmt.Print("Enter any positive integer : ")
	//fmt.Scan(&number)

	for i := bigInt; i > bigInt/10; i-- { //need better test condition
		for j := bigInt; j > bigInt/10; j-- { //need better test condition
			number = i * j
			temp = number
			reverse = 0

			//fmt.Println("made into first loop") //debug

			for {
				//fmt.Println("made into SECOND loop") //debug
				remainder = number % 10
				reverse = reverse*10 + remainder
				number /= 10

				if number == 0 {
					//fmt.Println("temp:", temp, "\treverse:", reverse)
					break // Break Statement used to exit from loop
				}
			}

			//fmt.Println("temp:", temp)
			//fmt.Println("reverse:", reverse)

			if temp == reverse {
				if final < temp {
					final = temp
				}
				break
			}
		}
	}
	fmt.Printf("%d is the largest Palindrome formed by multiplying two %d-digit numbers.\n", final, digits)

}
