package main

import "fmt"

func main() {

	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go myFoo(c)
	}
	fmt.Print(<-c)
}

func myFoo(c chan int) {
	c <- cap(c)
}
