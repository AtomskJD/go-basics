package main

import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	//ch1 <- 3
	//ch2 <- 4
	close(stop)
	fmt.Println(<-r)
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	rc := make(chan int)
	go func() {
		defer close(rc)

		select {
		case a := <-firstChan:
			rc <- a * a
		case a := <-secondChan:
			rc <- a * 3
		case <-stopChan:
			return
		}
	}()
	return rc
}
