/* степерь двоййки не превышающая N */

package main

import "fmt"

func main() {
	var n, a uint16
	a = 1

	fmt.Scan(&n)
	for {
		if a > n {
			break
		}
		fmt.Printf("%d ", a)
		a = a * 2
	}

}
