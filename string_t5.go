package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	for i := range r {
		if strings.Count(s, string(r[i])) == 1 {
			fmt.Printf("%c", r[i])
		}
	}

}
