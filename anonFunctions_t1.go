package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint = 123022
	fn := func(x uint) uint {
		var s string
		xs := strconv.FormatInt(int64(x), 10)
		r := []rune(xs)
		for _, re := range r {
			if re%2 == 0 && re != 48 {
				s += string(re)
			}
		}

		res, _ := strconv.ParseInt(s, 10, 64)
		if res == 0 {
			res = 100
		}
		return uint(res)
	}

	fmt.Println(fn(a))
}
