/**
Последовательност фибоначи по n найти Фn
*/
package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	prn := fibonaci(s)
	fmt.Println(prn)
}

func fibonaci(n int) int {
	var n1 int = 0
	var n2 int = 1
	var i int = 1
	tmp := 0
	if n < 2 {
		return 1
	}
	for {
		i++
		tmp = n1
		n1 = n2
		n2 = tmp + n2
		if i == n {
			return n2
		}
	}
}
