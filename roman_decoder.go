package main

import "fmt"

func main() {
	var rom string
	fmt.Scan(&rom)
	fmt.Println(Decode(rom))
}

func Decode(roman string) int {
	renderMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	r := []rune(roman)
	var sum int

	for i, v := range r {
		if len(r) > i+1 {
			if renderMap[r[i]] < renderMap[r[i+1]] {
				sum -= renderMap[r[i]]
			} else {
				sum += renderMap[r[i]]
			}
		} else {
			sum += renderMap[v]
		}
	}
	return sum
}
