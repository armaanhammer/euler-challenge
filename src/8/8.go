// Program for finding the 10 001st prime number
//
// Armaan Roshani

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func calcProduct(num []int) int {
	i := 1
	for _, j := range num {
		i = i * j
	}
	//time.Sleep(500 * time.Millisecond) // debug
	return i
}

func main() {
	var num bytes.Buffer
	var ints [1000]int
	//digits := 13
	product := 0
	finalProduct := 0

	num.WriteString("73167176531330624919225119674426574742355349194934")
	num.WriteString("96983520312774506326239578318016984801869478851843")
	num.WriteString("85861560789112949495459501737958331952853208805511")
	num.WriteString("12540698747158523863050715693290963295227443043557")
	num.WriteString("66896648950445244523161731856403098711121722383113")
	num.WriteString("62229893423380308135336276614282806444486645238749")
	num.WriteString("30358907296290491560440772390713810515859307960866")
	num.WriteString("70172427121883998797908792274921901699720888093776")
	num.WriteString("65727333001053367881220235421809751254540594752243")
	num.WriteString("52584907711670556013604839586446706324415722155397")
	num.WriteString("53697817977846174064955149290862569321978468622482")
	num.WriteString("83972241375657056057490261407972968652414535100474")
	num.WriteString("82166370484403199890008895243450658541227588666881")
	num.WriteString("16427171479924442928230863465674813919123162824586")
	num.WriteString("17866458359124566529476545682848912883142607690042")
	num.WriteString("24219022671055626321111109370544217506941658960408")
	num.WriteString("07198403850962455444362981230987879927244284909188")
	num.WriteString("84580156166097919133875499200524063689912560717606")
	num.WriteString("05886116467109405077541002256983155200055935729725")
	num.WriteString("71636269561882670428252483600823257530420752963450")

	//fmt.Println("This is the number:", num.String())
	//fmt.Println("This is the raw bytes:", num.Bytes())
	raw := num.Bytes()
	//fmt.Println("Test", reflect.TypeOf(raw[3]))
	//fmt.Println("Test", raw[3]*2)
	//fmt.Printf("%s", num)

	// convert uint8 value to int value
	// key and value pair
	for k, _ := range raw {
		//fmt.Println(k, v)
		//raw[k] = raw[k] - 48
		ints[k], _ = strconv.Atoi(string(raw[k]))
	}

	fmt.Println("Test", raw)
	fmt.Println("Test", ints)

	for k := range ints {
		product = calcProduct(ints[k : k+13])
		if product > finalProduct {
			finalProduct = product
		}
		fmt.Println("num:", k, " product: ", product, " final: ", finalProduct)
	}

}
