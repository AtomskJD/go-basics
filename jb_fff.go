package main

import "fmt"

func main() {
	x := 2
	if x < 0 {
		fmt.Println(f1(x))
	}
	fmt.Println(f2(x))
}

func f1(x int) int {
	// your code here
	return x * 2
}

func f2(x int) int {
	// your code here
	return x * 3
}
