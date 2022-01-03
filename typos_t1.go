package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(a, b string) int64 {
	aint, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		aint = extractNumbers(a)
	}
	bint, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		bint = extractNumbers(b)
	}

	return aint + bint
}

func extractNumbers(s string) int64 {
	var str string
	r := []rune(s)
	for i, _ := range r {
		if unicode.IsNumber(r[i]) {
			str += string(r[i])
		}
	}
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}
