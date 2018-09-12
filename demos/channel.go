package main

import "fmt"
import "time"

func main() {
	timeout := make(chan bool)
	cha := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()

	for i := 1; i < 10; i++ {
		go func() {
			cha <- i
		}()
	}
	for {
		select {
		case <-cha:
			fmt.Println("get data from cha")
		case <-timeout:
			fmt.Println("timeout")
		}
	}
}
