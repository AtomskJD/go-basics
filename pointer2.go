/**
Поменяйте местами значения переменных
на которые ссылаются указатели.
После этого переменные нужно вывести.
*/
package main

import "fmt"

func main() {
	x1 := 1
	x2 := 2
	test(&x1, &x2)
}
func test(x1 *int, x2 *int) {
	tmp := *x1
	*x1 = *x2
	*x2 = tmp
	fmt.Printf("%d %d", *x1, *x2)

}
