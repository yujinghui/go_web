package main

import "fmt"

func main() {
	// slice declare
	// var slice []int;
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice1 = arr[0:5]
	fmt.Println(slice1)
}
