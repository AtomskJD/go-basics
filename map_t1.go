/**
 Выделенного вам времени выполнения не хватит на
последовательную обработку каждого числа,
поэтому необходимо реализовать кэширование уже
готовых результатов и использовать их в работе.
*/
package main

import "fmt"

func work(x int) int {
	return x * 10
}

func main() {
	var k int
	m := make(map[int]int)

	for i := 0; i < 2; i++ {
		fmt.Scan(&k)
		if v, ok := m[k]; ok {
			fmt.Printf("%d ", v)
		} else {
			m[k] = work(k)
			fmt.Printf("*%d ", m[k])
		}
	}
}
