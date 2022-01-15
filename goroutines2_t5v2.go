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
	for i, prev := 0, 0; i < N; i++ {
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
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	var x int
	var x1, x2 []int
	for i := 0; i < n; i++ {
		go func() {
			wg.Add(2)
			go func() {
				defer wg.Done()
				x1 = append(x1, fn(<-in1))
				fmt.Println("x1 = ", x1)
			}()
			go func() {
				defer wg.Done()
				x2 = append(x2, fn(<-in2))
				fmt.Println("x2 = ", x2)
			}()

			wg.Wait()
			mu.Lock()
			fmt.Println("i = ", x)
			out <- x1[x] + x2[x]
			x++
			mu.Unlock()
		}()
	}
}
