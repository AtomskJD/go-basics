package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	r := []rune(s)
	for i := range r {
		if i%2 != 0 {
			fmt.Printf("%c", r[i])
		}
	}
}
