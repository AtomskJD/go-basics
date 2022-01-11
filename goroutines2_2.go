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
		go work(i, wg)
	}
	wg.Wait()
	fmt.Println("goroutines finished")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("GoRoutine %d start\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("GoRoutine %d ends\n", id)

}
