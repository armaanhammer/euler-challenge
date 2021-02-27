// Program for finding the greatest product of four adjacent numbers in the same
// direction (up, down, left, right, or diagonally) in a 20×20 grid.
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

func arrayFromFile(fileName string) [][]int {
	var mainArray [][]int
	var tempArray []int
	var tempString string

	fptr := flag.String(".", fileName, "file path to read from")
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

	// Create the scanner to scan lines
	outerScanner := bufio.NewScanner(f)
	for outerScanner.Scan() {

		tempString = outerScanner.Text()

		// Create the scanner to scan words
		innerScanner := bufio.NewScanner(strings.NewReader(tempString))
		innerScanner.Split(bufio.ScanWords)

		// Cycle through the individual words, conv from string to int
		temp := 0
		for innerScanner.Scan() {
			temp, _ = strconv.Atoi(innerScanner.Text())
			tempArray = append(tempArray, temp)
		}

		if err := innerScanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}

		mainArray = append(mainArray, tempArray)
		tempArray = nil // reset the temp array
	}
	err = outerScanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return mainArray
}

func goRight(mainArray [][]int, num int) int {
	product := 1 // multiplication identity

	return product
}

func goDown(mainArray [][]int, num int) int {
	product := 1 // multiplication identity

	return product
}

func goDiagDown(mainArray [][]int, num int) int {
	product := 1 // multiplication identity

	return product
}

func goDiagUp(mainArray [][]int, num int) int {
	product := 1 // multiplication identity

	return product
}

func main() {

	/* 	dat, err := ioutil.ReadFile("data.txt")
	   	check(err)
	   	fmt.Print(string(dat)) */

	/* 	f, err := os.Open("data.txt")
	check(err) */

	//tempString := "01 02 03 04 05 06 07 08 09 10"

	product := 0
	temp := 0
	num := 4 // number of digits to multiply
	var mainArray = arrayFromFile("data.txt")
	fmt.Println(mainArray)

	// cycle through all avaliable numbers
	for i := 0; i <= len(mainArray)-num; i++ {
		for j := 0; j <= len(mainArray[i])-4; j++ {
			fmt.Printf("%d ", mainArray[i][j]) // debug

			//test all four operations
			temp = goRight(mainArray, num)
			if temp > product {
				product = temp
			}
			temp = goDown(mainArray, num)
			if temp > product {
				product = temp
			}
			temp = goDiagDown(mainArray, num)
			if temp > product {
				product = temp
			}
			temp = goDiagUp(mainArray, num)
			if temp > product {
				product = temp
			}

		}
	}

	return
}
