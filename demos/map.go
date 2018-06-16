package main

import "fmt"


func main(){
	var ml map[string]int
	ml = make(map[string]int, 100)
	ml["hello"] = 10
	fmt.Print(ml)

	for k, v := range ml {
		fmt.Println(k, v)
	}
}