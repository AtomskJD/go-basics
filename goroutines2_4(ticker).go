package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(time.Second * 2)
	timer := time.After(time.Second)
	<-timer
	ticker := time.Tick(time.Second * 3)

	count := 0
	for {
		<-ticker
		fmt.Println("TICK")
		count++
		if count == 3 {
			break
		}
	}

	for _ = range ticker {
		fmt.Println("TICK NEW")
		count++
		if count == 5 {
			break
		}
	}
}
