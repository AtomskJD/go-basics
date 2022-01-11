package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second)
	go func() {
		<-t.C
		fmt.Println("END timer 3")
	}()
	fmt.Println("END timer 1")
	t.Stop()
	t.Reset(time.Second * 3)

	<-t.C
	fmt.Println("END timer 2")

}
