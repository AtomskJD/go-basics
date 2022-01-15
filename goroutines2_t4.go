package main

import "fmt"

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- 1
	}
	close(done)
	fmt.Println(<-result)
}
func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	rc := make(chan int)
	go func() {
		defer close(rc)
		var sum int
		for {
			select {
			case arg := <-arguments:
				sum += arg
			case <-done:
				rc <- sum
				return
			}
		}
	}()
	return rc
}
