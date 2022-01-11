package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			work()
			wg.Done()
		}()
	}
	wg.Wait()

}
func work() {
	fmt.Println("Goroutine start")
	time.Sleep(time.Second * 5)
	fmt.Println("Goroutine stop")
}
