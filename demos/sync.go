package main

import "fmt"
import "time"
import "sync"

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("this is the goroutine-", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("over")
}
