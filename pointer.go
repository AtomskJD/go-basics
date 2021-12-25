/**
Напишите функцию, которая умножает значения
на которые ссылаются два указателя и выводит
получившееся произведение в консоль.
*/
package main

import "fmt"

type In struct {
	x1 int
	x2 int
}

func main() {
	x1 := 3
	x2 := 2

	x := In{3, 2}

	test(&x1, &x2)
	test2(&x)
}

// считайте что fmt уже импортирован и main объявлен
func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}

// при использовании структур
func test2(x *In) {
	fmt.Println(x.x1 * x.x2)
}
