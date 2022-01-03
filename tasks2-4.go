package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)
	r := []rune(a)
	for _, el := range r {
		elInt, _ := strconv.Atoi(string(el))
		fmt.Print(elInt * elInt)
	}
}
