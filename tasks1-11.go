/* 1 - 99 korov */

package main

import "fmt"

func main() {
	var n byte
	var a string
	fmt.Scan(&n)

	copy := n
	if n > 20 {
		n = n % 10
	}

	if n == 1 {
		a = "korova"
	} else if n >= 2 && n <= 4 {
		a = "korovy"
	} else {
		a = "korov"
	}

	fmt.Printf("На лугу пасется %d %s", copy, a)

}
