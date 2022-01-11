package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second * 2)
	defer tick.Stop()

	wg := new(sync.WaitGroup)

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(i, tick.C, wg)
	}
	wg.Wait()
}

func worker(id int, limit <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	<-limit
	fmt.Printf("worker %d end working\n", id)
}
