/* находим порядковый номер фибоначи */

package main

import "fmt"

func main() {
	var n1 int = 0
	var n2 int = 1
	var n int16 = 2
	var s int
	fmt.Scan(&s)
	tmp := 0
	for {
		n++
		tmp = n1
		n1 = n2
		n2 = tmp + n2
		// fmt.Printf("%d - %d\n", n2, n)

		if n2 == s {
			fmt.Println(n)
			break
		}

		if n2 > 1000000000 {
			fmt.Println("-1")
			break
		}
	}
}
