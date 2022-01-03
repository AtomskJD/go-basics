package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func adding(s1 string, s2 string) int64 {
	n1, _ := strconv.Atoi(strings.TrimFunc(s1, func(r rune) bool { return !unicode.IsNumber(r) }))
	n2, _ := strconv.Atoi(strings.TrimFunc(s2, func(r rune) bool { return !unicode.IsNumber(r) }))

	return int64(n1 + n2)
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}
