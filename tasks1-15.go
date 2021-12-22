package main

import "fmt"

func main() {
	var a int
	var b int
	var tmp1, tmp2 int
	fmt.Scan(&a, &b)
	for a > 0 {
		if a%10 != b {
			tmp1 = tmp1 * 10
			tmp1 += a % 10
		}
		a = a / 10

	}

	for tmp1 > 0 {
		tmp2 = tmp2*10 + tmp1%10
		tmp1 = tmp1 / 10
	}

	fmt.Println(tmp2)
}
