/*Дан массив, состоящий из целых чисел. Нумерация элементов
начинается с 0. Напишите программу, которая выведет элементы 
массива, индексы которых четны (0, 2, 4...).*/
package main
import "fmt"
func main(){
	var N, ai, i int
	var a []int

	fmt.Scan(&N)
	if N >= 1 && N <= 100 {
		for i = 0; i < N; i++ {
			fmt.Scan(&ai)
			a = append(a, ai)
		}

		for idx, elem := range a {
			if idx % 2 == 0 {
				fmt.Println(elem)
			}
		}
	}
}