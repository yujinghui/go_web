package main

import "fmt"

func main() {

	// simple declare
	var arr [10]int
	var pstr [10]*float64
	a := [3]int{1, 2, 3}
	b := [5]int{11, 2, 3}
	c := [...]int{1, 2, 3}
	d := [...]int{8: 18}
	fmt.Println(arr, pstr, a, b, c, d)

	// element visit
	intArray := [...]int{1, 2, 4, 5, 6, 3}
	fmt.Println(len(intArray))
}
