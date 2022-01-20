/*
Print all odd numbers between 2 and 1023
using loops, each on a new line.
*/
package main

import "fmt"

func main() {
	for i := 2; i <= 1023; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
