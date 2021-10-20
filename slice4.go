/*Дана последовательность, состоящая из целых чисел. 
Напишите программу, которая подсчитывает количество положительных 
чисел среди элементов последовательности.*/
package main
import "fmt"
func main(){
	var N, ai, i int
	var a []int

	fmt.Scan(&N)
	if N >= 1 && N <= 100 {
		for i = 0; i < N; i++ {
			fmt.Scan(&ai)
			if ai > 0 {

				a = append(a, ai)
			}
		}
		fmt.Println(len(a))
	}
}