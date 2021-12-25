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
