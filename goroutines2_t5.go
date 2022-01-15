package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	Merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i <= N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

// Merge2Channels below

func Merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		wg := new(sync.WaitGroup)
		var x1, x2 []int
		for i := 0; i < n; i++ {
			x1 = append(x1, <-in1)
			x2 = append(x2, <-in2)
		}
		wg.Add(n * 2)
		for i := 0; i < n; i++ {
			go func(i int) {
				defer wg.Done()
				x1[i] = fn(x1[i])
			}(i)
			go func(i int) {
				defer wg.Done()
				x2[i] = fn(x2[i])
			}(i)
		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- x1[i] + x2[i]
		}
	}()
}
