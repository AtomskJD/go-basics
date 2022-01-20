package main

import "fmt"

func main() {
	var a string
	var rv []rune

	fmt.Scan(&a)
	r := []rune(a)
	for i := len(r) - 1; i >= 0; i-- {
		rv = append(rv, r[i])
	}
	if string(r) == string(rv) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
