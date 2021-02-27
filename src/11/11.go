// Program for finding the greatest product of four adjacent numbers in the same
// direction (up, down, left, right, or diagonally) in a 20×20 grid.
//
// Some portions copied from https://gobyexample.com/reading-files
//
// Armaan Roshani

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
	return
}

/* func readFile() [][]int {
	return
} */

/* func isPrime(toTest int) bool {
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
} */

func main() {

	/* 	dat, err := ioutil.ReadFile("data.txt")
	   	check(err)
	   	fmt.Print(string(dat)) */

	/* 	f, err := os.Open("data.txt")
	check(err) */

	//tempString := "01 02 03 04 05 06 07 08 09 10"
	var mainArray [][]int
	var tempArray []int
	var tempString string

	fptr := flag.String(".", "data.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	outerScanner := bufio.NewScanner(f)

	//i := 0
	count := 0
	for outerScanner.Scan() {

		fmt.Println(outerScanner.Text()) // debug
		tempString = outerScanner.Text()
		//fmt.Println(tempString)

		innerScanner := bufio.NewScanner(strings.NewReader(tempString))
		// Set the split function for the scanning operation.
		innerScanner.Split(bufio.ScanWords)

		// Count the words.

		temp := 0
		for innerScanner.Scan() {
			temp, _ = strconv.Atoi(innerScanner.Text())
			tempArray = append(tempArray, temp)

			//fmt.Println("Cycle", count) // debug
			//fmt.Println(tempArray) // debug

			//time.Sleep(500 * time.Millisecond) // debug

			count++
		}

		if err := innerScanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}

		mainArray = append(mainArray, tempArray)
		//fmt.Println(mainArray) // debug
		tempArray = nil
	}
	err = outerScanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	//mainArray = append(mainArray, tempArray)

	fmt.Printf("%d\n", count)
	fmt.Println(tempArray)
	fmt.Println(mainArray)

	return
}
