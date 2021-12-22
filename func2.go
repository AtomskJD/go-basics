package main

import "fmt"

func main() {
	var a [4]int
	var m int
	for i := 0; i < 4; i++ {
		fmt.Scan(&a[i])
	}
	m = a[0]
	for _, i2 := range a {
		if i2 < m {
			m = i2
		}
	}
	fmt.Println(m)
}
