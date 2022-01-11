package main

import "fmt"

func main() {
	c := make(chan int)
	chs := make(chan string, 5)
	inp := make(chan string)
	out := make(chan string)

	go task(c)
	c <- 5

	go task2(chs, "test")

	fmt.Println(<-c)
	fmt.Println(<-chs)
	fmt.Println(<-chs)
	fmt.Println(<-chs)
	fmt.Println(<-chs)

	go removeDuplicates(inp, out)
	inp <- "DODO"
	fmt.Println(<-out)
	inp <- "DODt"
	fmt.Println(<-out)
}

func task(ch chan int) {
	n := <-ch
	ch <- n + 1
}

func task2(ch chan string, str string) {
	ch <- str + " "
	ch <- str + " 1"
	ch <- str + " 2"
	ch <- str + " 3"
	ch <- str + " 4"
}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	defer close(outputStream)
	var str, tmp string

	for str = range inputStream {
		if str != tmp {
			tmp = str

			outputStream <- str
		}
	}
}
