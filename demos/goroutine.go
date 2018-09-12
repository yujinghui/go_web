package main

import "fmt"
import "time"

func main() {
	future := heavyCalculation()
	fmt.Println(<-future)
}

func heavyCalculation() chan int {

	future := make(chan int)
	go func() {
		//模拟耗时计算
		time.Sleep(1e9)
		future <- 666
	}()

	return future
}
