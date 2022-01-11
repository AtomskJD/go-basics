package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	go myFuncA(done)
	<-done

	// the same result but in pattern
	<-myFuncB()
}

func myFuncA(done chan struct{}) {
	fmt.Println("hello")
	close(done)
}

func myFuncB() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()
	return done
}
