/**
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений
ее аргументов x, y, z встречается чаще.
*/
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	output := vote(x, y, z)
	fmt.Println(output)
}

func vote(x, y, z int) int {
	summ := x + y + z
	if summ > 1 {
		return 1
	} else {
		return 0
	}
}
